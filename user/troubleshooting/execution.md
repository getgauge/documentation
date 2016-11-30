# Execution

It is advised to use the latest version of gauge and [gauge plugins](../plugins/README.md). See our [download page](http://getgauge.io/get-started/index.html) for links to latest installation

### Validation Errors
````
[WARN] Validation failed. The following steps have errors
...

````
These generally occur if step implementation is not found for a particular step.
* Ensure the [step implementation](../language_features/step_implementations.md) for the step has been added.
* The [step template](../language_features/step_implementations.md) marking the step in code is case sensitive and should match the step usage in the spec file.


### Compatibility errors
````
Failed to start a runner. Compatible runner version to 0.0.7 not found
````
* The language plugin installed is not compatible with the gauge version installed.
* Run ```gauge --install language_NAME ``` to install the latest compatible version. See [plugin installation](../Installations/install_language_runners.md) for more details


### Execution Errors
```
Error: too many open files
```
* This error occurs when the upper limit to open the number of files is too low. To fix the error, increase the upper limit by adding the command `ulimit -S -n 2048` to your `~/.profile` file and relogin.
