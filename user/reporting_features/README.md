# Reporting features

The test results report should be easy to comprehend and should be useful for all the stakeholders.

## Html report
  * A comprehensive test results report template prepared in a html format providing the overall summary with drill down of the test cases executed and effort spent during the testing for each stage and feature.
  * It provides the details for the defects found during the run.
  * It indicates the tests by color code - failed(red), passed(green) and skipped(grey).
  * The failure can be analyzed with the stacktrace and screenshot(captures unless overwritten not to).
  * The skipped tests can be analyzed with the given reason.
  * [Custom Report Messages](../language_features/custom_messages.md) allows users to add messages at runtime.

## Xml report
  * This follows the junit format which allows integration with CI.
  * It summarizes the success, failure(with stacktrace) and skipped tests with reason.

Gauge also allows custom reporting plugins to be added if required.
