# Plugin Installation
Plugins are installed using the flag `install`, this checks our plugin repository and downloads them.
```gauge --install <plugin_name>```

**Example:**
```
gauge --install java

```

To install a specific version of a plugin use the `--plugin-version` flag.
```
gauge --install ruby --plugin-version 0.0.2
```
### Offline Installation
If plugin should be installed from a zip file instead of downloading from plugin repository, use the `--file` or `-f` flag.
```
gauge --install java --file ZIP_FILE_PATH
```
Download the latest version of plugin from the `Releases` section of the respective repository. See [plugin list](../plugins/README.md) to find the repositories.

##Adding plugins to a project

Once plugins are installed, they can be added to the project by
using the `add-plugin` flag

`gauge --add-plugin <plugin_name>`

```
gauge --add-plugin xml-report
```

##Updating plugins

To update a plugin to the latest version of it, use the `--update` flag. This downloads the latest plugin from our plugin repository.

`gauge --update <plugin_name>`

Example:
```
gauge --update java
```
To update a plugin to a specific version, use the `--plugin-version` flag.
```
gauge --update java --plugin-version 0.3.2
```
You can also update all the installed plugins by running
```
gauge --update-all
```

Read the [Installation troubleshooting](../troubleshooting/installation.md) for more.
