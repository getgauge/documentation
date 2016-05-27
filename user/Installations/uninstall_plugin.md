#Uninstalling plugins

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
