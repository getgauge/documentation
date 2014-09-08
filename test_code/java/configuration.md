# Configuration

Java Specific configuration changes can be made in the ***env/default/java.properties*** file.

The following configurations can be made:

### gauge_java_home
Specify an alternate Java home if you want to use a custom version.


###gauge_custom _build_path
Use this property if you need to override the build path for the project.
````
Note: IntelliJ and Eclipse out directory will be usually auto-detected.
````


### gauge_additional_libs
* Specify the directory where additional libraries are kept.
* You can specify multiple directory names separated with a comma **','**
* ***libs*** directory in the gauge project is added by default.


### gauge_jvm_args
Specify the JVM arguments passed to java while launching.



