# Gauge Configuration

All the Gauge specific internal configurations are stored in `gauge.properties` file present in your Gauge install location (`GAUGE_ROOT`). These properties are key value pairs.

##gauge_repository_url

This property is set to an url, which acts as plugin repository for Gauge.

**Please do not change this url** or it will break the installation and update of Gauge plugins.

```
gauge_repository_url = http://raw.github.com/getgauge/gauge-repository/master
```

##gauge_templates_url

This property is set to an url, which acts as template repository for Gauge.

**Please do not change this url** or it will break the project initialization using templates.

```
gauge_templates_url = https://dl.bintray.com/gauge/Templates
```

###runner_connection_timeout

This property sets the timeout in milliseconds for making a connection to the langauge runner.

```
runner_connection_timeout = 30000
```

###plugin_connection_timeout

This property sets the timeout in milliseconds for making a connection to plugins (except language runner plugins).
```
plugin_connection_timeout = 10000
```

###plugin_kill_timeout

This property sets the timeout in milliseconds for a plugin to stop after a kill message has been sent.
```
plugin_kill_timeout = 10000
```

###runner_request_timeout

This property sets the timeout in milliseconds for requests from the langauge runner.

If the size of the project is too big, Gauge may timeout before the runner returns the response message. This value can be configured accordingly.

```
runner_request_timeout = 10000
```
