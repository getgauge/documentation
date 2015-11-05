# Install Language runner

Gauge is a first class command line utility. This means that you only **need** one plugin to use Gauge. That is your language runner.

The langauge runner determines the language you would use to write your tests. Langauge supported currently are [Java](../test_code/java/java.md), [C#](../test_code/csharp/csharp.md) ,and [Ruby](../test_code/ruby/ruby.md).

To install a langauge runner and then initialise the runner, use the following command

####Java
```
gauge install java

```
#### CSharp(C#)
```
gauge install csharp

```
####Ruby
```
gauge install ruby

```

# Check your installation

You can check the version of your plugin and Gauge core by executing the following command.
```
gauge -v
```
If this enlists a version then you're intallation and initialisation is has been successful. Your output will look like this:
```
Gauge version: <version number>

Plugins
-------
language(<version number>)

```

You can read more about plugins, the different types of plugins and how to install and uninstall them.

If you have Gauge and your langauge runner installed, then see how you can [create a Gauge project](creating_a_gauge_project.md).
