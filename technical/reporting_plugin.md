# Reporting Plugin
Gauge will send messages to Reporting Plugin at the end of execution and Reporting Plugin will recieve the message and perform action.

#### Get the result after ececution

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |`SuiteExecutionResult`|Structure representing the result of entire Suite execution.|   -     |


#### ShutDown after Execution.

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |`KillProcessRequest`|Tells the plugin to shutdown.|   -     |
