# Setting up Gauge

####Prerequisite Tools
* [golang](http://golang.org/)

Ensure golang dev environment is setup. [GOPATH](https://golang.org/doc/code.html#GOPATH) environment variable should be set. Add ```$GOPATH/bin``` to ```$PATH```

* Clone the [gauge repository](https://github.com/getgauge/gauge.git) to ```$GOPATH/src/github.com/getgauge/gauge```

* Gauge uses submodules. So issue the following commands before you attempt to build

```
  git submodule init
  git submodule update
```

* Gauge uses [godep](https://github.com/tools/godep) to manage Go dependencies.
Install godep by running

````
go get github.com/tools/godep
````

* To fetch all dependencies, in the gauge repository run

````
go get github.com/tools/godep
godep restore
````

##Building

````
go run build/make.go
````

This will generate gauge executable in the bin directory

## Running Tests

````
go test -v ./...
````
or
````
go run build/make.go --test
````
With Test coverage
````
go run build/make.go --test --coverage
````

## Installing

###Mac OS X and Linux

````
go run build/make.go --install
````

This installs gauge into __/usr/local__ by default.
To install into a custom location use a prefix for installation

````
go run build/make.go --install --prefix CUSTOM_PATH
````

###Windows

````
go run build\make.go --install --prefix CUSTOM_PATH
````

Set environment variable GAUGE_ROOT to the CUSTOM_PATH

## Verifying Installation

To verify the installation, run the following command :

````
gauge -v
````
The above command will print the version of Gauge and plugins installed.
