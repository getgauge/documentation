# Parallel Execution
Specs can be executed in parallel to run the tests faster and distribute the load.

This can be done by the command:

````
gauge --parallel specs
````
This creates a number of execution streams depending on the number of cores of the machine and distribute the load among workers.

The number of parallel execution streams can be specified by `--n` flag.

Example:
````
gauge --parallel --n=4 specs
````
This creates four parallel execution streams.
