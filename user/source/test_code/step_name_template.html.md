---
title: Step Name Template
---

# Step Name Template

The step name template is a simple generic representation for a step. It is used while defining language implementations for steps.


The Step template contains:
* The name of the step as it is used.
* The parameters are replaced by placeholders in angular brackets `< >`. The placeholder value determines what the parameter represents.
* Additional `<table>`  value is added if a step contains a table parameter.

###Examples
| Step Usage | Step name template|
| --    | -- |
| create a user| create a user |
| create a user "prateek"| create a user < username > |
| Verify text < file:test.txt > | Verify text < email text >|


**For Steps with tables:**
````
| Step usage                |       Step name template             |
|---------------------------|--------------------------------------|
| Create following users    |    Create following users <table>    |
|   |user-id|  name   |     |                                      |
|   |  123  | vishnu  |     |                                      |
|   |  456  |navaneeth|     |                                      |

````



