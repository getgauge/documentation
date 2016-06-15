# Table Parameter

* Table Parameters can be used in two ways
- When a scenario or multiple scenarios in a specification are to be executed for multiple sets of data then Data table execution can be used.
- Tables or inline tables can be passed to steps as parameters.

## Data driven execution
* A *data table* is defined in markdown table format in the beginning of the spec before any steps.
* The data table should have a header row and one or more data rows
* The header names from the table can be used in the steps within angular brackets `< >` to refer a particular column from the data table as a parameter.
* On execution each scenario will be executed for every data row from the table.
* Table can be easily created in IDE using template `table:<no of columns>`, and hit `Tab`.
* Table parameters are written in Multi-markdown table formats.

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

## External CSV for data table

Data Tables for a specification can also be passed from an external [CSV file](http://en.wikipedia.org/wiki/Comma-separated_values). The parameter contains a prefix table and the path to the csv file.

**Prefix** : The prefix is table

**Value** : The value is the path to the csv file. This can be absolute file path or relative to project.

Example:

```
Table Driven execution
=====================

table: /system/users.csv

Scenario
--------
* Say "hello" to <name>

Second Scenario
---------------
* Say "namaste" to <name>
```

Sample csv file :

```
Id,Name
1,The Way to Go On
2,Ivo Jay Balbaert
```

**The first row is considered as table header. Following rows are considered as the row values.**

## Data Table values in inline tables
Dynamic values from the data table can also be referred in table parameters passed into steps

Example:

```
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
```

In the above example the table parameter uses a dynamic value from the data table.
