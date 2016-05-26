# Parameters
---
Steps can be defined to take values as parameters so that they can be re-used with different parameter values.

```
* Check "product 1" exists
* Check "product 2" exists
```

The underlying [step implementation](language_features/step_implementations.md) in code must also take the same number of parameters as passed from the step.

The parameters passed into a step can be of different types:
- [Simple parameter](gauge_domain/parameters/simple_parameters.md)
- [Table parameter](gauge_domain/parameters/table_parameters.md)
- [Special parameter](gauge_domain/parameters/special_parameters.md)
