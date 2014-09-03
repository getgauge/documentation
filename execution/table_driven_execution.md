# Table driven execution

Using table driven execution, single spec can be run multiple times with different parameters

example spec:

```
Table driven execution
======================

     |id|name     |
     |--|---------|
     |1 |vishnu   |
     |2 |prateek  |
     |3 |navaneeth|

Scenario
--------

* Say "hello" to <name>
```

Each datatable row is executed and the dynamic parameter `name` in the step `say "hello" to <name>` points to that particular row's value.
