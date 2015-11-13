# Using build tools

You can use Gauge with any of the build tools that you like.

Here are the sample build files for
1. [Maven](#Maven)
2. [Gradle](#Gradle)
3. [Ant](#Ant)

###<a name="Maven"></a>Maven

````xml
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.foo</groupId>
    <artifactId>my-gauge-tests</artifactId>
    <version>1.0-SNAPSHOT</version>

    <dependencies>
        <dependency>
            <groupId>com.thoughtworks.gauge</groupId>
            <artifactId>gauge-java</artifactId>
            <version>0.2.2</version>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>com.thoughtworks.gauge.maven</groupId>
                <artifactId>gauge-maven-plugin</artifactId>
                <version>1.0.3</version>
                <executions>
                    <execution>
                        <phase>test</phase>
                        <configuration>
                            <specsDir>specs</specsDir>
                        </configuration>
                        <goals>
                            <goal>execute</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>
</project>
````

###<a name="Gradle"></a>Gradle
````groovy
apply plugin: 'java'
apply plugin: 'application'

group = "my-gauge-tests"
version = "1.0.0"

description = "My Gauge Tests"

sourceCompatibility = 1.7
targetCompatibility = 1.7

buildscript {
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath files("$projectDir/libs")
    }
}

repositories {
    mavenCentral()
}

sourceSets {
    main.java.srcDir 'src'
}

dependencies {
    compile(
            'com.thoughtworks.gauge:gauge-java:0.2.2'
    )
}

def executeGauge() {
    exec {
        executable "gauge"
        args "specs"
    }
}

task copyLibs(type: Copy) {
    from configurations.runtime
    into "$projectDir/libs"
}

task gauge(dependsOn: 'copyLibs') {
    doLast  {
        executeGauge()
    }
}
````

###<a name="Ant"></a>Ant Task
````xml
<target name="specs">
    <exec executable="gauge">
        <arg value="specs"/>
    </exec>
</target>
````
