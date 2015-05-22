# Configuration

Certain Gauge specific internal configurations can be made in **gauge.properties** file present in your Gauge install location (GAUGE_ROOT). These properties are key value pairs.

###gauge_repository_url

This property is set to an url, which acts as plugin repository for Gauge.

```
gauge_repository_url = http://raw.github.com/getgauge/gauge-repository/master
```
###api_refresh_interval

Gauge refreshes the api cache at a particular frequency set as value to this property. This property takes the value in milliseconds.


```
api_refresh_interval = 3000
```

###runner_connection_timeout

This property sets the timeout for making a connection to the langauge runner. The value is taken in milliseconds.

```
runner_connection_timeout = 25000
```

###plugin_connection_timeout

This property configures the timeout for making a connection to plugins.
```
plugin_connection_timeout = 10000
```

###plugin_kill_timeout

This property sets the timeout in milliseconds for a plugin to stop after a kill message has been sent.
```
plugin_kill_timeout = 10000
```

###runner_request_timeout

This property configures the timeout in milliseconds for requests from the langauge runner.

If the size of the project is too big, Gauge may timeout before the runner returns the response message. This value can be configured accordingly.

```
runner_request_timeout = 10000
```
