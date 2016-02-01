# Environments

Environment specific [variables](https://en.wikipedia.org/wiki/Environment_variable) can be managed using property files. The [property files](https://en.wikipedia.org/wiki/.properties) have set of key value pairs which are set as environment variables during execution.

The env directory structure for a `java` project:
````
├── env
  └── default
     ├── default.properties
     └── java.properties
````

Custom properties can be added to an existing property files or in a newly created one.

##Creating new environment

To create an environment called `ci`:

* Create a directory called `ci` in `env` directory
* Add property files (e.g. `user.properties`)


```
├── env
   ├── ci
      |── user.properties
   |── default
      ├── default.properties
      └── java.properties
```

##Executing with environment

The environment is specified using the `env` flag. For example if `ci` environment is used during execution
```
gauge --env ci specs
```


##Configuring the Properties

Certain properties can be configured in `env/default/default.properties`, which overrides the default properties and are scoped only to the current project.

These are key value pairs, key being the property name and the value is either the absolute or path relative to project root.

Example:

```
gauge_reports_dir = reports

logs_directory = GaugeLogs
```

##Precedence of Environments

Precedence to the env variable value is given in the below order.
   1. User shell / OS env variable values
   2. Project environment passed in the `--env` flag
   3. Gauge default env variable values, as below

|Property | Value|
|------------|---------|
| gauge_reports_dir | reports |
| overwrite_reports  | true       |
|screenshot_on_failure|true|
|logs_directory|logs|

Gauge loads the enviroment variables as below.

  *  When Gauge starts, the default values will be loaded.
  * Gauge will then load the environment passed by the user in the `--env` flag. If this flag is not passed by the user, `default` environment will be loaded.
  * These values can be overwritten by explicitly setting the respective OS environment variables.
  * If the environment mentioned in the `--env` flag is not found in the project, Gauge will end with a non-zero exit code.
  * Gauge project doesn't need to have a `default` env since Gauge will use the above values as default. User can still set the `default` env to either overwrite or add new env variables, but doesn't want to pass the `--env` flag.


### Examples

  * User executes `gauge specs`
    * If `<project_root>/env/default` is **not** present, Gauge will load the default values mentioned in the table above.
    * If `<project_root>/env/default` is present, Gauge will load the default values and overwrite them with the values mentioned in the `default` environment.
  * User executes `gauge --env=java_ci specs`
     * If `<project_root>/env/java_ci` is **not** present, Gauge will end with a non-zero exit code.
     * If `<project_root>/env/java_ci` is present, Gauge will load the default values and overwrite them with the values mentioned in the `java_ci` environment.
  * User executes `gauge_reports_dir=newReportsDir gauge specs` or user explicitly sets `gauge_reports_dir=newReportsDir` in shell and then runs `gauge specs`
    * Gauge will load all the default values except for the variable `gauge_reports_dir`. This variable's value will still continue to be `newReportsDir`.
  * User executes `gauge_reports_dir=newReportsDir gauge --env=java_ci specs` or user explicitly sets `gauge_reports_dir=newReportsDir` in shell and then runs `gauge --env=java_ci specs`
    * Gauge will load all the default values and overwrite them with the values mentioned in the `java_ci` environment, except for the variable `gauge_reports_dir`. This variable's value will still continue to be `newReportsDir`.

