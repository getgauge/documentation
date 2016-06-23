# Using build tools

You can use Gauge with any of the build tools that you like.

Here are the sample build files for
1. [Maven](#a-namemavenamaven)
2. [Gradle](#a-namegradleagradle)
3. [Ant](#a-nameantaant-task)

For more details, please take a look at the [Dependency Management](dependency_management_plugins/README.md) section.

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

###<a name="Ant"></a>Ant Task
````xml
<target name="specs">
    <exec executable="gauge">
        <arg value="specs"/>
    </exec>
</target>
````
