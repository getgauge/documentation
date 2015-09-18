# Features

### Creating a new Gauge Project

- Create a new Project of Type "Class Library".
- Add a Nuget reference to "Gauge.CSharp.Lib" using the Nuget Package Manager.

This should setup a new Gauge project, and add the required meta data for Gauge to execute this project.


### Syntax Highlighting

Gauge specs are in [Markdown](http://daringfireball.net/projects/markdown/syntax) syntax. This plugin highlights Specifications, Scenarios, Steps and Tags.

Steps with missing implementation are also highlighted.

![syntax highlighting](ide_support/visual_studio_screenshots/features/Syntax_highlighting.png
"syntax highlighting")

### Autocomplete

This plugin hooks into VisualStudio Intellisense, and brings in autocompletion of Step text. The step texts brought in is a union of steps already defined, concepts defined, and step text from implementation.

Hint: Hit <kbd>Ctrl</kbd> + <kbd>Space</kbd> to bring up the Intellisense menu.

![AutoComplete](ide_support/visual_studio_screenshots/features/AutoComplete.png
"Auto Complete")

### Navigation

Jump from Step text to it's implementation.

Usage: `Right Click` -> `Go to Declaration` or hit <kbd>F12</kbd>

### Smart Tag

Implement an unimplemented step - generates a method template, with a `Step` attribute having this Step Text.

![Quick fix](ide_support/visual_studio_screenshots/features/quickfix/QuickFix.gif
"Quick Fix")

### Test Runner

Open the Test Explorer : `Menu` -> `Test` -> `Windows` -> `Test Explorer`
All the scenarios in the project should be listed. Run one or more of these tests.

![Test Explorer](ide_support/visual_studio_screenshots/features/TestExplorer.png
"Test Explorer")
