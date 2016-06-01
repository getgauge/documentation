# Command Line Interface
------
Gauge has **first-class command line support**. If you have gauge [installed](../installations/README.md), running `gauge` in terminal gives you command usage and all the flags it supports.

The command-line interface works across platforms. On GNU/Linux and OSX, you can use any terminal. On Windows, you can use `cmd` or Powershell.

## Creating a project

To create or initialize a Gauge project use the `gauge --init` command. For details, see how to [create a Gauge project](../getting_started/creating_a_gauge_project.md).

**Example: Create a project**

{% codetabs name="Java", type="java" -%}
  $ gauge --init java
{%- language name="C#", type="cs" -%}
  $ gauge --init csharp
{%- language name="Ruby", type="ruby" -%}
  $ gauge --init ruby
{%- endcodetabs %}


## Executing tests

Inside a Gauge project, you can execute your tests by invoking `gauge` with path to [specifications](../gauge_terminologies/specifications.md). By convention, specifications are stored in the the `./specs/` sub-directory in the project root. The syntax is:

```sh
$ gauge [options] <path-to-specs>
```

The `gauge` command-line utility allows multiple ways to specify the specifications to be executed. A valid path for executing tests can be path to directories that contain specifications or path to specification files or path to scenarios or a mix of any of these three methods.

### Specify directories

You can specify a single directory in which specifications are stored. Gauge scans the directory and picks up valid specification files.

For example:

```sh
$ gauge specs/
```

You can also specify multiple directories in which specifications are stored. Gauge scans all the directories for valid specification files and executes them in one run.

For example:

```sh
$ gauge specs-dir1/ specs-dir2/ specs-dir3/
```

### Specify specification files

You can specify path to a specification files. In that case, Gauge executes only the specification files provided.

For example, to execute a single specification file:

```sh
$ gauge specs/spec1.spec
```

Or, to execute multiple specification files:

```sh
$ gauge specs/spec1.spec specs/spec2.spec specs/spec3.spec
```

### Specify scenarios

You can also specify a specific [scenario](../gauge_terminologies/scenarios.md) or a list of scenarios to execute. To execute scenarios, `gauge` takes path to a specification file, followed by a colon and a zero-indexed number of scenarios.

For example, to execute the second scenario of a specification file named `spec1.spec`, you would do:

```sh
$ gauge specs/spec1.spec:1
```

To specify multiple scenarios, add multiple such arguments. For example, to execute the first and third scenarios of a specification file named `spec1.spec`, you would do:

```sh
$ gauge specs/spec1.spec:0 specs/spec1.spec:2
```

## Verbose reporting

By default, `gauge` reports at the specification level when executing tests. You can enable verbose, step-level reporting by using the `--verbose` flag. For example:

```sh
$ gauge --verbose specs/
```

---

## Suggested reading

* [Command line flags](flags.md)
* [Specifications](../gauge_terminologies/specifications.md)
* [Scenarios](../gauge_terminologies/scenarios.md)
* [Tagged execution](../advanced_readings/execution/tagged_execution.md)
* [Parallel execution](../advanced_readings/execution/parallel_execution.md)
