# Gauge IntelliJ idea plugin Features

* [Project Creation](#1-project-creation)
* [Auto completion](#2-auto-completion)
* [Quick Fix for unimplemented steps](#3-step-quick-fix)
* [Formatting](#4-formatting)
* [Execution and Debugging](#5-execution)
* [Parallel Execution](#52-parallel-execution)
* [Run configurations](#6-run-configuration)
* [Rephrase steps](#7-rephrase-steps)
* [Find Usages](#8-find-usages)
* Navigation from step to implementation (cmd/ctrl + b)

###1. Project Creation
####1.1 Creating a simple java project
 * File -> New Project.
 * Choose 'Gauge'
 * Choose the project location and java sdk
 * Finish

````
Note: If gauge-java is not installed, it will download it for the first time.
````

![creation](intellij-screenshots/creation/creation.gif "project creation")

#### 1.2 Creating a maven project using gauge-maven-plugin

* File -> New Project
* Choose `Maven`
* Select `Create from Archetype`
* Select the gauge archetype - ```com.thoughtworks.gauge.maven```
* If the ```com.thoughtworks.gauge.maven``` archetype is not added select ```Add Archetype```
    * Enter GroupId: com.thoughtworks.gauge.maven
    * Enter ArtifactId: gauge-maven-plugin
    * Enter Version: 1.0.3 or the [latest version number](https://github.com/getgauge/gauge-maven-plugin/releases)
![maven add archetype](intellij-screenshots/creation/maven_add_archetype.png "maven add archetype")


* Enter the ```groupId``` and `artifactId` for your project.
* Enter `Project Name` and finish
* The project will be created in batch mode, watch the console for progress.
* After project creation ```close and re-open the project``` to enable auto-complete features.
* Enable ```auto-import``` for the project. Under ```File > Settings > Maven > Importing```, mark the checkbox  `Import Maven projects automatically`.

 See [gauge-maven-plugin](../dependency_management/maven-plugin.md) for more details on using the gauge maven plugin.



###2. Auto Completion
Steps present in the current project can be listed by invoking the auto completion pop up `ctrl+space` after the '*'. After choosing a step, it gets inserted with parameters highlighted, you can press `tab` to cycle between the parameters and edit them.

![creation](intellij-screenshots/auto completion/completion.gif "step completion")


###3. Step Quick Fix

If you have an unimplemented step in the spec file, it will be annotated saying 'undefined step'. `alt+enter` can be pressed to open the quick fix pop up. The destination of the implementation can be chosen, either a new class or from a list of existing classes. It will then generate the step with required annotation and parameters.

![step quick fix](intellij-screenshots/quick fix/fix.gif "step quick fix")

###4. Formatting

A spec file can be formatted by pressing `ctrl+alt+shift l` (in windows and linux) and `cmd+alt+shift l`(in mac).

This formats all the elements of current spec including indentation of tables and steps.

###5. Execution

* Specs can be executed by `right click -> Run spec`.
* Execute all specs inside a directory by `right click -> Run specifications`

####5.1. Single Scenario Execution

A single scenario can be executed by doing a right click on the scenario which should be executed and choosing the scenario.
`right click -> run -> Scenario Name`

````
Note: If the right click is done in context other than that of scenario, by default, first scenario will be executed.
````

![scenario execution](intellij-screenshots/execution/scenario.gif "scenario execution")

####5.2. Parallel Execution
To run multiple specifications in parallel
* Right click on the ```specs``` directory and select ```Create Specifications``` option.
* In the new Run configuration select ```In Parallel``` options. This will distribute specs execution based on number of cores the machine has.
* You can also specify the ```Number of parallel execution streams```. This is optional
````
Caution: Select parallel nodes based on current systems performance.
For example on a 2 core machine select upto 4 parallel streams.
A very large number may affect performance.
````
* Select ```ok```. Now you can run this new configuration for parallel execution of specs.


####5.3. Debugging
Debugging can be performed the same way spec execution works.
* Right click on a specification or specs directory -> Debug. Execution will halt on marked [breakpoints](https://www.jetbrains.com/idea/help/breakpoints.html).

###6. Run Configuration

You can edit the run configuration to make changes to:
* The scenario or spec file to be executed
* The environment to run against
* Add a tag filter to the execution
* Choose the number of parallel streams


![run configuration](intellij-screenshots/execution/config.gif "run configuration")

###7. Rephrase Steps
* Press **Shift+F6** on a step to rephrase it.
* The parameters will be in ```< >``` in the rephrase dialog. They can be reordered,removed or new parameters can be added.
* The rephrase change will reflect across **all the specs** in the project.

###8. Find Usages
* Press **cmd/ctrl+alt+F7** on step/concept to see the usages.

![find usages](intellij-screenshots/find_usages/find_usages.gif "find usages")


