# CSharp

This project adds C# support for Gauge.

This repository hosts three projects, `Core`, `Lib` and `Runner`.

## Gauge Core

### *Internal Gauge Use Only*
`Core` is a dll that brings in Gauge's API and connection to C#, and also holds the protobuf communication code. It is hosted in [Nuget](https://www.nuget.org/packages/Gauge.CSharp.Core/), and is not listed in the search results. This package is used by `Lib` and `Runner` only.

## Gauge Lib
`Lib` is a dll that brings in Gauge's data types to C#. It is hosted in [Nuget](https://www.nuget.org/packages/Gauge.CSharp.Lib/).

## Gauge CSharp Runner
`Runner` is an executable that is invoked by Gauge Core. The `Runner` acts a bridge between C# test code and Gauge's API.

Read more about `Runner` [here](runner.md).

## Gauge-Proto
This is a submodule of the repository `https://github.com/getgauge/gauge-proto`. This repository holds the `.proto` files that act as contracts between Gauge and the plugins. This submodule needs to be fetched to generate the protobuf classes.

## Developer notes

### Install

#### Runner
Installing the runner is done via Gauge Core. This means that you need Gauge installed as a pre-requisite. You can download Gauge [here](http://getgauge.io/download.html).

Once you have Gauge installed, add the csharp plugin using:

    gauge --install csharp

#### Lib
The Lib is a reference that you can add to your test project, as you would do with any Nuget package.

You will need Gauge and Gauge-csharp plugin installed before installing `Lib`.

### Build

    run.bat build

### Test

    run.bat test

### Package

    run.bat package

All artifacts are genereated in `.\artifacts` folder.

The plugin is zipped to `.\artifacts\gauge-csharp-<version>.zip`
The Lib Nuget package is put at `.\artigacts\Gauge.CSharp.Lib.<version>.nupkg`
The Core Nuget package is put at `.\artigacts\Gauge.CSharp.Core.<version>.nupkg`

### Install the plugin

To install a local version of the plugin:

    run.bat install

### Regenerate the API messages (protocol buffer api changes)

Update the submodule and run the `gen-proto` command:

    git submodule update
    run.bat gen-proto
