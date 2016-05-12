---
title: Gauge Project Templates
---

# Gauge Project Templates

Gauge provides templates that can be used to bootstrap the process of initializing a Gauge project along with a suitable build dependency tool, webdriver etc.

To list all the Gauge project templates available, run the following command:

```
$ gauge --list-templates
```

These templates can also be found in [Bintray Gauge Templates](https://bintray.com/gauge/Templates/gauge-templates/view#files).

## Initialize a Gauge project with Template

Say you want to initialize  as a driver of choice. You can quickly setup such project which is ready to start writing tests with selenium by using `java_maven_selenium` Gauge template.

To initialize a Gauge project with a template, choose a name from the list shown on running `gauge --list-templates` and pass that name as an argument when initializing the Gauge project.

For example, to create a Gauge project with the `java_maven_selenium` template, you need to run this command:

```
$ gauge --init java_maven_selenium
```

This template creates a Gauge project with Maven as build tool and the selenium Webdriver. This will download the Gauge template `java_maven_selenium` and setup your project with useful sample code.

Now, you can start writing [Specifications](../specifications/README.md) and execute them.
