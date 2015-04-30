# Gauge Architecture - Responsibilities of components

## Gauge Core

- install plugins

- check version compatibility

- init project
 - Copy common files
 - Send Runner an `init` request

- Loads environment

- validate specs
 - request the runner for list of all steps defined
 - scan spec files for step invocations, detect unimplemented step

- orchestrate test run
 - determine scenarios to be executed
 - orchestrate the test run
  - hold the spec->[context step]?->scenario->step structure.
 - aggregate test run results
 - broadcast event hooks
 - parallelize test execution
 - filter by tags
 - Request runner to Initialize datastore (at various event)

- refactor
 - send refactor request to language runner

- daemonized run

- expose API with helpers for plugins
 - all steps in project
 - all concepts in the project
 - parse step/spec to parameterized value

- Console report

## Plugins
### Language Runners
 - Execute hooks (when requested to)
 - Execute method corresponding to a step
  - send execution response back to Gauge
 - Refactor method signatures with changes defined by Gauge
 - Initialize a project with template
 - Read/write to datastore (currently, datastore is intra process, hence parallel processes cannot use a shared datastore)
 - Holds reference to Step text and Method / execution block.


### Others

#### Reporters
   - HTML Report (via plugin)

#### IDE (Intellij, Visual Studio, Eclipse)
   - Syntax highlighting
   - Autocomplete
   - Goto definition
   - Refactoring
   - Test runner
