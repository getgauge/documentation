# Creating a gauge  project

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

### Learn more
* [Gauge project structure](project_structure.md)
* [Executing Specs](../execution/README.md)
