# Execution

To execute a spec or set of specs in a directory use the following command.

```
$ gauge specs/login_test.spec
```
or
```
$ gauge specs/
```
This will give a colored console output with details of the execution as well an execution summary.

## Single scenario execution

A single scenario of a specification can be executed by specifying the line number in the span of that scenario in the spec. To execute a `Admin Login` scenario in the following spec use `gauge specs/login_test.spec:4` command.
```
1>   Configuration
2>   =============
3>
4>   Admin Login
5>   -----------
6>   * User must login as "admin"
7>   * Navigate to the configuration page
```

This executes only the scenario present at line number `4` i.e `Admin Login` in `login_test.spec`. In the above spec, specifying line numbers 4-7 will execute the same scenario because of the span.

Multiple scenarios can be executed selectively as follows :

```
$ gauge specs/helloworld.spec:4 specs/helloworld.spec:7
```

These scenarios can also belong to different specifications.

## Errors during execution

### 1. Parse error in a spec file:

This occurs if the spec file doesn't follow the expected [specifications](../../gauge_terminologies/specifications.md) syntax or parameters could not be resolved.

**Example**

```
[ParseError] hello_world.spec : line no: 25, Dynamic parameter <product> could not be resolved
```



###2. Unimplemented steps present in spec file
If the spec file has a step that does not have an implementation in the projects programming language there will be a validation error.

Appropriate underlying code implementation has to be provided for all the steps in the specs to be executed.

**Example**

````
login.spec:33: Step implementation not found. login with "user" and "p@ssword"
````

###3. Failure to launch the language runner plugin
If the language specific plugin for the project has not been installed then the execution will fail.
