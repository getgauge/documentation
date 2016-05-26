# Table Parameter
----
Tables or inline tables can be passed to steps as parameters. They will be available as a language specific table data structure in the underlying implementation.

* Table parameters are written in Multi-markdown table formats.
* The first row contain the table headers. Following rows contains the row values.
* A separator between the header and the other rows is optional

```
* Step that takes "this" and a table
 | id   |  name   |
 | 123  |  John   |
 | 456  | Mcclain |

* Another step with a table parameter
 | id   |  name    |
 |----- | ---------|
 |123   |   John   |
 |456   |  Mcclain |
```
