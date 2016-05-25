# Step implementations

Every [step](../gauge_domain/steps.md) needs to have a language specific implementation that gets executed on the spec execution.

###Examples

#### Simple step

**Step name**
```
* Say "hello" to "gauge"
```

**Implementation**

{% codetabs name="Java", type="java" -%}
// This Method can be written in any java class as long as it is in classpath.

public class StepImplementation {

    @Step("Say <greeting> to <product name>")
    public void helloWorld(String greeting, String name) {
        // Step implementation
    }
}
{%- language name="C#", type="csharp" -%}
// The Method can be written in **any C# class** as long as it is part of the project.
public class StepImplementation {

    [Step("Say <greeting> to <product name>")]
    public void HelloWorld(string greeting, string name) {
        // Step implementation
    }
}
{%- language name="Ruby", type="ruby" -%}
step 'Say <greeting> to <product name>' do |greeting, name|
 # Code for the step
end
{%- endcodetabs %}

> The text passed to the annotation is the [step template](step_name_template.md), which is a simple representation of the step.

#### Step with table
**Step:**

````
* Create following "hobbit" characters
|id |   name  |
|---|---------|
|123| frodo |
|456| bilbo  |
|789|samwise|
````

**Implementation:**

{% codetabs name="Java", type="java" -%}
// Table is a custom data structure defined by gauge.
public class Users {

    @Step("Create following <race> characters <table>")
    public void createCharacters(String type, Table table) {
        // Step implementation
    }
}
{%- language name="C#", type="csharp" -%}
// Here Table is a custom data structure defined by gauge.
// This is available by adding a reference to the Gauge.CSharp.Lib.
// Refer : http://nuget.org/packages/Gauge.CSharp.Lib/
public class Users {

    [Step("Create following <role> users <table>")]
    public void HelloWorld(string role, Table table) {
        // Step implementation
    }
}
{%- language name="Ruby", type="ruby" -%}
# Here table is a custom data structure defined by gauge-ruby.
step 'Create following <race> characters <table>' do |role, table|
  puts table.rows
  puts table.columns
end
{%- endcodetabs %}

## Multiple Step names (Alias for Step names)
The same implementation can point to multiple step names. To do this add all the [step name template](step_name_template.md) to the `[Step]` attribute.

The number and type of parameters for all the steps names must match the number of parameters on the C# implementation.

####Example

**Step Names**
````
* Create a wizard "Gandalf"
* Create another wizard "Saruman"
````
**Implementation**


{% codetabs name="Java", type="java" -%}
public class Users {

    @Step({"Create a wizard <wizard>", "Create another wizard <wizard>"})
    public void helloWorld(String role) {
        // Step implementation
    }
}
{%- language name="C#", type="csharp" -%}
public class Users {

    [Step({"Create a wizard <wizard>", "Create another wizard <wizard>"})]
    public void HelloWorld(string role) {
        // Step implementation
    }
}
{%- language name="Ruby", type="ruby" -%}
step "Create a wizard <wizard>", "Create another wizard <wizard>" do |username|
 #step code
end
{%- endcodetabs %}

---

Learn More
* [Configuration](configuration.md)
* [Execution Hooks](execution_hooks.md)
* [Data Store](data_store_to_share_data.md)
