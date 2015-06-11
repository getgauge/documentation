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

By default, Gauge clears the state after each scenario so that new objects are created for next scenario execution.

You can [configure](configuration.md) to change the level at which Gauge clears cache.

## Filtering Hooks execution based on tags
* You can specify tags for which the execution hooks can run. This will ensure that the hook runs only on scenarios and specifications that have the required tags.

````java

    // A before spec hook that runs when tag1 and tag2 is present in the current scenario and spec.
    @BeforeSpec(tags = {"tag1, tag2"})
    public void loginUser() {
       // Code for before scenario
    }


    // A after step hook runs when tag1 or tag2 is present in the current scenario and spec.
    // Default tagAggregation value is Operator.AND.
    @AfterStep(tags = {"tag1", "tag2"}, tagAggregation = Operator.OR)
    public void performAfterStep() {
       // Code for after step
    }

````
```
Note: Tags cannot be specified on @BeforeSuite and @AfterSuite hooks
```


## Current Execution Context in the Hook
* To get additional information about the **current specification, scenario and step** executing, an additional **ExecutionContext** parameter can be added to the hooks method.

````java
public class ExecutionHooks {

    @BeforeScenario
    public void loginUser(ExecutionContext context) {
      String scenarioName = context.getCurrentScenario().getName();
      // Code for before scenario
    }

    @AfterSpec
    public void performAfterSpec(ExecutionContext context) {
      Specification currentSpecification = context.getCurrentSpecification();
      // Code for after step
    }
}
````





