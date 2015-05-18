# Execution hooks Java

The [execution hooks](../../execution/execution_hooks.md) can be defined at all levels in java by annotating methods with the following annotations:


| Execution Hook | annotation  |
|----------------| ------------|
| Before Suite   | `@BeforeSuite`|
| After Suite    | `@AfterSuite`|
| Before Specification   | `@BeforeSpec`|
| After Specification   | `@AfterSpec`|
| Before Scenario | `@BeforeScenario`|
| After Scenario   | `@AfterScenario`|
| Before Step | `@BeforeStep` |
|After Step| `@AfterStep`|
|BeforeClassSteps|`@BeforeClassSteps`|
|AfterClassSteps|`@AfterClassSteps`|

###Example
````java
public class ExecutionHooks {

    @BeforeSpec
    public void loginUser() {
    // Code for before spec
    }

    @AfterStep
    public void performAfterStep() {
    // Code for after step
    }
}
````
Add parameter **`SpecificationInfo`** to the hooks for getting information about the current running **`specification`**.
###Example
````java
public class ExecutionHooks {

    @BeforeScenario
    public void loginUser(SpecificationInfo info) {
    // Code for before scenario
    }

    @AfterStep
    public void performAfterStep(SpecificationInfo info) {
    // Code for after step
    }
}
````


By default, Gauge cleares the state after each scenario so that new objects are created for next scenario execution.

You can [configure](configuration.md) to change the level at which Gauge cleares cache.

