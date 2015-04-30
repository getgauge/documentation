# Installation

Plugins are installed using the flag `install`, this checks our plugin repository and downloads them.

```gauge --install <plugin-id>```

**Example:**
```
gauge --install java

gauge --install html-report
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

### Plugin install location
* **Windows** - `%APPDATA%\gauge\plugins` on windows
* **MacOS**   - `<home>/.gauge/plugins`
* **Linux**   - `<home>/.gauge/plugins`


##Adding plugins to a project

Once plugins are installed, they can be added to the project by
using the `add-plugin` flag

```
gauge --add-plugin <plugin-id>
```

```
gauge --add-plugin xml-report
```

##Updating plugins

To update a plugin to the latest verion of it, use the `--update` flag. This downloads the latest plugin from our plugin repository.

`gauge --update <plugin-id>`

Example:
````
gauge --update java

gauge --update ruby
````
