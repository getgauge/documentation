#### Datastore

Datastore is a way for the tests to exchange data. Datastores are basically a key-value store that have lifecycles attached to it. It is quite similar to `Session` in web application world.

##### Types

- Suite Datastore : Lives throughout the lifecycle of the entire test suite.
- Spec Datastore : Lives throughout the lifecycle of a spec, it is destroyed after the spec execution completes, and a new one is created
- Scenario Datastore : Lives throughout the lifecycle of a scenario, it is destroyed after the scenario execution completes, and a new one is created
- Step Datastore : Lives throughout the lifecycle of a step, it is destroyed after the step execution completes, and a new one is created

> Datastores are In-Memory at the moment, so they are expected to be shared only within the same machine.
