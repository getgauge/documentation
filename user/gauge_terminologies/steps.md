# Steps
------
Steps are the executable components of your specification. They  are written as markdown unordered list items (bulleted points).

They are written inside a specification as
- [context steps](contexts.md)
- [tear down steps](tear_down_steps.md)
- steps inside a scenario or concepts

Every step has an underlying code implementation for the programming language used. This is executed when the steps inside a spec are executed.

See how to write [step implementations](language_features/step_implementations.md) for different languages.

## Example

```
* Login into my app
* Search for "gauge"
* Search for "gauge-java"
```

The values written in __quotes__ are parameters which are passed into the underlying step implementation as a language specific structure.

> Note: The following characters are reserved for parameters, these cannot be used in step text.
- "
- <
- >

## Further reading
* [context steps](contexts.md)
* [tear down steps](tear_down_steps.md)
* [Scenarios](scenarios.md)
* [Concepts](concepts.md)
* [Parameters](parameters/README.md)
