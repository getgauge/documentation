# Go
[Go](http://www.go.cd/), is a continuous integration and deployment tool.

##Integrating Gauge with Go

* [Setup a new pipeline on Go](http://www.go.cd/documentation/user/current/configuration/quick_pipeline_setup.html)

    ![pipeline](images/Gauge_Pipeline.png "gauge pipeline")
* [Download](http://getgauge.io/get-started/index.html) and Install Gauge on the Go Agents
* Install the required gauge [language plugin](../../../installations/install_language_runners.md) on the Go agents.

### Create execution task

* Create a new task which will run `gauge specs`.
* If you want to run only a subset of specs, you can use [tags](../execution_types/tagged_execution.md). Eg. ```gauge --tags "tag1 & tag2" specs```
* Adding a flag `-p` runs them in [parallel](../execution_types/parallel_execution.md).
* Run against specific [environments](../managing_environments.md) using the ```--env``` flag
* See the [Gauge CLI](../../cli/README.md) for list of all flags that can be used.

    ![configuring](images/Configuring_Gauge.png "adding new task")

### Reports

* Gauge generates **html-reports** after execution which can be configured in Go by adding a new artifact in Artifacts tab.

    ![artifact](images/Configuring_Artifacts.png "artifact")


* Artifacts can be viewed in the artifacts tab.

    ![artifact](images/Arifacts.png "artifact")

* **Console output** can be seen while execution of job and reports can be seen after execution.

     ![console](images/Console_Output.png "console")

* You can also add a [custom tab](http://www.go.cd/documentation/user/current/configuration/managing_artifacts_and_reports.html#using-tabs) to view your html reports generated.
