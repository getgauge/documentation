# Configuration

Apart from the default configuration, we can create and use other configuration files which have logically grouped configurations.
Here are the configurations generated while using the Java language runner.

## Java

Java Specific configuration changes can be made in the ***env/default/java.properties*** file.

### gauge_java_home
Specify an alternate Java home if you want to use a custom version.

Example:
```
gauge_java_home = PATH_TO_JAVA_HOME
```

###gauge_custom_build_path

```
Note: IntelliJ out directory will be usually auto-detected.
```

Use this property if you need to override the build path for the project.

Example:
```
gauge_custom_build_path = PATH_TO_CUSTOM_BUILDPATH
```

### gauge_additional_libs
* Specify the directory where additional libraries are kept.
* You can specify multiple directory names separated with a comma **','**
* ***libs*** directory in the gauge project is added by default.

Example:
```
gauge_additional_libs = libs/*, PATH_TO_NEW_LIBRARY
```

### gauge_jvm_args
Specify the JVM arguments passed to java while launching.
