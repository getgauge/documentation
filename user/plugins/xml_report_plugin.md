# XML Report

[Gauge](https://github.com/getgauge/gauge) plugin to generate JUnit style XML reports in the format described [here](https://windyroad.com.au/dl/Open%20Source/JUnit.xsd)
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
