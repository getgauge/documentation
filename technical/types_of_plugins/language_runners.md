## Language Runners

Gauge allows one to write implementations in multiple languages, and this is made possible via such plugins. Language runners are alive for the full cycle of an execution.

### Responsibilites

- Execute hooks (when requested to)
- Execute method corresponding to a step
- Send execution response back to Gauge
- Refactor method signatures with changes defined by Gauge
- Initialize a project with template
- Read/write to datastore (currently, datastore is intra process, hence parallel processes cannot use a shared datastore)
- Holds reference to Step text and Method / execution block.

### Consituents

- **Runner binary** : This is responsible for scanning the test library and also invoke methods when requested to.
- **Skeleton files (templates)** : These files create a sample `hello world` project that can be used as a quick-start for creating a new project.
- **Metadata file (usually `<language>.json`)** : Holds meta information about the plugin. Also holds the commands to invoke for `init` and `run` phases of the plugin.

#### Format of a Metadata file(`<langauge>.json`)
```json
{
    "id": "{plugin-name}",
    "version": "{plugin-version}",
    "description": "{plugin-description}",
    // Commands to execute pre-installation
    "preInstall": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Commands to execute post-installation
    "postInstall": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Commands to invoke for run phase of the plugin
    "run": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Commands to invoke for init phase of the plugin
    "init": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Commands to execute pre-uninstallation
    "preUnInstall": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Commands to execute post-uninstallation
    "postUnInstall": {
        "windows": [],
        "linux": [],
        "darwin": []
    },
    // Location of libraries for the plugin relative to plugin directory
    "lib": "{libs}",
    "gaugeVersionSupport": {
        "minimum": "{minimum gauge version supported}",  // mandatory
        "maximum": "{maximum gauge version supported}"   // optional
    }
}
```

### Installation

Installing a plugin happens via Gauge. Gauge's `properties` file contains a URL that points to the `Gauge-Repository`, which holds meta information, which Gauge uses to determine the highest compatible version of the plugin.
This makes plugin installation a two step process:
- Download `<language>-install.json`, determine the highest compatible plugin version
- Verify if version is already installed, if not, download the plugin using the URL specified in the `<language>-install.json`.
