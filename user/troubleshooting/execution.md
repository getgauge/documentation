# Execution

It is advised to use the latest version of gauge and [gauge plugins](../plugins/README.md). See our [download page](http://getgauge.io/download) for links to latest installation

### Validation Errors
````
[WARN] Validation failed. The following steps have errors
...

````
These generally occur if step implementation is not found for a particular step.
* Ensure the [test_code](../test_code/step_implementations.md) for the step has been added.
* The [step_template](../test_code/step_name_template.md) marking the step in code is case sensitive and should match the step usage in the spec file.


### Compatibility errors
````
Failed to start a runner. Compatible runner version to 0.0.7 not found
````
* The langauge plugin installed is not compatible with the gauge version installed.
* Run ```gauge --install LANGAUGE_NAME ``` to install the latest compatible version. See [plugin installation](../plugins/installation.md) for more details


### Execution Errors
```
Error: too many open files
```
* This error occurs when the upper limit to open the number of files is too low. To fix the error, increase the upper limit by adding the command `ulimit -S -n 2048` to your `~/.profile` file and relogin.

