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
