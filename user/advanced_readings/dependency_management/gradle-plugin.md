# Gauge Gradle Plugin

Use the gauge-gradle-plugin to execute specifications in your [Gauge](http://getgauge.io) java project and manage dependencies using [Gradle](http://gradle.org//).

### Using plugin in project

Apply plugin ***gauge*** and add classpath to your ***build.gradle***. Here is a sample gradle file,

````groovy
apply plugin: 'java'
apply plugin: 'gauge'
apply plugin: 'application'

group = "my-gauge-tests"
version = "1.0.0"

description = "My Gauge Tests"

buildscript {
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath 'com.thoughtworks.gauge.gradle:gauge-gradle-plugin:+'
    }
}

repositories {
    mavenCentral()
}

dependencies {
}

// configure gauge task here (optional)
gauge {
    specsDir = 'specs'
    inParallel = true
    nodes = 2
    env = 'dev'
    tags = 'tag1'
    additionalFlags = '--verbose'
}

````

The plugin is also available at [Gradle Plugin Portal](https://plugins.gradle.org/). Find more details [here](https://plugins.gradle.org/plugin/com.thoughtworks.gauge)..

### Executing specs using gradle
To execute gauge specs,

````groovy
gradle gauge
````

#### Execute specs in parallel
```groovy
gradle gauge -PinParallel=true -PspecsDir=specs
```
#### Execute specs by tags
```groovy
gradle gauge -Ptags="!in-progress" -PspecsDir=specs
```
#### Specifying execution environment
```groovy
gradle gauge -Penv="dev" -PspecsDir=specs
```

Note : Pass specsDir parameter as the last one.

### All additional Properties
The following plugin properties can be additionally set:

|Property name|Usage|Description|
|-------------|-----|-----------|
|specsDir| -PspecsDir=specs| Gauge specs directory path. Required for executing specs|
|tags    | -Ptags="tag1 & tag2" |Filter specs by specified tags expression|
|inParallel| -PinParallel=true | Execute specs in parallel|
|nodes    | -Pnodes=3 | Number of parallel execution streams. Use with ```parallel```|
|env      | -Penv=qa  | gauge env to run against  |
|additionalFlags| -PadditionalFlags="--verbose" | Add additional gauge flags to execution|


See gauge's [command line interface](../../cli/README.md) for list of all flags that be used with **-PadditionalFlags** option.

### Adding/configuring custom Gauge tasks
It is possible to define new custom Gauge tasks by extending `GaugePlugin` class. It can be used to create/configure tasks specific for different environments. For example,

````groovy
task gaugeDev(type: GaugeTask) {
    doFirst {
        gauge {
            specsDir = 'specs'
            inParallel = true
            nodes = 2
            env = 'dev'
            additionalFlags = '--verbose'
        }
    }
}

task gaugeTest(type: GaugeTask) {
    doFirst {
        gauge {
            specsDir = 'specs'
            inParallel = true
            nodes = 4
            env = 'test'
            additionalFlags = '--verbose'
        }
    }
}
````
