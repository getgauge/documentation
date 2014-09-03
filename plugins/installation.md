# Installation

Plugins are installed using the flag `install`, this checks our plugin repository and downloads them.

eg
```
gauge --install <plugin-id>
```

since runners are also plugins, they can be installed the same way.

Plugins are placed in `<home>/.gauge/plugins` on *nix and `%APPDATA%\gauge\plugins` on windows

##Adding plugins to project

Once plugins are installed, they can be added to the project by
using the `add-plugin` flag

```
gauge --add-plugin <plugin-id>
```
additional plugin information can be passed using `plugin-args`

eg: adding webdriver version 2.37

```
gauge --add-plugin webdriver --plugin-args="selenium_version=2.37"
```

