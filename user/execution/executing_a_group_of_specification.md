# Executing a group of specification
Specifications can be distributed into groups and `--group` | `-g` flag provides the ability to execute a specific group.

This can be done by the command:

````
gauge -n=4 -g=2 specs

````
This creates 4 groups (provided by -n flag) of specification and selects the 2nd group (provided by -g flag) for execution.

Specifications are sorted by alphabetical order and then distributed into groups, which guarantees that every group will have the same set of specifications, no matter how many times it is being executed.

Example:
````
gauge -n=4 -g=2 specs

gauge -n=4 -g=2 specs
````
The above two commands will execute the same group of specifications.
