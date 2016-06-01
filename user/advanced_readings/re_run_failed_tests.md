# Re-run failed tests

Gauge provides you the ability to re-run only the scenarios which failed in previous execution. Failed scenarios can be run using the `--failed` flag of Gauge.

Say you run `gauge specs` and 3 scenarios failed, you can run re-run only failed scenarios instead of executing all scenarios by following command.

```
gauge --failed
```

This command will even set the flags which you had provided in your previous run. For example, if you had executed command as

```
gauge --env="chrome" --verbose specs
```

and 3 scenarios failed in this run, the `gauge --failed` command sets the `--env` and `--verbose` flags to corresponding values and executes only the 3 failed scenarios.
In this case `gauge --failed` is equivalent to command
```
gauge --env="chrome" --verbose specs <path_to_failed_scenarios>
```
