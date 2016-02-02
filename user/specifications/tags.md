# Tags

Tags are used to associate labels with specifications or scenarios. Tags are written as comma separated values in the spec with a prefix **'Tags:'**' .

* Both scenarios and specifications can be separately tagged
* Only **one** set of tags can be added to a single specification or scenario.

They help in filtering specs or scenarios based on tags used.

## Example

Both the **Login specification** and the scenario **Successful login scenario** have tags in the below example.

```
Login specification
===================
 Tags: login, admin, user-abc


Successful login scenario
-------------------------
 Tags: login-success, admin
```

A tag applied to a spec automatically applies to a scenario.
