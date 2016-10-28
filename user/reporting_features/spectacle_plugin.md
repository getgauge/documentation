# Spectacle

This is a Gauge plugin that generates static HTML from Specification/Markdown files. Ability to filter specifications and scenarios are available.

## Installation
To install :
```
gauge --install spectacle
```

To install a specific version of spectacle plugin use the `--plugin-version` flag.
````
gauge --install spectacle --plugin-version 0.0.2
````
__Offline Installation__ :

If plugin should be installed from a zipfile instead of downloading from plugin repository, use the `--file` or `-f` flag.
````
gauge --install spectacle --file ZIP_FILE_PATH
````
Download the plugin zip from the [Github Releases](https://github.com/getgauge/spectacle/releases)

## Export to HTML

Run the following command to export to HTML in a Gauge project

```
gauge --docs spectacle <path to specs dir>
```

**Sample Spectacle Report**

![Sample spectacle report](images/spectacle.png "Sample spectacle report")

**Filter Specification/Scenario based on Tags**

Tags allow you to filter the specs and scenarios. Add the tags to the textbox in the report to view all the specs and scenarios which are labelled with certain tags. Tag expressions with operators `|`, `&`, `!` are supported.

In the following image the specs/scenarios are filtered using a tag expression(`refactoring & !api`).

![Filter Specification/Scenario](images/filter.png "Filter Specification/Scenario")