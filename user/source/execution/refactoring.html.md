---
title: Refactoring
---
# Refactoring

## Rephrase steps

Gauge allows you to rephrase a step across the project. To rephrase a step run:

```
$ gauge --refactor "old step <name>" "new step name"
```

Here `<` and `>` are used to denote parameters in the step. __Parameters can be added, removed or changed while rephrasing.__

This will change all spec files and code files (for language plugins that support refactoring).

## Example

Let's say we have the following steps in our `spec` file:

```
* create user "john" with id "123"
* create user "mark" with id "345"
```

Now, if we now need to add an additional parameter, say `last name`, to this step we can run the command:

```
$ gauge --refactor "create user <name> with id <id>" "create user <name> with <id> and last name <watson>"
```

This will change all spec files to reflect the change.

````
* create user "john" with id "123" and last name "watson"
* create user "mark" with id "345" and last name "watson"
````
