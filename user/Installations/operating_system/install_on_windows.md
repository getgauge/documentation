# Install Gauge on Windows
There are two ways to install Gauge on Windows.

* [Using the installer](#1-using-installer)
* [Using chocolatey](#2-using-chocolatey)

## Using the installer

[Download](http://getgauge.io/get-started) the installer based on your system configuration and run it. Click your way through till you're asked to select your language.

The Windows installer allows you to select the language plugin(s) as part of the the installation process. A language plugin is essentially the language(s) that you would use to write your tests in. Check the box(es) you want to install. You can select more than one language plugin to install.

![Select language runner](images/install-lang-runner.jpg)

Gauge is installed in `%PROGRAMFILES%` by default. But you can select where you want to install Gauge and complete Gauge installation.

## Using Chocolatey
You can install Gauge using Chocolatey as well.

```
> choco install gauge

```

If you're upgrading to a newer version, then use the command:

```
> choco upgrade gauge

```
