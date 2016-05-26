# Current Execution Context in the Hook

* To get additional information about the **current specification, scenario and step** executing, an additional **ExecutionContext** parameter can be added to the [hooks](language_features/execution_hooks.md) method.

{% codetabs name="Java", type="java" -%}
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
{%- language name="C#", type="csharp" -%}
This feature is not yet supported in Gauge-CSharp.
Please refer to https://github.com/getgauge/gauge-csharp/issues/53 for updates.
{%- language name="Ruby", type="ruby" -%}
before_spec do |execution_info|
    puts execution_info.inspect
end

after_spec do |execution_info|
    puts execution_info.inspect
end
{%- endcodetabs %}
