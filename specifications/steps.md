# Steps

Steps are the executable components of your specification. They  are written as markdown unordered list items (bulleted points).

Every step has an underlying code implementation for the programming language used. This is executed when the steps inside a spec are executed.

See [test code](../test_code/README.md) to understand how to write step implementations for different languages.

###Example steps
````
* Login into my app
* Search for "gauge"
* Search for "gauge-java"
````

The values written in __quotes__ are parameters which are passed into the underlying step implementation as a language specific structure.

###Learn more
* [Contexts](contexts.md)
* [Parameters](parameters.md)


