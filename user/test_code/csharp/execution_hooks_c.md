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

## Filtering Hooks execution based on tags
* You can specify tags for which the execution hooks can run. This will ensure that the hook runs only on scenarios and specifications that have the required tags.

````csharp
    // A before scenario hook that runs when tag1 and tag2 is present in the current scenario and spec.
    [BeforeScenario("tag1", "tag2")]
    public void LoginUser() {
        // Code for before scenario
    }

    // A after scenario hook runs when tag1 or tag2 is present in the current scenario and spec.
    // Default TagAggregationBehaviour value is TagAggregation.AND.
    [AfterScenario("tag1", "tag2")]
    [TagAggregationBehaviour(TagAggregation.Or)]
    public void PerformAfterScenario() {
        // Code for after scenario
    }

````

