# Table Parameter

* Table Parameters can be used in two ways
- When a scenario or multiple scenarios in a specification are to be executed for multiple sets of data then Data table execution can be used.
- Tables or inline tables can be passed to steps as parameters.

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

First scenario
--------------

* Create the following projects
     |project name| username |
     |------------|----------|
     | Gauge java | <name>   |
     | Gauge ruby | <name>   |
```

In the above example the table parameter uses a dynamic value from the data table.

## Further reading
* [External data source](special_parameters.md) as parameters
