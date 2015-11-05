# Project Structure

On initialization of a gauge project for a particular language a project skeleton is created with the following files

##1. Common Gauge files

````
├── env
│  └── default
│     └── default.properties
├── manifest.json
├── specs
   └── hello_world.spec
````
### Env Directory
The env directory contains multiple environment specific directories. Each directory has  [.property files](https://en.wikipedia.org/wiki/.properties) which define the environment variables set during execution for that specific environment.

A **env/default** directory is created on project initialization which contains the default environment variables set during execution.

Learn more about [managing environments](../managing_environments/README.md).

### Specs Directory

The specs directory contains all the [spec](../specifications/README.md) files for the project. They are the business layer specifications written in simple markdown format.

A simple example spec (**hello_world.spec**)  is created in the specs directory to better understand the format of specifications.

Learn more about [specifications](../specifications/README.md)

### Manifest file
The **manifest.json** contains gauge specific configurations which includes the information of plugins required in the project.

After project initialization, the `manifest.json` will have the following content.

```
{
  "Language": {{ LANGAUGE }},
  "Plugins": [
    "html-report"
  ]
}
```

* **language** : Programming language used for the test code. Gauge uses the corresponding language runner for executing the specs.

* **Plugins** : The gauge plugins used for the project. Some plugins are used by default on each gauge project. The plugins can be added to project by running the following command :
    ```
    gauge --add-plugin {{PLUGIN_NAME}}

    ```
    Example :
    ```
    gauge --add-plugin xml-report
    ```


After running the above command, the manifest.json would have the following content:
```
{
  "Language": {{ LANGAUGE }},
  "Plugins": [
    "html-report",
    "xml-report"
  ]
}
```

##2. Language specific files

Along with common gauge files certain language specific files and directories are created during [project creation](creating_a_gauge_project.md). You can find the more details about language specific files created under each [programming language](../test_code/README.md).
