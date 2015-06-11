# Data Store
 Data (Java objects) can be shared in steps defined in different classes at runtime using DataStores exposed by Gauge.

There are 3 different types of DataStores based on the lifecycle of when it gets cleared.
### 1. ScenarioStore
This data store keeps values added to it in the lifecycle of the scenario execution. Values are cleared after every scenario executes
````java
// Adding value
DataStore scenarioStore = DataStoreFactory.getScenarioDataStore();
scenarioStore.put("element-id", "455678");

// Fetching Value
DataStore scenarioStore = DataStoreFactory.getScenarioDataStore();
String elementId = (String) scenarioStore.get("element-id");
````

### 2. SpecStore
This data store keeps values added to it during the lifecycle of the specification execution. Values are cleared after every specification executes

````java
// Adding value
DataStore specStore = DataStoreFactory.getSpecDataStore();
specStore.put("key", "455678");

// Fetching value
DataStore specStore = DataStoreFactory.getSpecDataStore();
String elementId = (String) specStore.get("key");
````

### 3. SuiteStore
This data store keeps values added to it during the lifecycle of entire suite execution. Values are cleared after entire suite execution.


````java
// Adding value
DataStore suiteStore = DataStoreFactory.getSuiteDataStore();
suiteStore.put("element-id", "455678");

// Fetching value
DataStore suiteStore = DataStoreFactory.getSuiteDataStore();
String elementId = (String) suiteStore.get("element-id");
````

````
Note: SuiteStore is not advised to be used when executing specs in parallel.
      The values are not retained between parallel streams of execution.
````

