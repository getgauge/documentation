# HTML Reports

Reports are generated using `html-report` plugin. By default html-report is added to the project.

When the specs are executed, the html report is generated in `reports` directory in the project by default.

## Configuration
The HTML report plugin can be configured by the properties set in the `env/default.properties` file in the project.

The configurable properties are:

#### *gauge_reports_dir*
* Specifies the path to the directory where the execution reports will be generated.

* Should be either relative to the project directory or an absolute path.
By default it is set to `reports` directory in the project

#### ***overwrite_reports***
* Set to `true` if the reports **should be overwritten** on each execution hence maintaining only the latest execution report.

* If set to `false` then a **new report** will be generated on each execution in the reports directory in a nested time-stamped directory.

Default value is `true`
