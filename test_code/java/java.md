# Java

## Step implementations

Step implementations for java are methods in java classes. They can be any class as long as they are in class path.

The method implementation for the particular step has to be annotated with`@Step` annotation which takes the [step name template](../step_name_template.md) as value.

###Examples


####1. Simple step

**Step name**
```
* Say "hello" to "gauge"
```

**Java Implementation:**
````java

public class StepImplementation {

    @Step("Say <greeting> to <product name>")
    public void helloWorld(String greeting, String name) {
        // Step implementation
    }
}
````

* The text passed to the is the [step template](../step_name_template.mdd), whichis a simple representation of the step.
* The Method can be written in **any java class** as long as it is in classpath.


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

**Java Implementation:**
````java
public class Users {

    @Step("Create following <role> users <table>")
    public void helloWorld(String role, Table table) {
        // Step implementation
    }
}

````
* Here **Table** is a custom data structure defined by gauge.

## Multiple Step names
The same implementation can point to multiple step names. To do this add all the [step name template](../step_name_template.md) to the ``@Step`` annotations.

The number and type of parameters for all the steps names must match the number of parameters on the java implementation.

####Example

***Step Names :***
````
* Create a user "prateek"
* Create another user "vishnu"
````
***Implementation :***

````java
public class Users {

    @Step({"Create a user <username>", "Create another user <username>"})
    public void helloWorld(String role, Table table) {
        // Step implementation
    }
}

````

Learn More
* [Configuration](configuration.md)

