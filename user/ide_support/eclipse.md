# Eclipse

There is an eclipse plugin that helps in authoring Java implementation of Gauge specs.

This plugin only supports [Java](../test_code/java.md).

## Versions supported

The plugin has been tested on  below versions of Eclipse
- 4.4 (Luna)
- 4.3.2 (Kepler SR2)

## Installation

This plugin depends on XText 2.6 or newer. This plugin can be installed via an update-site.


Steps to install:

1) Click on Help -> Install New Software

2) Ensure that the below update sites are listed in the "Available Software Sites" list:

    1) XText : http://download.eclipse.org/modeling/tmf/xtext/updates/composite/releases/
    2) Gauge : http://getgauge.io/eclipse

3) In the Dialog that opens up, choose the Gauge update site (entered above) in the "Work with" textbox, press enter.

3) Eclipse should fetch and list the below

Category: `Gauge`
Feature: `Gauge Eclipse`

4) Select the `Gauge Eclipse` feature and Click "Next"

5) Confirm the installation of "Gauge Eclipse" by clicking "Next"

6) Accept the License terms, and Click "Finish".

Eclipse should then download and setup the plugin. This will require a restart of Eclipse.


