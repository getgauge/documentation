# Configuration

* [Java](#java)
* [CSharp](#csharp)
* [Ruby](#ruby)

## Java

Java Specific configuration changes can be made in the ***env/default/java.properties*** file.

### gauge_java_home
Specify an alternate Java home if you want to use a custom version.

Example:
```
gauge_java_home = PATH_TO_JAVA_HOME
```


###gauge_custom _build_path

````
Note: IntelliJ and Eclipse out directory will be usually auto-detected.
````

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

### gauge_clear_state_level
Specify the level at which cached objects should get removed while execution.

Possible values for this property are `suite`,`spec` and `scenario`. By default, Gauge clears state at scenario level.

Example:
```
gauge_clear_state_level = spec
```
This clears the objects after the execution of each specification, so that new objects are created for next execution.

## CSharp

## Ruby

