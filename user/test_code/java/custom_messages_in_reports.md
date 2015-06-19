# Custom messages in reports

Cutom messsages/data can be added to execution reports using the below API from the step implementations or hooks.

````java
Gauge.writeMessage("Custom message for report");

String id = "4567";
Gauge.writeMessage("User id is %s", id);
````
These messages will be appear under steps in the execution reports.
