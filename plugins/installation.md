# Installation

Plugins are installed using the flag `install`, this checks our plugin repository and downloads them.

e.g.
```
gauge --install <plugin-id>
```

To install a specific version of a plugin use the `--plugin-version` flag.
````
gauge --install ruby --plugin-version 0.0.2
````

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
Additional plugin information can be passed using `plugin-args` flag.
These depend on the plugin being installed.

e.g: On the Webdriver plugin, Adding webdriver version 2.37

```
gauge --add-plugin webdriver --plugin-args="selenium_version=2.37"
```

