# Logs

* Gauge logs are created under the ```logs``` directory in the project.
* 2 log files are created
  - **gauge.log** - logs for test execution
  - **api.log**   - logs for gauge core api exposed for plugins
* To customize logs directory set the ```logs_directory``` property in the ```env/default/default.properties``` file to a custom logs directory path.

````
logs_directory = my_logs_dir
````
* For **non-project specific actions** like plugin installation log files are created in the following location.

````
 Windows - %APPDATA%\gauge\logs
 MacOS*  - <user_home>/.gauge/logs
 Linux   - <user_home>/.gauge/logs
````
