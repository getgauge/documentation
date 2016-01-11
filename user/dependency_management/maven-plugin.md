# Gauge-maven-plugin

Use the gauge-maven-plugin to execute specifications in your gauge java project and manage dependencies using [maven](https://maven.apache.org/).

### Creating a new project from archetype

```bash
mvn archetype:generate -DgroupId={projectGroupId}
-DartifactId={projectArtifactId}
-DarchetypeArtifactId=gauge-maven-plugin
-DarchetypeGroupId=com.thoughtworks.gauge.maven

```
Set **{projectGroupId}** and **{projectArtifactId}** based on your project.
See [maven docs](https://maven.apache.org/pom.html#Maven_Coordinates) to understand what groupId and artifactId mean in a maven project.

### Gauge maven project creation in IDE
* [Intellij idea](../ide_support/features.md#12-creating-a-maven-project-using-gauge-maven-plugin)


The generated **pom.xml** in the project will have the** gauge-java** dependency and a **gauge:execute** goal defined in the test phase.

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

### Executing specs using maven
If the execute goal is added for test phase (see above xml) then running maven test phase will also execute gauge specs in the project
````
mvn test
````

####To only run gauge specs,
```
mvn gauge:execute -DspecsDir=specs
```

####Execute specs In parallel
```
mvn gauge:execute -DspecsDir=specs -DinParallel=true
```
####Execute specs by tags
```
mvn gauge:execute -DspecsDir=specs -Dtags="!in-progress"
```
####Specifying execution environment
```
mvn gauge:execute -DspecsDir=specs -Denv="dev"
```
### All additional Properties
The following plugin properties can be additionally set:

|Property name|Usage|Description|
|-------------|-----|-----------|
|specsDir| -DspecsDir=specs| Gauge specs directory path. Required for executing specs|
|tags    | -Dtags="tag1 & tag2" |Filter specs by specified tags expression|
|inParallel| -DinParallel=true | Execute specs in parallel|
|nodes    | -Dnodes=3 | Number of parallel execution streams. Use with ```parallel```|
|env      | -Denv=qa  | gauge env to run against  |
|dir  | -Ddir=. | Set working directory for gauge.  Default is project.basedir. |  
|flags| -Dflags="--verbose" | Add additional gauge flags to execution|


See gauge's [command line interface](../cli/README.md) for list of all flags that be used with **-Dflags** option.


