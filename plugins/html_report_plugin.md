# HTML Report plugin

Reports are generated using `html-report` plugin. By default html-report is added to the project. When the specs are executed, the html report is generated in `reports` directory.


There are two configurable properties added to the default.properties in the env directory.

* The path to the gauge reports directory. Should be either relative to the project directory or an absolute path
`gauge_reports_dir = reports`

* Set as false if gauge reports should not be overwritten on each execution. A new time-stamped directory will be created on each execution.
`overwrite_reports = true`
