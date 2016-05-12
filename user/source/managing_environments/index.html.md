---
title: Managing Environments
---

# Managing Environments

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
   3. Project environment present in the `env/default` dir (if present)
   3. Gauge default env variable values, as below

|Property | Value|
|------------|---------|
| gauge_reports_dir | reports |
| overwrite_reports  | true       |
|screenshot_on_failure|true|
|logs_directory|logs|

Gauge loads the enviroment variables as below.

  *  When Gauge starts, the environment passed by the user in the `--env` flag will be loaded. If this flag is not passed by the user, `default` environment will be loaded.
  * Gauge will then load the `default` environment. Only the values which are not yet set will be loaded. This step won't overwrite the variables which are set in step 1.
  * Finally, Gauge will load the environment variables which are not yet set, as per the table above.
  * These values can be overwritten by explicitly setting the respective OS environment variables.
  * If the environment mentioned in the `--env` flag is not found in the project, Gauge will end with a non-zero exit code.
  * Gauge project doesn't need to have a `default` env since Gauge will use the above values as default. User can still set the `default` env to either overwrite or add new env variables, but doesn't want to pass the `--env` flag.


### Examples

  * User executes `gauge specs`
    * If `<project_root>/env/default` is **not** present, Gauge will set the default env variables with values mentioned in the table above.
    * If `<project_root>/env/default` is present, Gauge will set the env variables mentioned in the `default` environment. It will then set any env variable (which is not already set) as per the table above.
  * User executes `gauge --env=java_ci specs`
     * If `<project_root>/env/java_ci` is **not** present, Gauge will end with a non-zero exit code.
     * If `<project_root>/env/java_ci` is present, Gauge will set the env variables mentioned in the `java_ci` environment. It will then load other variables from the `default` environment which are not already set. Finally, it will the set the env vars with values mentioned in the table above (if they are not already set).
  * User executes `gauge_reports_dir=newReportsDir gauge specs` or user explicitly sets `gauge_reports_dir=newReportsDir` in shell and then runs `gauge specs`
    * Gauge will set all the default env variables from `env/default` directory and then from the above table, except for the variable `gauge_reports_dir`. This variable's value will still continue to be `newReportsDir`.
  * User executes `gauge_reports_dir=newReportsDir gauge --env=java_ci specs` or user explicitly sets `gauge_reports_dir=newReportsDir` in shell and then runs `gauge --env=java_ci specs`
    * Gauge will set the env variables mentioned in the `java_ci` environment. It will then load other variables from the `default` environment which are not already set. Finally, it will the set the env vars with values mentioned in the table above (if they are not already set). However variable `gauge_reports_dir`, which is explicitly set in the shell will not be overwritten. This variable's value will still continue to be `newReportsDir`.

