# Creating a gauge  project

## Command line
To create or initialize a Gauge project use the `gauge --init` command. In an empty directory run

````
gauge --init language
````
Where **language** is the programming language to write test code in. See [supported programming languages](../test_code/README.md).

**Example:**
````
gauge --init java
````
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
* [Test code](../test_code/README.md)
* [IDE support](../ide_support/README.md)
