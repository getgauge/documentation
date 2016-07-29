# Install Gauge on Linux

## On Debian, Ubuntu

### Setup

Add Gauge's GPG key:

```
$ sudo apt-key adv --keyserver hkp://pool.sks-keyservers.net --recv-keys 023EDB0B
```

#### Stable

For stable releases, run this command to add URL to repository list:

```
$ echo deb https://dl.bintray.com/gauge/gauge-deb stable main | sudo tee -a /etc/apt/sources.list
```

#### Nightly

Nightly releases are latest development snapshots of Gauge. They have the latest features being developed, but are unstable. If you want to try out Gauge nightly, do this:

```
$ echo deb https://dl.bintray.com/gauge/gauge-deb nightly main | sudo tee -a /etc/apt/sources.list
```

### Install

```
$ sudo apt-get update
$ sudo apt-get install gauge
```

To set up necessary environment variables and download basic reporting plugins, run this command as a regular user to complete installation:

```
$ gauge_setup
```

## On RHEL, Fedora, CentOS

### Setup

#### Stable

For stable releases, create file `/etc/yum.repos.d/gauge-stable.repo` with the following content:

```
[gauge-stable]
name=gauge-stable
baseurl=http://dl.bintray.com/gauge/gauge-rpm/gauge-stable
gpgcheck=0
enabled=1
```

You can use this command to do it in one step:

```
$ echo -e "[gauge-stable]\nname=gauge-stable\nbaseurl=http://dl.bintray.com/gauge/gauge-rpm/gauge-stable\ngpgcheck=0\nenabled=1" | sudo tee /etc/yum.repos.d/gauge-stable.repo
```

#### Nightly

Note: Nightly releases are latest development snapshots and can be unstable.

For nightly releases, `create /etc/yum.repos.d/gauge-nightly.repo` with the following content:

```
[gauge-nightly]
name=gauge-nightly
baseurl=http://dl.bintray.com/gauge/gauge-rpm/gauge-nightly
gpgcheck=0
enabled=1
```

You can use this command to do it in one step:

```
$ echo -e "[gauge-nightly]\nname=gauge-nightly\nbaseurl=http://dl.bintray.com/gauge/gauge-rpm/gauge-nightly\ngpgcheck=0\nenabled=1" | sudo tee /etc/yum.repos.d/gauge-nightly.repo
```

### Install

Install on Fedora:

```
$ sudo dnf install gauge
```

Install on CentOS/RHEL:

```
$ sudo yum install gauge
```

To set up necessary environment variables and download basic reporting plugins, run this command as a regular user to complete installation:

```
$ gauge_setup
```

## Install manually

[Download](http://getgauge.io/get-started) the zip file. Choose the archive file appropriate for your installation. And run the command below to install Gauge.

```
$ unzip gauge-$VERSION-$OS.$ARCH.zip
$ ./install.sh
```

Having trouble with installation? Here is a detailed [installation troubleshooting guide](../../troubleshooting/installation.md) that can help you.
