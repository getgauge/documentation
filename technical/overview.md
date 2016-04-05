# Overview

## Repositories

* [Gauge](https://github.com/getgauge)

    This the core repository. The Binaries generated are
        gauge/gauge.exe
        gauge_screenshot/gauge_screenshot.exe


* [Gauge-proto](https://github.com/getgauge/gauge-proto)

This is a repository which contains the  [google protobuff](https://github.com/google/protobuf) messages. An IDL that defines contracts between plugins and gauge core.

Several components uses this repository as a sub module. It takes gauge_proto compiles it into the specific language.

* [Common](https://github.com/getgauge/common)

    Shared utility code between components written in Golang.

* Language runners

    Enables gauge support for a specific language. Below is a list of all currently supported language runners.
    1. [Java](https://github.com/getgauge/gauge-java)
    2. [C#](https://github.com/getgauge/gauge-csharp)
    3. [Ruby](https://github.com/getgauge/gauge-ruby)
    4. [Javascript](https://github.com/getgauge-contrib/gauge-js)
    5. [Python](https://github.com/kashishm/gauge-python)

            Javascript and Python are community contributed language runners.


* Reporting plugins

    The summary of the test result can be viewed in the following formats
    1. [Html](https://github.com/getgauge/html-report)
    2. [Xml](https://github.com/getgauge/xml-report)

* IDE Plugins

    Enables gauge support in IDE. Below is the list of all supported IDEs.
    1. [IntelliJ](https://github.com/getgauge/Intellij-Plugin)
    2. [Visual Studio C#](https://github.com/getgauge/gauge-visualstudio)
    3. [Eclipse](https://github.com/getgauge/gauge-eclipse)

