# Environments

Environment specific [variables](https://en.wikipedia.org/wiki/Environment_variable) can be managed using property files. The [property files](https://en.wikipedia.org/wiki/.properties) have set of key value pairs which are set as environment variables during execution.

The env directory structure for a `java` project:
````
├── env
  └── default
     ├── default.properties
     └── java.properties
```

Custom properties can be added to an existing property files or in a newly created one.

##Creating new environment

To create an environment called `ci`:

* Create a directory called `ci` in `env` directory
* Add property files (e.g. `user.properties`)


```
├── env
   ├── ci
   │  └── user.properties
   └── default
      ├── default.properties
      └── java.properties
```

##Executing with environment

The environment is specified using the `env` flag. For example if `ci` environment is used during execution
```
gauge --env ci specs/
```
On execution
* The properties from the `default` directory are first set.
* The properties from the specified environment `ci`  are then set.
* For common values,  environment specific variables from `ci` override the default values.

For e.g.

```
gauge_custom_build_path: out/production
```

present in **env/default/java.properties** is overridden by the

```
gauge_custom_build_path: target
```

present in the **env/ci/java.properties**
