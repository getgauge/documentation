# Enum as Step parameter

The constant values of an Enum data type can be used as parameters to a Step. However, the type of parameter should match the Enum name itself in step implementation.

Step:

```
* Navigate towards "SOUTH"
```

Implementation:

{% codetabs name="Java", type="java" -%}
public enum Direction {
    NORTH, SOUTH, EAST, WEST;
}

@Step("Navigate towards <direction>")
public void navigate(Direction direction) {
    // code here
}
{%- endcodetabs %}



