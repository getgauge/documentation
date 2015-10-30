# Language Plugin
Gauge will send requests to Language Plugin and Language Plugin will perform actions based on request and send the response back.

#### Execute hooks (when requested to)

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

#### Execute method corresponding to a step

|Request      |Purpose                               |Response                 |
     |-------------|--------------------------------------|-------------------------|
     |`ExecuteStep`|Execute method corresponding to a step|`ExecutionStatusResponse`|

#### Refactor method signatures with changes defined by Gauge

|Request      |Purpose                               |Response                 |
     |-------------|--------------------------------------|-------------------------|
     |`StepNameRequest`|Request for step text present in the implementaion for a Step.|`StepNameResponse`|
     |`RefactorRequest`|Refactors the method signature and the step text|`RefactorResponse`|

#### Read/write to datastore

|Request                |Purpose                            |Response                 |
     |-----------------------|-----------------------------------|-------------------------|
     |`SuiteDataStoreInit`   |initialize Suite level DataStore.  |`ExecutionStatusResponse`|
     |`SpecDataStoreInit`    |initialize Spec level DataStore    |`ExecutionStatusResponse`|
     |`ScenarioDataStoreInit`|initialize Scenario level DataStore|`ExecutionStatusResponse`|
*All Data Stores should be cleared at their respective levels.*

#### Holds reference to Step text and Method / execution block.

|Request              |Purpose                                                                |Response              |
     |---------------------|-----------------------------------------------------------------------|----------------------|
     |`StepNamesRequest`   |Get all the step names present in the implementaion side.              |`StepNamesResponse`   |
     |`StepValidateRequest`|To check if there is an implementation defined for the given Step Text.|`StepValidateResponse`|

#### ShutDown after Execution.

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |`KillProcessRequest`|Tells the runner to shutdown.|   -     |

#### Send Response when a unsupported message request is sent.

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |Unknown Request|Request not supported by plugin.|   `UnsupportedMessageResponse`     |

