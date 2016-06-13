# Step alias

Multiple Step names for the same implementation. The number and type of parameters for all the steps names must match the number of parameters on the implementation.

## Use case
There may be situations where while authoring the specs, you may want to express the same functionality in different ways in order to make the specs more readable.

### Example 1

```
User Creation
=============
Multiple Users
--------------
* Create a user "user 1"
* Verify "user 1" has access to dashboard
* Create another user "user 2"
* Verify "user 2" has access to dashboard
```

In the scenario named Multiple Users, the underlying functionality of the first and the third step is the same. But the way it is expressed is different. This helps in conveying the intent and the functionality more clearly. In such situations like this, step aliases feature should be used so that you can practice DRY principle at code level, while ensuring that the functionality is expressed clearly.

#### Implementation

{% codetabs name="Java", type="java" -%}
public class Users {

    @Step({"Create a user <user_name>", "Create another user <user_name>"})
    public void helloWorld(String user_name) {
        // create user user_name
    }
}
{%- language name="C#", type="csharp" -%}
public class Users {

    [Step({"Create a user <user_name>", "Create another user <user_name>"})]
    public void HelloWorld(string user_name) {
      // create user user_name
    }
}
{%- language name="Ruby", type="ruby" -%}
step 'Create a user <user_name>',
     'Create another user <user_name>' do |user_name|
     // create user user_name
end
{%- endcodetabs %}


### Example 2
```
User Creation
-------------
* User creates a new account
* A "welcome" email is sent to the user

Shopping Cart
-------------
* User checks out the shopping cart
* Payment is successfully received
* An email confirming the "order" is sent
```

In this case, the underlying functionality of the last step (sending an email) in both the scenarios is the same. But it is expressed more clearly with the use of aliases. The underlying step implementation could be something like this.

#### Implementation

{% codetabs name="Java", type="java" -%}
public class Users {

    @Step({"A <email_type> email is sent to the user", "An email confirming the <email_type> is sent"})
    public void helloWorld(String email_type) {
        // Send email of email_type
    }
}
{%- language name="C#", type="csharp" -%}
public class Users {

    [Step({"A <email_type> email is sent to the user", "An email confirming the <email_type> is sent"})]
    public void HelloWorld(string email_type) {
        // Send email of email_type
    }
}
{%- language name="Ruby", type="ruby" -%}
step 'A <email_type> email is sent to the user',
     'An email confirming the <email_type> is sent' do |email_type|

    email_service.send email_type

end
{%- endcodetabs %}
