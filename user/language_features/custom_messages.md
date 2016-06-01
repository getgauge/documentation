# Custom messages in reports

Custom messsages/data can be added to execution reports using the below API from the step implementations or hooks.

These messages will appear under steps in the execution reports.

{% codetabs name="Java", type="java" -%}
Gauge.writeMessage("Custom message for report");

String id = "4567";
Gauge.writeMessage("User id is %s", id);

{%- language name="C#", type="csharp" -%}
GaugeMessages.WriteMessage("Custom message for report");

var id = "4567";
GaugeMessages.WriteMessage("User id is {0}", id);
{%- language name="Ruby", type="ruby" -%}
Gauge.write_message("Custom message for report")

id = "4567"
Gauge.write_message("User id is" + id)
{%- endcodetabs %}
