# C# Project files

When creating a new Gauge C# project, the csharp specific project files created in the project are:

````
├── foo.csproj
├── foo.sln
├── manifest.json
├── packages.config
├── StepImplementation.cs
│
├── env
│   └───default
│           default.properties
│
├───packages
    └───<Nuget Package Binaries>
├───Properties
│       AssemblyInfo.cs
│
└───specs
        hello_world.spec
````

###packages.config
For `nuget`. Contains the dependencies for Gauge. One can add more to this list, depending on your project needs.

### StepImplementation.cs
Contains the implementations for the sample steps defined in `hello_world.spec`.

### default.properties
This defines default configurations for gauge csharp runner plugin. Currently the configuration parameters are:

- `gauge_reports_dir` - The path to the gauge reports directory. Should be either relative to the project directory or an absolute path
- `overwrite_reports` - Set as false if gauge reports should not be overwritten on each execution. A new time-stamped directory will be created on each execution. This is `true` by default.
