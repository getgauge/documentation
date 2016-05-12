# Intellij idea plugin

It is advised to use the latest version of [gauge](http://getgauge.io/download) and [Intellij-gauge](https://plugins.jetbrains.com/plugin/7535?pr=idea).

### Intellij idea Errors

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
* Ensure the project is compiling. Press ctrl/cmd + F9 to build project or select `Build -> Make project`.
* Ensure ```Module SDK``` is set to a valid JDK under ```Module settings```.
* Restart Intellij or close and reopen the project.

#### Check dependencies
##### For a gauge maven project
* The gauge-java dependency should be added in the pom.xml
* Enable auto-import for the project. Under ```File > Settings > Maven > Importing```, mark the checkbox  `Import Maven projects automatically`.

##### For a simple gauge java project
* Under `Project Settings -> Modules` select the gauge module. Under the `dependencies` tab should be `gauge-lib` and `project-lib`.
* If not present restart Intellij or close and re-open project. They should be added

####3. Project Build failing with compilation error but the Java Files do not mark any errors.
* The project compilation fails however the java files do not mark any errors in the file.
* This is a specific issue with Java <= 1.7 on Windows.
* To resolve set **-Duser.home=USER_HOME** in the **IDEA_INSTALLATION\bin\idea.exe.vmoptions** file.

````
-Duser.home=C:\\Users\\johnk
````
* See the [Intellij idea forum post](https://devnet.jetbrains.com/message/5545889#5545889) for more details

