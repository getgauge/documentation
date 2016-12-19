# Uninstallation

## Uninstalling plugins

Plugins can be uninstalled using the `uninstall` flag. The command is

`gauge --uninstall <plugin-id>`

Example:
````
gauge --uninstall java
````

To uninstall a specific version of the plugin, use the `--plugin-version` flag.
````
gauge --uninstall java --plugin-version 0.3.2
````

## Uninstalling Gauge

To uninstall Gauge, run the following commands:

#### OS X/Linux
```
rm -rf /usr/local/bin/gauge /usr/local/share/gauge /usr/local/share/gauge_screenshot ~/.gauge
```

If Gauge is installed in custom location, user will have to remove corresponding files/directory. 

#### Windows
Run the executable `uninst.exe` found in Gauge install location.

More on Gauge install location can be found [here](../troubleshooting/installation.md).