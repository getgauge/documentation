# Ruby

## Step implementations

Step implementations in ruby are method calls calling the ***step*** method. The [step name template](../step_name_template.md) is passed as the first parameter to the method followed by a code block which takes the necessary number of parameters.

###Examples


####1. Simple step

**Step name**
```
* Say "hello" to "gauge"
```

**Ruby Implementation:**
````ruby
step 'Say <greeting> to <product name>' do |greeting, name|
 # Code for the step
end
````



####2. Step with table
**Step:**

````
Create following "admin" users
|id |   name  |
|---|---------|
|123| prateek |
|456| vishnu  |
|789|navaneeth|
````

**Ruby Implementation:**
````ruby
step 'Create following <role> users <table>' do |role, table|
  puts table.rows
  puts table.columns
end


````
* Here **table** is a custom data structure defined by gauge-ruby.

## Multiple Step names
The same implementation can point to multiple step names. To do this pass all the [step name template](../step_name_template.md) as parameters to the method call.


####Example

***Step Names :***
````
* Create a user "prateek"
* Create another user "vishnu"
````
***Implementation :***

````ruby
step "Create a user <username>", "Create another user <username>" do |username|
 #step code
end

````
