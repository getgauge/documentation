# XML-Report

This plugin generates Junit style XML reports of the test execution result provided by Gauge.

Build from Source
-----------------

### Requirements
* [Golang](http://golang.org/)

### Compiling

````
go run build/make.go
````

For cross platform compilation

````
go run build/make.go --all-platforms
````

### Installing
After compilation

````
go run build/make.go --install
````

Installing to a CUSTOM_LOCATION

````
go run build/make.go --install --plugin-prefix CUSTOM_LOCATION
````

### Creating distributable

Note: Run after compiling

````
go run build/make.go --distro
````

For distributable across platforms os, windows and linux for bith x86 and x86_64

````
go run build/make.go --distro --all-platforms
````

New distribution details need to be updated in the xml-report-install.json file in  [gauge plugin repository](https://github.com/getgauge/gauge-repository) for a new verison update.
