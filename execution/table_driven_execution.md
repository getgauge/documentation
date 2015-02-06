# Table driven execution
When a scenario or multiple scenarios in a specification are to be executed for multiple sets of data then Data table execution can be used.

* A **data table** is defined in markdown table format in the beginning of the spec before any steps.
* The data table should have a header row and one or more data rows
* The header names from the table can be used in the steps within angular brackets `< >` to refer a particular column from the data table as a parameter.
* On execution each scenario will be executed for every data row from the table.

**Example:**

```
Table driven execution
======================

     |id| name    |
     |--|---------|
     |1 |vishnu   |
     |2 |prateek  |
     |3 |navaneeth|

Scenario
--------
* Say "hello" to <name>

Second Scenario
---------------
* Say "namaste" to <name>
```

In the above example the step uses the `name` column from the data table as a dynamic parameter.

Both `Scenario` and `Second Scenario` are executed first for the first row values `1, vishnu` and then consecutively for the second and third row values from the table.


### Data Table values in inline tables
Dynamic values from the data table can also be referred in table parameters passed into steps

Example
````
Create projects
===============

|id| name |
|--|------|
|1 | john |
|2 | mike |

first scenario
--------------

* Create the following projects
     |project name| username |
     |------------|----------|
     | Gauge java | <name>   |
     | Gauge ruby | <name>   |

````

In the above example the table parameter uses a dynamic value from the data table.