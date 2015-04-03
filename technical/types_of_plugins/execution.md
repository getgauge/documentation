#### Execution

Gauge execution can be one of the three types as mentioned below.

##### Setup
This phase of execution initializes a new Gauge project. The setup consists of below sequence of steps:
- Gauge checks if the plugin for the specified language is installed.
- If plugin isn't installed, it queries the Gauge-Repository to see if there is a plugin available
- Throws out an error if an invalid language is specified
- Gauge installs the language runner plugin if it isn't installed already
- Copies Gauge specific files, `manifest.json`, `hello_world.spec` etc.
- Invokes the runner with a `--setup` flag.
- Runner responds to command line argument
- Runner the copies the language specific files.

> Currently there is only one template that is copied over, but there are [plans to add more templates](https://github.com/getgauge/gauge/issues/89).

##### Scan
This phase is implicit. This happens as a pre-requisite to running specifications. The Runner is invoked with a `--start` flag, which the Runner responds to by:
- Building the target project (when required)
- Scan and build a cache of Types, Step Implementations and Hook implementations.
- Wait for execution requests.

##### Run Specs
As the name suggests, this phase would run the implementations. However there are some pre and post run tasks that also get executed:
- Datastore initialization/Clear: Initialize a datastore at the right level. Clear datastore at appropriate points of the execution lifecycle.
- Hooks execution: Execute hooks at appropriate points of the execution lifecycle.
