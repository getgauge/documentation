# Contexts

__Contexts__ or __Context steps__ are steps defined in a spec before any scenario.

They allow you to specify a set of conditions that are necessary for executing scenarios in a spec. Context steps can be used to set up data before running scenarios. They can also perform a setup or tear down function. 

* Any regular [step](steps.md) can be used as a context.
* Contexts are executed before every scenario in the spec.

````
Delete project
==============
These are context steps

* User is logged in as "mike"
* Navigate to the project page

Delete single project
---------------------
* Delete the "example" project
* Ensure "example" project has been deleted

Delete multiple projects
------------------------
* Delete all the projects in the list
* Ensure project list is empty
````

In the above example spec the context steps are ***User is logged in as Mike*** and ***Ensure example project has been deleted***, they are defined before any scenario.

These steps are executed before the execution of each scenario ***Delete single project*** and ***Delete multiple projects***.

The spec execution flow would be:
1. Context steps execution
2. ``Delete single project`` scenario execution
3. Context steps execution
3. ``Delete multiple projects`` scenario execution



