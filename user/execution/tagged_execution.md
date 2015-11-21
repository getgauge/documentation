# Tagged Execution

Tags allow you to filter the specs and scenarios quickly for execution. To execute all the specs and scenarios which are labelled with certain tags, use the following command.

````
gauge --tags tag1,tag2 specs
````
or
````
gauge --tags "tag1, tag2" specs
````
This executes only the scenarios and specifications which are tagged with tag1 and tag2.

Example:

![Spec](../specifications/images/spec.png "Specification")

In the above spec, if all the scenarios tagged with "search" and "admin" should be executed, then use the following command.

````
gauge --tags "search & admin" SPEC_FILE_NAME
````


### Tags can be selected using expressions
__Examples__
* __!TagA:__  Selects specs/scenarios that do not have TagA.
* __TagA & TagB:__  Selects specs/scenarios that have both TagA and TagB.
* __TagA & !TagB:__  Selects specs/scenarios that have TagA and not TagB.
* __TagA | TagB:__  Selects specs/scenarios that have either TagA or TagB.
* __(TagA & TagB) | TagC:__  Selects specs/scenarios that have either TagC or both the tags TagA and TagB
* __!(TagA & TagB) | TagC:__  Selects specs/scenarios that have either TagC or do not have both the tags TagA and TagB
* __(TagA | TagB) & TagC:__  Selects specs/scenarios that either TagA and TagC or TagB and TagC

