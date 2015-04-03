#### Communication

Gauge and Runners exchange data using one of the below:

- Command line arguments: Gauge tells the runner to run `Setup` or `Start` phase by passing in `--setup` or `--start` argument respectively, while invoking the runner.
- Environment Variables: Used to exchange configuration properties. Gauge loads all values in `default.properties` as well as `environment.properties` from the project environment, and passes these to the Runner.
- `GAUGE_INTERNAL_PORT` : This environment variable holds a TCP port, which Gauge and the Runner exchange [core messages](https://github.com/getgauge/gauge-proto/blob/master/doc/gauge-proto-doc.md#messages.proto).
- `GAUGE_API_PORT`: This environment variable holds a TCP port, which Gauge and the Runner exchange [api messages](https://github.com/getgauge/gauge-proto/blob/master/doc/gauge-proto-doc.md#api.proto).

The data passed via TCP use [Protocol Buffers](https://developers.google.com/protocol-buffers/).
