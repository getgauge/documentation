# Download and install

This section explains the different ways in which Gauge can be installed. The next step is to [install language runners](.install_language_runner.md).

##Windows

#### Using the exe
There are two ways to install Gauge on Windows. Below we show the process using the .exe. The other option is to use chocolatey, as described in the section below.

[Download](http://getgauge.io/get-started) the exe based on your system configuration and run it. Click your way through till you're asked to select your language.

The Windows installer allows you to select the language plugin(s) as part of the the installation process. A language plugin is essentially the language(s) that you would use to write your tests in. Check the box(es) you want to install. You can select more than one language plugin to install.

![Select langauge runner](images/install-lang-runner.jpg)

The windows installer installs Gauge to %PROGRAMFILES% by default. But you can select where you want to install Gauge and complete Gauge installation.

#### Using Chocolatey

Chocolatey is a package manager, very similar to brew for mac. This provides for a simple command line way to install Gauge, if you don't like to click through a user interface.

Install Gauge using Chocolatey with the following command

```
choco install gauge

```

If you're upgrading to a newer version, then use the command:

```
choco upgrade gauge

```
The windows installer installs Gauge to %PROGRAMFILES% by default

##Mac

There are 2 ways you can download an install Gauge. We recommend using brew because it's very simple. However, if you  have limitations then you can use the offline installation instructions.

### Home Brew

The following command installs Gauge.

For this to work, you will need to install [homebrew](http://brew.sh/). If you have brew installed then all you need to run is this command; it will downlaod and install Gauge.  This requires you to be connected to the internet.

```
brew update
brew install gauge
```

If you already have a version of Gauge then the following command will update to the latest version of Gauge.

```
brew update
brew upgrade gauge
```

### Offline Installation

You can [download the Gauge installer](http://getgauge.io/get-started). This is a pkg file, so you can juse click your way through to finish installing Gauge.

To install a language runner and other plugins, go [here]().

##Linux

[Download](http://getgauge.io/get-started) the zip file. Chooe the archive file appropriate for your installation. And run the command below to install Gauge.

```
$ unzip gauge-$VERSION-$OS.$ARCH.zip
~ $ ./install.sh
```

Having trouble with installation? Here is a detailed [installation troubleshooting guide](../troubleshooting/installation.md) that will help.
