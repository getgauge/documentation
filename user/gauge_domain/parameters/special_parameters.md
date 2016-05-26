# Special Parameters
-----
Special parameters provide the ability to pass larger and richer data into the steps as parameters.
* They are entered in angular brackets - `<>` in the step. There are 2 types of special parameters available in Gauge
* They contain 2 parts separated by a colon `:`

**Prefix** : This defines the type of special parameter. e.g. file, table.

**Value** : This defines the value for the type of special parameter.

```
<prefix:value>

```

The different special parameter types are:

## File

These are used to read files and pass the file content as a string parameter to the underlying steps.

The prefix and value are below:

**Prefix** : The prefix is ***file***

**Value**  : The value is the path to the file.

```
* Verify email text is <file:email.txt>
* Check if <file:/work/content.txt> is visible
```
The path to the file can be the relative path from the Gauge project or an absolute path to the file.

## CSV
Tables are used to pass table value into steps read from an external CSV file. The parameter text in the step contains a  prefix table and the path to the csv file.

**Prefix** : The prefix is ***table***

**Value**  : The value is the path to the csv file.

```
* Step that takes a table <table:data.csv>
* Check if the following users exist <table : /Users/john/work/users.csv>
```

**Sample csv file**:

```
Id,Name
1,The Way to Go On
2,Ivo Jay Balbaert
```

The first row is considered as table header. Following rows are considered as the row values.
