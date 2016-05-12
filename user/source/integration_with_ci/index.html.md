---
title: Integration with CI
---

# Integration with CI

Gauge can be easily integrated with any [Continuous Integration](http://martinfowler.com/articles/continuousIntegration.html) environment.

Since Gauge supports first class command line, invoking it from any CI/CD tool is very straightforward.

Steps to Integrate Gauge with CI tool:

* Install the Gauge and language plugin on CI machine
* Add gauge commands as tasks in CI to run tests.

    For example, to run the specs use `gauge specs`
* If you want to run specific instance of gauge on CI, set `GAUGE_ROOT` as environment variable to the path of specific instance.
* Gauge returns html-reports, console output as result of execution which can be configured to view on CI.
