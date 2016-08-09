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

##Execute selected data table rows
By default, scenarios in a spec are run against all the data table rows. It can be run against selected data table rows with flag `--table-rows` and specifying the row number against which the scenarios should be executed.

Example:
````
$ gauge --table-rows "1" specs/hello.spec
````

Range of table rows can also be specified, against which the scenarios are run.

Example:
````
$ gauge --table-rows "1-3" specs/hello.spec
````

This executes the scenarios against table rows 1, 2, 3.
