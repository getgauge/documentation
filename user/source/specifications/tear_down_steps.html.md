---
title: Tear Down Steps
---

# Tear Down Steps

__Tear Down Steps__ are the steps defined in a spec  after the last scenario.
They allow you to specify a set of clean-up steps after every execution of scenario in a spec. They are used to perform a tear down function.

* Any regular [step](steps.md) can be used as a tear down step.
* Tear down steps are executed after every scenario in the spec.

## Syntax

`___`: Three or more consecutive underscores will indicate the start of tear down. Steps that are written in tear down(after three or more consecutive underscores) will be considered as tear down steps.
```
___
* Tear down step 1
* Tear down step 2
* Tear down step 3
```

## Example

```
Delete project
==============

* Sign up for user "mike"
* Log in as "mike"

Delete single project
---------------------
* Delete the "example" project
* Ensure "example" project has been deleted

Delete multiple projects
------------------------
* Delete all the projects in the list
* Ensure project list is empty

____________________
These are teardown steps

* Logout user "mike"
* Delete user "mike"
```

In the above example spec, the tear down steps are ***Logout user "mike"*** and ***Delete user "mike"***, they are defined after three or more consecutive underscores.

The spec execution flow would be:

1. Context steps execution
2. `Delete single project` scenario execution
3. Tear down steps execution
4. Context steps execution
5. `Delete multiple projects` scenario execution
6. Tear down steps execution
