# Execution hooks Java

The [execution hooks](../../execution/execution_hooks.md) can be defined at all levels in java by annotating methods with the following annotations:

Before Suite hooks :

| Execution Hook | annotation  |
|----------------| ------------|
| Before Suite   | @BeforeSuite|
| After Suite    | @AfterSuite|
| Before Specification   | @BeforeSpec|
| After Specification   | @AfterSpec|
| Before Scenario | @BeforeScenario|
| After Scenario   | @AfterScenario|
| Before Step | @BeforeStep |
|After Step| @AfterStep|

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



