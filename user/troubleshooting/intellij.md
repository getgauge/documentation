# Intellij idea plugin

It is advised to use the latest version of [gauge](http://getgauge.io/download) and [Intellij-gauge](https://plugins.jetbrains.com/plugin/7535?pr=idea).

### Intellij Errors

####1. Gauge API error

````
Could not start gauge api: Could not find executable in PATH or GAUGE_ROOT. Gauge is not installed.

````
This can occur because of following reasons :
* Gauge is not installed
* Gauge is installed at custom location and ```custom_install_location/bin``` is not in `PATH`.

To Solve this :
* If gauge is not installed, [install gauge](installation.md)
* If gauge is installed at custom location, add ```custom_install_location/bin``` to `PATH`
* On custom installation location Set `GAUGE_ROOT` to custom_install_location
* Restart Intellij


####2. Steps marked as unimplemented

If steps have implementation code and are still marked as unimplemented
* Ensure that ```src/test/java``` directory is marked as test sources root in the project. Right click on the ```src/test/java``` directory and select **Mark Directory as -> Test sources root**
* Restart Intellij or close and reopen the project.

