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
