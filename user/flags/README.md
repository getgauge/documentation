# Flags

Gauge has **first-class command line support**. If you have gauge installed, running `gauge` in terminal gives you command usage and all the flags it supports.

Here is a list of flags that can be used with `gauge`.

| Flag  | Description |
| ------| ----------- |
|--add-plugin=""| Adds the specified non-language plugin to the current project.|
|  --api-port=""  |    Specifies the api port to be used.<br> Eg: ``` gauge --daemonize --api-port 7777```|
|  --daemonize=false    |  Run as a **daemon**.|
|  --dir="."|Set the working directory for the current command, accepts a path relative to current directory.|
|  --env="" | Specifies the **environment**. If not specified, default will be used.|
|  --file, -f=""    | Installs the plugin from zip file. This is used with --install.<br> Eg: ```gauge --install java -f ZIP_FILE```|
|  --format=""       |Formats the specified spec files. <br>Eg: ```gauge --format specs``` |
|  -g, --group   |  Specify which group of specs to execute based on -n flag.<br>Eg: ```gauge -n=3 -g=1 specs```|
|  --init=""| Initializes project structure in the current directory.<br> Eg: ```gauge --init java```|
|  --install="" |  Downloads and installs a plugin.<br> Eg: ```gauge --install java```|
|  --log-level="" |     Set level of logging to debug, info, warning, error or critical.|
|  -n          |Specify number of **parallel execution** streams.<br>Eg: ```gauge -p -n=4 specs``` <br> This executes the specs in 4 parallel streams.|
|  --parallel, -p=false         |    Execute specs in parallel.|
|  --plugin-args="" | Specify additional arguments to the plugin. This is used together with --add-plugin.|
|  --plugin-version=""     |         Version of plugin to be installed. This is used with --install.|
|  --refactor=""   |      **Refactor/Rephrase** steps.<br>Eg: ```gauge --refactor old_step new_step```|
|  --simple-console=false     | Removes colouring and simplifies the console output.|
|  --sort, -s=false          |       Run specs in Alphabetical Order. <br>Eg: ```gauge -s specs```|
|  --table-rows=""             |     Executes the specs and scenarios only for the selected rows of data table.<br> Eg: ```gauge --table-rows "1-3" specs/hello.spec ```|
|  --tags=""                |        Executes the specs/ scenarios tagged with given tags. <br>Eg: ```gauge --tags tag1,tag2 specs```<br>This filtering can also be done based on **tag expression**.<br>Eg: ```gauge --tags "tag1 & tag2" specs```|
|  --update=""           |Updates a plugin. <br>Eg: ```gauge --update java```|
|  -v, --version, -version=false   | Print the current version and exit.<br> Eg: ```gauge --version```|
|  --verbose=false         |         Enable verbose logging for debugging.|
