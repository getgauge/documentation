# Execution Hooks Ruby

The [execution hooks](../../execution/execution_hooks.md) can be defined at all levels in ruby by calling the necessary execution hook method and passing a code block with thecode to run at the level.

| Execution Hook | Method  |
|----------------| ------------|
| Before Suite   | `before_suite`|
| After Suite    | `after_suite`|
| Before Specification   | `before_spec`|
| After Specification   | `after_spec`|
| Before Scenario | `before_scenario`|
| After Scenario   | `after_scenario`|
| Before Step | `before_step` |
|After Step| `after_step`|


###Examples
````ruby
 before_suite do
   # Code for the hook
 end

 after_step do
   put "Inside after step hook"
 end

````



