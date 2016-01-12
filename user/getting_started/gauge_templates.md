# Gauge Project Templates

The Gauge Templates can be used to bootstrap the process of initializing a Gauge project along with a suitable build dependency tool, webdriver etc.

To list all the Gauge project templates available, run the following command:

```
gauge --list-templates
```

These templates can also be found in [Bintray Gauge Templates](https://bintray.com/gauge/Templates/gauge-templates/view#files).

## Initialize a Gauge project with Template

Say you want to initialize a Gauge project with Maven as build tool and selenium Webdriver as a driver of choice. You can quickly setup such project which is ready to start writing tests with selenium by using `java_maven_selenium` Gauge template.

To do this, run the following command:
```
gauge --init java_maven_selenium
```
This will download the Gauge template `java_maven_selenium` and setup your project ready to use.

Now, start writing your [Specifications](../specifications/README.md) and execute them.
