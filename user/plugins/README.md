# Plugins

Plugins are an easy way to extend the features of gauge. There are various types of plugins that gauge currently supports.

## Language Plugins

Also known as **language runners**, these plugins enable the users to write the implemantation of specs in a language of their choice. For a list of available language runners, see how to [install language runners](../installations/install_language_runners.md). Allows you to write your tests in

* [Java](http://github.com/getgauge/gauge-java)
* [CSharp](http://github.com/getgauge/gauge-csharp)
* [Ruby](http://github.com/getgauge/gauge-ruby)

## Reporting plugins

These are plugins that work during the execution of the specifications to give summary/details of the run.
Example: [HTML Report](../reporting_features/html_report_plugin.md), [XML Report](../reporting_features/xml_report_plugin.md)

## IDE plugins

Gauge has a bunch of plugins so that users can easily author specs on IDE. For more details, check [IDE Support](../ide_support/README.md).

* Integration with [IntelliJ Idea](../ide_support/intellij_idea.md)
* Integration with [Visual Studio 2013](../ide_support/visual_studio.md) (Visual Studio 2015 support is currently under testing).
* Integration with [Eclipse](../ide_support/eclipse.md)

## Reporting

 * [HTML Report](http://github.com/getgauge/html-report)
 * [XML Report](http://github.com/getgauge/xml-report) to help integrated with CI/CD systems.

>Some plugins are added to every new project created in gauge.
 - [HTML Report Plugin](../reporting_features/html_report_plugin.md)

## Other

* [Maven](https://github.com/getgauge/gauge-maven-plugin)
* [Gradle](https://github.com/manupsunny/gauge-gradle-plugin)
