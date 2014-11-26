# Execution

To execute a spec or set of specs in a directory use the following command.

```
gauge specs/login_test.spec
```
or
```
gauge specs/
```

This will give a colored console output with details of the execution as well an execution summary.
## Single scenario execution

A single scenario of a specification can be executed by specifying the index of that scenario in the spec (index starts from 0). To execute a single scenario in a spec use the following command.

````
gauge specs/login_test.spec:3
````
or
````
gauge login_test.spec:3
````

This executes only the scenario with index 3 in login_test.spec.

Multiple scenarios can be executed selectively as follows :

````
gauge specs/helloworld.spec:4 specs/helloworld.spec:7
````
These scenarios can also belong to different specifications.

## Errors during execution

### 1. Parse error in a spec file:

This occurs if the spec file doesn't follow the expected [specifications](../specifications/README.md) syntax or parameters could not be resolved.

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
