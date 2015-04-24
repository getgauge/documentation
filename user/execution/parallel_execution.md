# Parallel Execution
Specs can be executed in parallel to run the tests faster and distribute the load.

This can be done by the command:

````
gauge --parallel specs

gauge -p specs
````
This creates a number of execution streams depending on the number of cores of the machine and distribute the load among workers.

The number of parallel execution streams can be specified by `-n` flag.

Example:
````
gauge --parallel -n=4 specs
````
This creates four parallel execution streams.

**Note**:
The number of streams should be specified depending on the machine, beyond which it affects the performance of the system.

## Rerun one execution stream
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
