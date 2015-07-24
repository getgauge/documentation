# Data Store to share data
Data (CSharp objects) can be shared in steps defined in different classes at runtime using DataStores exposed by Gauge.

DataStore should be initialized for the first time before using it. It can be ScenarioDataStore, SpecDataStore or SuiteDataStore.

````
DataStoreFactory.ScenarioDataStore.Initialize();
````

There are 3 different types of DataStores based on the lifecycle of when it gets cleared.
### 1. ScenarioStore
This data store keeps values added to it in the lifecycle of the scenario execution. Values are cleared after every scenario executes
````chsarp
// Adding value
DataStore scenarioStore = DataStoreFactory.ScenarioDataStore;
scenarioStore.Add("element-id","67869");

// Fetching Value
DataStore scenarioStore = DataStoreFactory.ScenarioDataStore;
String elementId = (String)scenarioStore.Get("element-id");
````
The key is of type `string` and the data stored can be an object of any type.

````
// Adding value
DataStore scenarioStore = DataStoreFactory.ScenarioDataStore;
List<string> aList = new List<string>();
aList.Add("hello!");
scenarioStore.Add("myList", aList);
````

### 2. SpecStore
This data store keeps values added to it during the lifecycle of the specification execution. Values are cleared after every specification executes

````csharp
// Adding value
DataStore specStore = DataStoreFactory.SpecDataStore;
specStore.Add("element-id","67869");

// Fetching Value
DataStore specStore = DataStoreFactory.SpecDataStore;
String elementId = (String)specStore.Get("element-id");
````

### 3. SuiteStore
This data store keeps values added to it during the lifecycle of entire suite execution. Values are cleared after entire suite execution.


````java
// Adding value
DataStore suiteStore = DataStoreFactory.SuiteDataStore;
suiteStore.Add("element-id","67869");

// Fetching Value
DataStore suiteStore = DataStoreFactory.SuiteDataStore;
String elementId = (String)suiteStore.Get("element-id");
````

````
Note: SuiteStore is not advised to be used when executing specs in parallel.
      The values are not retained between parallel streams of execution.
````

