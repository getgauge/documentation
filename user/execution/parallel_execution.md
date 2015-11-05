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

````
The above two commands will execute the same group of specifications.

## Run your test suite with lazy assignment of tests
This features allows you to dynamically allocate your specs to streams during execution instead of at the start of execution.

This allows Gauge to optimise the resources on your agent/execution environment. This is useful because some specs may take much longer than other, either because of the number of scenarios in them or the nature of the feature under test

The following command will assign tests lazily across the specified number of streams:

````
gauge -n=4 --strategy="lazy" specs
````
Say you have 100 tests, which you have chosen to run across 4 streams/cores; lazy assignment will dynamically, during execution, assign the next spec in line to the stream that has completed it's previous execution and is waiting for more work.

Eager assignment of tests is the default behaviour. In this case, the 100 tests are distributed before execution, thus making them an equal number based distribution.

*Note:* The 'lazy' assignment strategy only works when you do NOT use the -g flag. This is because grouping is dependent on allocation of tests before the start of execution. Using this in conjunction with a lazy strategy will have no impact on your test suite execution.
