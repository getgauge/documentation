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

In the above spec, if all the scenarios tagged with "successful" should be executed, then use the following command.

````
gauge --tags "successful" specs
````
Here, last parameter "specs" specifies the specifications directory.

