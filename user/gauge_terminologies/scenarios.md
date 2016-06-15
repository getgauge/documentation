# Scenarios

Each scenario represents a single flow in a particular specification. A specification must contain at least one scenario.

A scenario starts after a scenario heading or a scenario name. The scenario heading is written in markdown **`<H2>`** syntax. This can be written in 2 ways:

```
Scenario heading
----------------
```

or


```
## Scenario heading
```

* A scenario contains one or more [steps](steps.md) under it.
* A scenario can be tagged using [tags](tags.md).

## Example spec with a scenario:

```
Configuration
=============

The Admin user should be able to switch permissions for other users.

Admin Login
-----------
* User must login as "admin"
* Navigate to the configuration page
* Change permissions for user "john" to "admin"
* User "john" should have admin permissions
```

## Further reading
* [Steps](steps.md)
* [Tags](tags.md)
