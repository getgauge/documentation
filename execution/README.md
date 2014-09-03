# Execution

To execute a spec or set of specs in a directory

```
gauge specs/Login Test.spec
```
or
```
gauge specs/
```

## Errors during execution

### Parse error in a spec file:
This is because the spec file doesn't follow the expected syntax or parameters could not be resolved

eg

`
[ParseError] hello_world.spec : line no: 25, Dynamic parameter <product> could not be resolved
`


### Unimplemented steps present in spec file
If the spec file has a step that is not implemented in the runner side, there will be a validation error. So a proper implementaion has to be provided for all the steps and concepts

eg

`login.spec:33: Step implementation not found. login with "user" and "p@ssword"`

### Not able to launch the runner
If the language runtimes haven't been installed(java, ruby) then the runner can't be launched.
