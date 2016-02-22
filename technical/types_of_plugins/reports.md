## Reports

Reports plugins come in towards the end of the execution. In Gauge's context the definition of Report is the result of a test run.

#### HTML Report

The HTML report holds the output of the test execution, with screenshots (on failure).

HTML report is a plugin, and it listens to Gauge's messages. On Receiving [`SuiteExecutionResult`](https://github.com/getgauge/gauge-proto/blob/master/doc/gauge-proto-doc.md#gauge.messages.SuiteExecutionResult), the plugin serializes the message to JSON.

The view of this report is an angular-js application that binds the JSON output. The result looks like [this](https://gauge-reports-ruby.herokuapp.com/)

#### JUnit style report

Conceptually everything is similar to the HTML report, except that instead of a JSON and angular-js view, this plugin serializes the result into a JUnit format XML.
