# XML Report
XML Report plugin creates JUnit XML test result document that can be read by tools such as Go, Jenkins.
When the specs are executed, the xml report is generated in reports directory in the project. The format of XML report is based on [JUnit XML Schema](https://windyroad.com.au/dl/Open%20Source/JUnit.xsd)

**Sample XML Report Document** : 
```
<testsuites>
	<testsuite id="1" tests="1" failures="0" package="specs/hello_world.spec" time="0.002" timestamp="2015-09-09T13:52:00" name="Specification Heading" errors="0" hostname="INcomputer.local">
		<properties></properties>
		<testcase classname="Specification Heading" name="First scenario" time="0.001"></testcase>
		<system-out></system-out>
		<system-err></system-err>
	</testsuite>
</testsuites>
```
## Installation
To install XML Report plugin :
```
gauge --install xml-report
```

To install a specific version of XML report plugin use the `--plugin-version` flag.
````
gauge --install xml-report --plugin-version 0.0.2
````
__Offline Installation__ :

If plugin should be installed from a zipfile instead of downloading from plugin repository, use the `--file` or `-f` flag.
````
gauge --install xml-report --file ZIP_FILE_PATH
````
Download the plugin zip from the [Github Releases](https://github.com/getgauge/xml-report/releases)

## Configuration

To add XML report plugin to your project, run the following command :
```
gauge --add-plugin xml-report
```
The XML report plugin can be configured by the properties set in the `env/default.properties` file in the project.

The configurable properties are:

__gauge_reports_dir__
* Specifies the path to the directory where the execution reports will be generated.

* Should be either relative to the project directory or an absolute path.
By default it is set to `reports` directory in the project

__overwrite_reports__
* Set to `true` if the reports **should be overwritten** on each execution hence maintaining only the latest execution report.

* If set to `false` then a **new report** will be generated on each execution in the reports directory in a nested time-stamped directory.

Default value is `true`
