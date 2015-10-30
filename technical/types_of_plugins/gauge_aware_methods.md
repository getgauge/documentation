#### Gauge aware methods

These are methods/blocks that Gauge would understand and execute at various points.

##### Step
These are implementation of one or more business steps defined in the `spec` files. These methods are either annotated with `@Step`(Java), `[Step]`(C#) or used with`step <block>`(Ruby).

##### Hook
These are methods that get invoked at various points in the execution life cycle. The hooks available are:
- `before_suite`
	- `before_spec`
		- `before_scenario`
			- `before_step`
			- `after_step`
		- `after_scenario`
	- `after_spec`
- `after_suite`

#### Datastore

Datastore is a way for the tests to exchange data. Datastores are basically a key-value store that have lifecycles attached to it. It is quite similar to `Session` in web application world.

##### Types

- Suite Datastore : Lives throughout the lifecycle of the entire test suite.
- Spec Datastore : Lives throughout the lifecycle of a spec, it is destroyed after the spec execution completes, and a new one is created
- Scenario Datastore : Lives throughout the lifecycle of a scenario, it is destroyed after the scenario execution completes, and a new one is created

> Datastores are In-Memory at the moment, so they are expected to be shared only within the same machine.
