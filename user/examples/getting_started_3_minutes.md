# Getting started in 3 minutes

## Prerequisites

* JDK 6+ / .NET framework / Ruby 2.0+ is installed
* [Gauge](../../installations/operating_system) is installed.
* Install the language runner
{% codetabs name="Java", type="java" -%}
$ gauge --install java

{%- language name="C#", type="csharp" -%}
$ gauge --install csharp

{%- language name="Ruby", type="ruby" -%}
$ gauge --install ruby
{%- endcodetabs %}

One can use IDEs to create the projects and run specifications, for this example, we are using the command line options.

## Creating a Project

{% codetabs name="Java", type="java" -%}
$ gauge --init java

{%- language name="C#", type="csharp" -%}
$ gauge --init csharp

{%- language name="Ruby", type="ruby" -%}
$ gauge --init ruby
{%- endcodetabs %}

On getting the message of `Successfully initialized the project`, one should be able to run the specifications.

## Running the specs

```
gauge specs
```
The details of the run are displayed on the command line followed by a statistics summary.

## Interpret results

The console report would give you the details of the run
```
Specifications:	(w) executed	(x) passed	(y) failed	(z) skipped
Scenarios:	(a) executed	(b) passed	(c) failed	(d) skipped
```
The statistics of the scenarios indicate
* (a) is the total number of scenarios executed.
* (b) is the total number of scenarios passed.
* (c) is the total number of scenarios failed.
* (d) is the total number of scenarios skipped.

The statistics of the specifications indicate
* (w) is the total number specifications executed.
* (x) is the total number specifications with all scenarios passed.
* (y) is the total number specifications with atleast one scenario failed.
* (z) is the total number specifications with all scenarios skipped.

> Other [formats of reporting](../reporting_features/README.md) are available.
