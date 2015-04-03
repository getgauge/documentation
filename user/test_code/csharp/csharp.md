C#
==

## Step implementations

Step implementations for C# are methods in C# classes.

The method implementation for the particular step has to be decorated with`[Step]` attribute which takes the [step name template](../step_name_template.md) as value.

###Examples


####1. Simple step

**Step name**
```
* Say "hello" to "gauge"
```

**C# Implementation:**
````csharp

public class StepImplementation {

    [Step("Say <greeting> to <product name>")]
    public void HelloWorld(string greeting, string name) {
        // Step implementation
    }
}
````

* The text passed to the annotation is the [step template](../step_name_template.mdd), which is a simple representation of the step.
* The Method can be written in **any C# class** as long as it is part of the project.


####2. Step with table
**Step:**

````
Create following "admin" users
|id |   name  |
|---|---------|
|123| prateek |
|456| vishnu  |
|789|navaneeth|
````

**C# Implementation:**
````csharp
public class Users {

    [Step("Create following <role> users <table>")]
    public void HelloWorld(string role, Table table) {
        // Step implementation
    }
}

````
* Here **Table** is a custom data structure defined by gauge. This is available by adding a reference to the [Gauge.CSharp.Lib](http://nuget.org/packages/Gauge.CSharp.Lib/) nuget package.

## Multiple Step names
The same implementation can point to multiple step names. To do this add all the [step name template](../step_name_template.md) to the `[Step]` attribute.

The number and type of parameters for all the steps names must match the number of parameters on the C# implementation.

####Example

***Step Names :***
````
* Create a user "srikanth"
* Create another user "vishnu"
````
***Implementation :***

````csharp
public class Users {

    [Step({"Create a user <username>", "Create another user <username>"})]
    public void HelloWorld(string role, Table table) {
        // Step implementation
    }
}

````

Learn More
* [Execution Hooks](execution_hooks_c.md)

