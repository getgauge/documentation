# Installation
------
Plugins are installed using the flag `install`, this checks our plugin repository and downloads them. Read the [plugins directory](../troubleshooting/installation.md) to know more about the plugin installation directory and how to customize it. 

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
