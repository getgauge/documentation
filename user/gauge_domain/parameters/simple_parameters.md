# Simple parameters
------
They are values passed into the steps in double quotes.

```
* Create a “gauge-java” project
* Write “100” line specification
```
>Note
Renaming the parameter will not rename the usages inside the method.

>By design, the renamed parameter is considered as a new parameter. Therefore the usage of the old parameter(if any) has to be fixed manually to resolve the corresponding compilation issue.
