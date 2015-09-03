# Custom messages in reports

Custom messsages/data can be added to execution reports using the below API from the step implementations or hooks.

````csharp
GaugeMessages.WriteMessage("Custom message for report");

GaugeMessages.WriteMessage("Welcome to {0}", "gauge");
````
These messages will appear under steps in the execution reports.
