# Data Store
 Data (Objects) can be shared in steps defined in different classes at runtime using DataStores exposed by Gauge.

There are 3 different types of DataStores based on the lifecycle of when it gets cleared.

### 1. ScenarioStore
This data store keeps values added to it in the lifecycle of the scenario execution. Values are cleared after every scenario executes

{% codetabs name="Java", type="java" -%}
// Import Package
import com.thoughtworks.gauge.datastore.*;

// Adding value
DataStore scenarioStore = DataStoreFactory.getScenarioDataStore();
scenarioStore.put("element-id", "455678");

// Fetching Value
DataStore scenarioStore = DataStoreFactory.getScenarioDataStore();
String elementId = (String) scenarioStore.get("element-id");
{%- language name="C#", type="csharp" -%}
using Gauge.CSharp.Lib;

// Adding value
var scenarioStore = DataStoreFactory.ScenarioDataStore;
scenarioStore.Add("element-id", "455678");

// Fetching Value
var scenarioStore = DataStoreFactory.ScenarioDataStore;
var elementId = (string) scenarioStore.Get("element-id");

// avoid type cast by using generic Get<T>
var anotherElementId = scenarioStore.Get<string>("element-id");
{%- language name="Ruby", type="ruby" -%}
// Adding value
scenario_store = Gauge::DataStoreFactory.scenario_datastore;
scenario_store.put("element-id", "455678");

// Fetching Value
scenario_store = Gauge::DataStoreFactory.scenario_datastore;
element_id = scenario_store.get("element-id");
{%- endcodetabs %}

### 2. SpecStore
This data store keeps values added to it during the lifecycle of the specification execution. Values are cleared after every specification executes

{% codetabs name="Java", type="java" -%}
// Import Package
import com.thoughtworks.gauge.datastore.*;

// Adding value
DataStore specStore = DataStoreFactory.getSpecDataStore();
specStore.put("key", "455678");

// Fetching value
DataStore specStore = DataStoreFactory.getSpecDataStore();
String elementId = (String) specStore.get("key");
{%- language name="C#", type="csharp" -%}
using Gauge.CSharp.Lib;

// Adding value
var specStore = DataStoreFactory.SpecDataStore;
specStore.Add("element-id", "455678");

// Fetching Value
var specStore = DataStoreFactory.SpecDataStore;
var elementId = (string) specStore.Get("element-id");

// avoid type cast by using generic Get<T>
var anotherElementId = specStore.Get<string>("element-id");
{%- language name="Ruby", type="ruby" -%}
// Adding value
spec_store = Gauge::DataStoreFactory.spec_datastore;
spec_store.put("element-id", "455678");

// Fetching Value
spec_store = Gauge::DataStoreFactory.spec_datastore;
element_id = spec_store.get("element-id");
{%- endcodetabs %}

### 3. SuiteStore
This data store keeps values added to it during the lifecycle of entire suite execution. Values are cleared after entire suite execution.

> Warning: SuiteStore is not advised to be used when executing specs in parallel.
      The values are not retained between parallel streams of execution.

{% codetabs name="Java", type="java" -%}
// Import Package
import com.thoughtworks.gauge.datastore.*;

// Adding value
DataStore suiteStore = DataStoreFactory.getSuiteDataStore();
suiteStore.put("element-id", "455678");

// Fetching value
DataStore suiteStore = DataStoreFactory.getSuiteDataStore();
String elementId = (String) suiteStore.get("element-id");
{%- language name="C#", type="csharp" -%}
using Gauge.CSharp.Lib;

// Adding value
var suiteStore = DataStoreFactory.SuiteDataStore;
suiteStore.Add("element-id", "455678");

// Fetching Value
var suiteStore = DataStoreFactory.SuiteDataStore;
var elementId = (string) suiteStore.Get("element-id");

// avoid type cast by using generic Get<T>
var anotherElementId = suiteStore.Get<string>("element-id");
{%- language name="Ruby", type="ruby" -%}
// Adding value
suite_store = Gauge::DataStoreFactory.suite_datastore;
suite_store.put("element-id", "455678");

// Fetching Value
suite_store = Gauge::DataStoreFactory.suite_datastore;
element_id = suite_store.get("element-id");
{%- endcodetabs %}
