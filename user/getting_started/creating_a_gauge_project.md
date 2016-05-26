# Creating a Gauge project
-----
## Command line

To create or initialize a Gauge project use the `gauge --init` command. In an empty directory run

```
$ gauge --init <language>
```

Replace *<language>* with the programming language in which you want to write the test code. See [supported programming languages](../test_code/step_implementations.md).

**Note:** Currently this command creates the Gauge project in the current directory. If you would like to specify a location, then you will have to manually create that dir and navigate to it before you run `--init langauge`

```
$ mkdir <location for Gauge project>
$ cd <location for Gauge project>

$ gauge --init <language>
```

**Examples:**

```
$ gauge --init java
```

```
$ gauge --init csharp
```

```
$ gauge --init ruby
```

This will initialize a gauge project with all the necessary project files.

![init](images/gauge-init.png "init")

````
Note: This will download and install the specific language plugin if it is not currently installed.
````

## IDE
* Create your project using the project creation wizard in your favorite [IDE](http://en.wikipedia.org/wiki/Integrated_development_environment). See [IDE support](../ide_support/README.md) for supported addons.

### Learn more
* [Gauge project structure](project_structure.md)
* [Executing Specs](../execution/README.md)
* [Specifications](../specifications/README.md)
* [Test code](../test_code/step_implementations.md)
* [IDE support](../ide_support/README.md)
