#Specifications


**Specifications** are business layer test cases which can also act as your feature documentation. They are written in the Business language. Typically a spec or specification describe a particular feature of the application under test.

* They are written in a **.spec** file. Gauge also support **.md** file format.

* The Markup for a Specification file is based on [markdown](https://en.wikipedia.org/wiki/Markdown) syntax.


#### A simple Spec:

![Spec](images/spec.png "Specification")


## Specification Heading

A Spec must begins with a spec heading and a single specification can contain only one spec heading.

It is written in H1 syntax of markdown. This can be in two forms

````
Spec Heading
============
````
 or

````
#Spec Heading
````

* Every spec must contain one or more [scenarios](scenarios.md).
* Every spec can be marked with labels using [tags](tags).


## Scenarios

Each scenario represents a single flow in a particular specification.

Consider the above example, **successful search** and **unsuccessful search** are 2 scenarios for the **search specification**. In business language terms they represent the entire search feature in the app.

Learn more about the structure of a [Scenario](scenarios.md)

## Steps

Steps are the executable part of the specification. They have underlying implementations in the programmming language for execution of the specs.

They are written inside a specification ([context steps](contexts.md)) or a scenario.

#### Example
````
* Search for product "die hard"
* User "admin" must be logged in
````

Learn more about [steps](steps.md).







