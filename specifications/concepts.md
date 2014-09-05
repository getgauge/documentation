# Concepts

Concepts provide the ability to combine multiple steps into a single step. It provides a higher level abstraction of a business intent by combining steps.

They are defined in **.cpt** format files in the **specs** directory in the project. They can be inside nested directories inside the **specs** dircetory.

* Concepts are used inside spec just like any other step. The appropriate parameters are passed to them.
* On execution all teh steps under the concepts are executred in the defined order.

````
Note: A single .cpt file can contain multiple concept definitions.
````

## Defining a concept
Create a **.cpt** file under specs directory with the concept definition.

The concept definition contians the 2 parts:

### 1. Concept header
The concept header defines the name of the conecpt and the parameters that it takes. It is written in the markdown **H1** format.

* All parameters are defined in angular brackets `< >`.
* A concept definition must have a concept header.

Example:
````
# concept name with <param0> and <param1>
````
### 2. Steps
The concept header is followed by the steps that are used inside the concept. They are defined in the usual [step](steps.md) structure.

* All the parameters used from the concept header will be in `< >` brackets.
* Fixed static parameter values are writen in quotes `" "`.

###Example of Concept definition
````
# login as user <username> and create project <project_name>
* login as user <username> and "password"
* navigate to project page
* create a project <project_name>
````

In the above example
* The first line is the concept header
* The following 3 steps are abstracted into the concept




