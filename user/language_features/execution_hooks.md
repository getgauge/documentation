# Execution hooks

Test execution hooks can be used to run arbitrary test code as different levels during the test suite execution.

## Before Suite Hook
Executes before the entire suite execution begins i.e, before the execution of all specifications in the project.

## After Suite Hook
Executes after the entire suite execution finishes i.e, after the execution of all specifications in the project.

## Before Specification hook
Executes before every spec executes.

## After Specification hook
Executes after every spec executes.

## Before Scenario hook
Executes before every scenario executes.

## After Specification hook
Executes after every scenario executes.

## Before Step hook
Executes before every step executes.

## After Step hook
Executes after every scenario executes.

## Hooks specific to Java
- BeforeClassSteps (Java only)-`@BeforeClassSteps`
- AfterClassSteps (Java only)-`@AfterClassSteps`

### Example

{% codetabs name="Java", type="java" -%}
public class ExecutionHooks {
    @BeforeSuite
    public void BeforeSuite() {
      // Code for before suite
    }

    @AfterSuite
    public void AfterSuite() {
      // Code for after suite
    }

    @BeforeSpec
    public void BeforeSpec() {
      // Code for before spec
    }

    @AfterSpec
    public void AfterSpec() {
      // Code for after spec
    }

    @BeforeScenario
    public void BeforeScenario() {
      // Code for before scenario
    }

    @AfterScenario
    public void AfterScenario() {
      // Code for after scenario
    }

    @BeforeStep
    public void BeforeStep() {
      // Code for before step
    }

    @AfterStep
    public void AfterStep() {
      // Code for after step
    }
}
{%- language name="C#", type="csharp" -%}
public class ExecutionHooks {
    [BeforeSuite]
    public void BeforeSuite() {
      // Code for before suite
    }

    [AfterSuite]
    public void AfterSuite() {
      // Code for after suite
    }

    [BeforeSpec]
    public void BeforeSpec() {
      // Code for before spec
    }

    [AfterSpec]
    public void AfterSpec() {
      // Code for after spec
    }

    [BeforeScenario]
    public void BeforeScenario() {
      // Code for before scenario
    }

    [AfterScenario]
    public void AfterScenario() {
      // Code for after scenario
    }

    [BeforeStep]
    public void BeforeStep() {
      // Code for before step
    }

    [AfterStep]
    public void AfterStep() {
      // Code for after step
    }
}
{%- language name="Ruby", type="ruby" -%}
before_spec do
    // Code for before spec
end

after_spec do
    // Code for after step
end
{%- endcodetabs %}

> By default, Gauge clears the state after each scenario so that new objects are created for next scenario execution.
You can [configure](../advanced_readings/configuration/README.md) to change the level at which Gauge clears cache.
