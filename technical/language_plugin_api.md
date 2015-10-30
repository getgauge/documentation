# Language Plugin

### Responsibilities & Request/Response:

 - Execute hooks (when requested to)

|Request                    |Purpose                      |Response                 |
     |---------------------------|-----------------------------|-------------------------|
     |`ExecutionStarting`        |Execute `BeforeSuite` Hooks   |`ExecutionStatusResponse`|
     |`ExecutionEnding`          |Execute `AfterSuite` Hooks    |`ExecutionStatusResponse`|
     |`SpecExecutionStarting`    |Execute `BeforeSpec` Hooks    |`ExecutionStatusResponse`|
     |`SpecExecutionEnding`      |Execute `AfterSpec` Hooks     |`ExecutionStatusResponse`|
     |`ScenarioExecutionStarting`|Execute `BeforeScenario` Hooks|`ExecutionStatusResponse`|
     |`ScenarioExecutionEnding`  |Execute `AfterScenario` Hooks |`ExecutionStatusResponse`|
     |`StepExecutionStarting`    |Execute `BeforeStep` Hooks    |`ExecutionStatusResponse`|
     |`StepExecutionEnding`      |Execute `AfterStep` Hooks     |`ExecutionStatusResponse`|



 - Execute method corresponding to a step

|Request      |Purpose                               |Response                 |
     |-------------|--------------------------------------|-------------------------|
     |`ExecuteStep`|Execute method corresponding to a step|`ExecutionStatusResponse`|

 - Refactor method signatures with changes defined by Gauge

|Request      |Purpose                               |Response                 |
     |-------------|--------------------------------------|-------------------------|
     |`StepNameRequest`|Request for implementaion binding details on a Single Step.|`StepNameResponse`|
     |`RefactorRequest`|Refactors the method signature and the step text|`RefactorResponse`|

 - Initialize a project with template
 - Read/write to datastore (currently, datastore is intra process, hence parallel processes cannot use a shared datastore)

|Request                |Purpose                            |Response                 |
     |-----------------------|-----------------------------------|-------------------------|
     |`SuiteDataStoreInit`   |initialize Suite level DataStore.  |`ExecutionStatusResponse`|
     |`SpecDataStoreInit`    |initialize Spec level DataStore    |`ExecutionStatusResponse`|
     |`ScenarioDataStoreInit`|initialize Scenario level DataStore|`ExecutionStatusResponse`|
***All Data Stores should be cleared at their respective levels.***

 - Holds reference to Step text and Method / execution block.

|Request              |Purpose                                                                |Response              |
     |---------------------|-----------------------------------------------------------------------|----------------------|
     |`StepNamesRequest`   |Get all the step names present in the implementaion side.              |`StepNamesResponse`   |
     |`StepValidateRequest`|To check if there is an implementation defined for the given Step Text.|`StepValidateResponse`|

- ShutDown after Execution.

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |`KillProcessRequest`|Tells the runner to shutdown.|        ||
