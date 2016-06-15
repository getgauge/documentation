# IntelliJ Gauge Plugin
This plugin helps to create and execute Gauge project in IntelliJ.

## Setting up

 * Install the intelliJ SDK
 * Install necessary plugins Grammar-kit, PsiViewer
  http://www.jetbrains.org/intellij/sdk/docs/tutorials/custom_language_support/prerequisites.html
 * Create a run configuration (Plugin -> Use Classpath of module -> Choose the current module), hit run
 * A new instance of IntelliJ will be launched with the plugin installed

## Lexing and parsing

 * The lexing rules are defined in `_SpecLexer.flex` and `_SpecLexer.java` is generated using JFlex
 * The grammar rules are defined in `specification.bnf` and `SpecParser.java` is generated using grammar kit
   SpecParser builds an structure with PsiElements and we can add custom behavior for each element (eg. step has get specstep has getStepValue)

## Available Features

 * Project Creation
 * Syntax Highlighting
 * Auto completion
 * Navigation from step to implementation
 * Quick Fix for unimplemented steps
 * Find Usages
 * Formatting Specification
 * Refactoring


## Interacting with the core

Auto completion,step value extraction and other api calls are made from the plugin. Each module is associated with an api connection(GaugeConnection).
This Gauge connection is created during the module initialization phase after the gauge process is started in background. Gauge connection talks to the core using protobuff.

## Some useful links

[Antlr guy's tips](https://github.com/antlr/jetbrains/blob/master/doc/plugin-dev-notes.md)

[Psi Cook book](http://www.jetbrains.org/intellij/sdk/docs/basics/psi_cookbook.html)
