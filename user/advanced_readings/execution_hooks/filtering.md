# Filtering Hooks execution based on tags

* You can specify tags for which the execution [hooks](../../../language_features/execution_hooks.md) can run. This will ensure that the hook runs only on scenarios and specifications that have the required tags.

{% codetabs name="Java", type="java" -%}
// A before spec hook that runs when tag1 and tag2
// is present in the current scenario and spec.
@BeforeSpec(tags = {"tag1, tag2"})
public void loginUser() {
    // Code for before scenario
}


// A after step hook runs when tag1 or tag2
// is present in the current scenario and spec.
// Default tagAggregation value is Operator.AND.
@AfterStep(tags = {"tag1", "tag2"}, tagAggregation = Operator.OR)
public void performAfterStep() {
    // Code for after step
}
{%- language name="C#", type="cs" -%}
// A before spec hook that runs when tag1 and tag2
// is present in the current scenario and spec.
[BeforeSpec("tag1, tag2")]
public void LoginUser()
{
    // Code for before scenario
}


// A after step hook runs when tag1 or tag2
// is present in the current scenario and spec.
// Default tagAggregation value is Operator.AND.
[AfterStep("tag1", "tag2")]
[TagAggregationBehaviour(TagAggregation.Or)]
public void PerformAfterStep()
{
    // Code for after step
}
{%- language name="Ruby", type="ruby" -%}
# A before spec hook that runs when tag1 and tag2 is present in the current scenario and spec.
before_spec({tags: ['tag2', 'tag1']}) do
    // Code for before scenario
end

# A after step hook runs when tag1 or tag2 is present in the current scenario and spec.
# Default tagAggregation value is Operator.AND.
after_spec({tags: ['tag2', 'tag1'], operator: 'OR'}) do
    // Code for after step
end
{%- endcodetabs %}

> Note: Tags cannot be specified on @BeforeSuite and @AfterSuite hooks
