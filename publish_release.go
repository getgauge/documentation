package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	newDirPermissions = 0755
	ghPages           = "gh-pages"
	bookDir           = "_book"
	defaultUserName   = "gaugeci"
	defaultUserEmail  = "getgauge@googlegroups.com"
)

var updateCurrent = flag.Bool("update-current", false, "Update current docs symlink to point to version")
var version = flag.String("version", "", "Version of documentation to update")

func main() {
	flag.Parse()
	if *version == "" {
		log.Panic("Docs version not specified")
	}
	updateDocs()
}

func updateDocs() {
	bookPath := buildGitBook()
	switchToGitBranch()
	copyDocs(bookPath, *version)
	commitAndPushChanges()
	cleanUp()
}

func commitAndPushChanges() {
	setCredentials()
	runCommand("git", "add", filepath.Join("user"))
	runCommand("git", "commit", "-m", fmt.Sprintf("Updating docs for version %v", *version))
	pushChanges()
}

func pushChanges() {
	defer func() {
		if error := recover(); error != nil {
			runCommand("git", "reset", "HEAD~1")
			runCommand("git", "checkout", "-f", "master")
		}
	}()
	runCommand("git", "push", "origin", ghPages)
}

func copyDocs(bookPath string, version string) {
	mirrorDir(bookPath, filepath.Join("user", version))
	if *updateCurrent {
		runCommand("rm", filepath.Join("user", "current"))
		runCommand("ln", "-s", version, filepath.Join("user", "current"))
	}
}

//Todo: Use git2go
func switchToGitBranch() {
	output := getCommandOutput("git", "branch", "-l")
	if strings.Contains(output, ghPages) {
		runCommand("git", "checkout", ghPages)
	} else {
		runCommand("git", "checkout", "-b", ghPages, "--track", "origin/gh-pages")
	}
}

func buildGitBook() string {
	runCommand("gitbook", "install")
	runCommand("gitbook", "build", ".")
	bookPath := filepath.Join(os.TempDir(), bookDir)
	mirrorDir(bookDir, bookPath)
	return bookPath
}

func cleanUp() {
	getCommandOutput("git", "checkout", "-f", "master")
	getCommandOutput("rm", "-rf", bookDir)
}

func setCredentials() {
	uName := getCommandOutput("git", "config", "user.name")
	if strings.TrimSpace(uName) == "" {
		runCommand("git", "config", "user.name", defaultUserName)
	}
	uEmail := getCommandOutput("git", "config", "user.email")
	if strings.TrimSpace(uEmail) == "" {
		runCommand("git", "config", "user.email", defaultUserEmail)
	}
}

func runCommand(command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	log.Printf("Running => %s %s", command, arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}

func getCommandOutput(command string, arg ...string) string {
	cmd := exec.Command(command, arg...)
	log.Printf("Running => %s %s", command, arg)
	output, _ := cmd.Output()
	return string(output)
}

func isExecMode(mode os.FileMode) bool {
	return (mode & 0111) != 0
}

func mirrorFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if sfi.Mode()&os.ModeType != 0 {
		log.Fatalf("mirrorFile can't deal with non-regular file %s", src)
	}
	dfi, err := os.Stat(dst)
	if err == nil &&
		isExecMode(sfi.Mode()) == isExecMode(dfi.Mode()) &&
		(dfi.Mode()&os.ModeType == 0) &&
		dfi.Size() == sfi.Size() &&
		dfi.ModTime().Unix() == sfi.ModTime().Unix() {
		// Seems to not be modified.
		return nil
	}

	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, newDirPermissions); err != nil {
		return err
	}

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	n, err := io.Copy(df, sf)
	if err == nil && n != sfi.Size() {
		err = fmt.Errorf("copied wrong size for %s -> %s: copied %d; want %d", src, dst, n, sfi.Size())
	}
	cerr := df.Close()
	if err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Chmod(dst, sfi.Mode())
	}
	if err == nil {
		err = os.Chtimes(dst, sfi.ModTime(), sfi.ModTime())
	}
	return err
}

func mirrorDir(src, dst string) {
	log.Printf("Copying '%s' -> '%s'\n", src, dst)
	err := filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		suffix, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("Failed to find Rel(%q, %q): %v", src, path, err)
		}
		return mirrorFile(path, filepath.Join(dst, suffix))
	})
	if err != nil {
		log.Panic(err)
	}
}
