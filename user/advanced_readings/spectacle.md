# Spectacle

This is a Gauge plugin that generates static HTML from Specification/Markdown files. Ability to filter specifications and scenarios are available.

## Install through Gauge

```
gauge --install spectacle
```

## Export to HTML

Run the following command to export to HTML in a Gauge project

```
gauge --docs spectacle <path to specs dir>
```

## Build from Source

### Pre-requisites

Golang

### Compiling
```
go run build/make.go
```

For cross-platform compilation

```
go run build/make.go --all-platforms
```

### Installing

After compilation
```
go run build/make.go --install
```

### Creating distributable

Note: Run after compiling
```
go run build/make.go --distro
```

For distributable across platforms: Windows and Linux for both x86 and x86_64
```
go run build/make.go --distro --all-platforms
```
