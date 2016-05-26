# Installation
------
Plugins are installed using the flag `install`, this checks our plugin repository and downloads them.

```gauge --install <plugin-id>```

**Example:**
```
gauge --install java

```

To install a specific version of a plugin use the `--plugin-version` flag.
````
gauge --install ruby --plugin-version 0.0.2
````
### Offline Installation
If plugin should be installed from a zipfile instead of downloading from plugin repository, use the `--file` or `-f` flag.
````
gauge --install java --file ZIP_FILE_PATH
````
Download the latest version of plugin from the `Releases` section of the respective repository. See [plugin list](list.md) to find the repositories.

### Custom Plugin Install location

By default the plugins are stored at `%APPDATA%\gauge\plugins` for windows and `~/.gauge/plugins` in mac and linux.

To install plugins at different location, set `GAUGE_HOME` environment variable to the custom location. After setting the `GAUGE_HOME` env, run the install command. The plugin will get installed at the `GAUGE_HOME` custom location.

##Adding plugins to a project

Once plugins are installed, they can be added to the project by
using the `add-plugin` flag

`gauge --add-plugin <plugin-id>`

```
gauge --add-plugin xml-report
```

##Updating plugins

To update a plugin to the latest verion of it, use the `--update` flag. This downloads the latest plugin from our plugin repository.

`gauge --update <plugin-id>`

Example:
````
gauge --update java
````
To update a plugin to a specific version, use the `--plugin-version` flag.
````
gauge --update java --plugin-version 0.3.2
````
You can also update all the installed plugins by running
````
gauge --update-all
````
##Uninstalling plugins
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


## Troubleshooting

Plugins are installed in the `.gauge/plugins` directory in user's home. You can check this directory to manually install / uninstall plugins as well as to verify the installed plugins.

The plugin installation directory for various operating systems are listed below.

* **Windows:** `%APPDATA%\.gauge\plugins`
* **MacOS:** `~/.gauge/plugins`
* **Linux:** `~/.gauge/plugins`
