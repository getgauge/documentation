Execution hooks C\#
===================

The [execution hooks](../../execution/execution_hooks.md) can be defined at all levels in C# by decorating methods with the one (or more) of the following attributes:


| Execution Hook | Attributes  |
|----------------| ------------|
| Before Suite   | `BeforeSuite`|
| After Suite    | `AfterSuite`|
| Before Specification   | `BeforeSpec`|
| After Specification   | `AfterSpec`|
| Before Scenario | `BeforeScenario`|
| After Scenario   | `AfterScenario`|
| Before Step | `BeforeStep` |
|After Step| `AfterStep` |

###Example
````csharp
public class ExecutionHooks {

    [BeforeSpec]
    public void LoginUser() {
    // Code for before spec
    }

    [AfterStep]
    public void PerformAfterStep() {
    // Code for after step
    }
}
````



