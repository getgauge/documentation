# Syntax Cheat Sheet

####Specification
```
Specification             # Specification
=============
```

####Scenario

```
Scenario                  ## Scenario
--------
```

####Tags
````
Specification Level        Scenario Level
===================        --------------
Tags: login, admin         Tags: login-success, admin
````

####Concept
```
Concept Heading             # Concept Heading
===============
```

####Steps
````
* Step Name
````

####Parameters

* `"`Static Arg`"`
````
    Check "product" exists
````

* `<`Dynamic Arg`>`
````
 Check <product> exists
````

* `|`Table Parameter`|`
````
 Step that takes a table
 | id   |  name   |
 | 123  |  John   |
 | 456  | Mcclain |
````

* `<`Special`:`Parameters`>`
````
 <prefix:value>
````

 * *File*
````
   * Check if <file:/work/content.txt> is visible
````

 * *Table*
````
   * Check if the users exist <table : /Users/john/work/users.csv>
````
