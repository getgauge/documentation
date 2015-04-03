#### Installation

Installing a plugin happens via Gauge. Gauge's `properties` file contains a URL that points to the `Gauge-Repository`, which holds meta information, which Gauge uses to determine the highest compatible version of the plugin.
This makes plugin installation a two step process:
- Download `<language>-install.json`, determine the highest compatible plugin version
- Verify if version is already installed, if not, download the plugin using the URL specified in the `<language>-install.json`.
