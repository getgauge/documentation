# IDE Plugin

IDE Plugin will send requests to Gauge and Gauge will perform actions based on request and send the response back.

#### Autocomplete

|Request                |Purpose                                              |Response                |
     |-----------------------|-----------------------------------------------------|------------------------|
     |`GetAllStepRequest`    |Fetches all the steps in the gauge project as a list.|`GetAllStepResponse`    |
     |`GetAllConceptsRequest`|Fetches all the concepts in the gauge project as a list.|`GetAllConceptsResponse`|

*Combine the results of the above requests to show all the steps and concepts present in gauge project.*

#### Format Specification.

|Request             |Purpose                      |Response|
     |--------------------|-----------------------------|--------|
     |`FormatSpecsRequest`|Request contains list of spec file names to format, Response will contain errors & warnings occurred on formatting.|`FormatSpecsResponse`|

#### Gets language plugin libs location.

|Request                          |Purpose                                                        |Response                          |
     |---------------------------------|---------------------------------------------------------------|----------------------------------|
     |`GetLanguagePluginLibPathRequest`|Request to get the location of language plugin's Lib directory.|`GetLanguagePluginLibPathResponse`|

#### Get Gauge Installation Path

|Request                     |Purpose                                                     |Response                     |
     |----------------------------|------------------------------------------------------------|-----------------------------|
     |`GetInstallationRootRequest`|Request to get the Root Directory of the Gauge installation.|`GetInstallationRootResponse`|

#### Get Project Root

|Request                     |Purpose                                                     |Response                     |
     |----------------------------|------------------------------------------------------------|-----------------------------|
     |`GetProjectRootRequest`|Request to get the Root Directory of the project.|`GetProjectRootResponse`|

#### Refactoring

|Request                    |Purpose                                                   |Response                    |
     |---------------------------|----------------------------------------------------------|----------------------------|
     |`PerformRefactoringRequest`|request to gauge to perform a step name refactoring on the project.|`PerformRefactoringResponse`|

#### Extract to Concept

|Request                |Purpose                                           |Response                |
     |-----------------------|--------------------------------------------------|------------------------|
     |`ExtractConceptRequest`|Request to perform Extract to Concept refactoring.|`ExtractConceptResponse`|

#### Get Parsed Step

|Request              |Purpose                                                                          |Response              |
     |---------------------|---------------------------------------------------------------------------------|----------------------|
     |`GetStepValueRequest`|Gets the step value which contains param, parsed text for a particular step name.|`GetStepValueResponse`|

#### Get Specs in Project

|Request             |Purpose                                                                                                   |Response             |
     |--------------------|----------------------------------------------------------------------------------------------------------|---------------------|
     |`GetAllSpecsRequest`|Request to get all Specs in the project and returns a collection of Specs that are defined in the project.|`GetAllSpecsResponse`|

