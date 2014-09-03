#Specifications


**Specifications** are business layer test cases which can also act as your feature documentation. They are written in the Business language. Typically a spec or specification describe a particular feature of the application under test.

* They are written in a **.spec** file. Gauge also support **.md** file format.

* The Markup for a Specification file is based on [markdown](https://en.wikipedia.org/wiki/Markdown) syntax.


#### A simple Spec:

![Spec](images/spec.png "Specification")


### Specification Heading

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




