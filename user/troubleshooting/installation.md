# Troubleshooting

## Gauge Installation
### Windows

* **GAUGE_ROOT** [environment variable](http://www.computerhope.com/issues/ch000549.htm) should be set to the gauge installation path.

* The default installation location is %ProgramFiles%\gauge.

* ```Gauge_install_location\bin``` should be in PATH to run from command line. See [how to add folder to path](http://windowsitpro.com/systems-management/how-can-i-add-new-folder-my-system-path)
* Gauge plugins are installed at ```%APPDATA%\gauge\plugins``` directory.

* [APPDATA](http://windows.microsoft.com/en-in/windows-8/what-appdata-folder) directory is usually located at ```C:\Users\USER_NAME\AppData\Roaming```.

### MacOs X
* The default installation location is ```/usr/local/```.

* On custom installation location Set the environment variable ```GAUGE_ROOT``` to custom_install_location.

* ```usr/local/bin/``` should be in PATH.

* Run `brew update` before installing the latest version of gauge.

* If installation is failing [Upgrade homebrew](https://github.com/Homebrew/homebrew/blob/master/share/doc/homebrew/FAQ.md#faq)

* Gauge plugins are installed under ```~/.gauge/plugins``` directory

### Linux
* The default installation location is ```/usr/local/```.

* On custom installation location Set the environment variable ```GAUGE_ROOT``` to custom_install_location.

* ```usr/local/bin/``` or ```custom_install_location/bin``` should be in PATH.

* Gauge plugins are installed under ```~/.gauge/plugins``` directory

## Plugin installation

* If **[plugin installation](../plugins/installation.md)** fails due to a network connection issue, you can **manually download** the plugin distributable zip and install it using the ```-f``` flag.

````
gauge --install plugin_name -f path_to_zip_file

eg:
gauge --install html-report -f html-report-1.0.3-darwin.x86.zip
````
* Find the plugin zip files under Releases section of the plugin github repositories. See the [gauge plugins list](../plugins/list.md) for plugin repositories details.


## Plugins directory

Plugins are installed in the `.gauge/plugins` directory in user's home. You can check this directory to manually install / uninstall plugins as well as to verify the installed plugins.

The plugin installation directory for various operating systems are listed below.

* **Windows:** `%APPDATA%\.gauge\plugins`
* **MacOS:** `~/.gauge/plugins`
* **Linux:** `~/.gauge/plugins`

### Custom Plugin Install location

By default the plugins are stored at `%APPDATA%\gauge\plugins` for windows and `~/.gauge/plugins` in mac and linux.

To install plugins at different location, set `GAUGE_HOME` environment variable to the custom location. After setting the `GAUGE_HOME` env, run the install command. The plugin will get installed at the `GAUGE_HOME` custom location.
