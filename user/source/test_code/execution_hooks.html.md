---
title: Execution hooks
---

# Execution hooks 

The [execution hooks](../execution/execution_hooks.md) can be defined at all levels in java by annotating methods with the following annotations:


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
|BeforeClassSteps (Java only)|`@BeforeClassSteps`|
|AfterClassSteps (Java only)|`@AfterClassSteps`|

### Example

#### Java

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

#### C# #

````csharp
public class ExecutionHooks {

    [BeforeSpec]
    public void LoginUser() 
    {
      // Code for before spec
    }

    [AfterStep]
    public void PerformAfterStep() 
    {
      // Code for after step
    }
}
````

#### Ruby

````ruby
before_spec do
    // Code for before spec
end

after_spec do
    // Code for after step
end
````

> By default, Gauge clears the state after each scenario so that new objects are created for next scenario execution.
You can [configure](configuration.md#gaugeclearstatelevel) to change the level at which Gauge clears cache.

## Filtering Hooks execution based on tags

* You can specify tags for which the execution hooks can run. This will ensure that the hook runs only on scenarios and specifications that have the required tags.

#### Java

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

#### C# #

````csharp
    // A before spec hook that runs when tag1 and tag2 is present in the current scenario and spec.
    [BeforeSpec("tag1, tag2")]
    public void LoginUser() 
    {
       // Code for before scenario
    }


    // A after step hook runs when tag1 or tag2 is present in the current scenario and spec.
    // Default tagAggregation value is Operator.AND.
    [AfterStep("tag1", "tag2")
    [TagAggregationBehagiour(TagAggregation.OR)]
    public void PerformAfterStep() 
    {
       // Code for after step
    }

````

#### Ruby

````ruby
    # A before spec hook that runs when tag1 and tag2 is present in the current scenario and spec.
    before_spec({tags: ['tag2', 'tag1']}) do
       // Code for before scenario
    end

    // A after step hook runs when tag1 or tag2 is present in the current scenario and spec.
    // Default tagAggregation value is Operator.AND.
    after_spec({tags: ['tag2', 'tag1'], operator: 'OR'}) do
       // Code for after step
    end

````

> Note: Tags cannot be specified on @BeforeSuite and @AfterSuite hooks


## Current Execution Context in the Hook

* To get additional information about the **current specification, scenario and step** executing, an additional **ExecutionContext** parameter can be added to the hooks method.

### Java

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

### C# #

> This feature is not yet supported in Gauge-CSharp. Please refer to [this issue](https://github.com/getgauge/gauge-csharp/issues/53) for updates.

### Ruby

````ruby
before_spec do |execution_info|
    puts execution_info.inspect
end

after_spec do |execution_info|
    puts execution_info.inspect
end
````
