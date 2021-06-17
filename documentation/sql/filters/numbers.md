---
title: Numbers
enterprise: false
---

# Numbers

## `add`

Adds the argument to the value.

For example:
```sql
{{ 2|default:4 }}
```

Will return:
```
6
```

## `default`

If the value evaluates to `0` or is missing, uses the given default. Otherwise,
uses the value.

For example:
```sql
{{ 0|default:"1" }}
```

Will return:
```
1
```

## `divisibleby`

Returns `True` if the value is divisible by the argument.

For example:
```sql
{{ 9|divisibleby:"3" }}
```

Will return:
```
True
```

## `filesizeformat`

Formats the value like a "human-readable" file size.

For example:
```sql
{{ 123456789|filesizeformat }}
```

Will return:
```
117.7 MB
```

## `floatformat`

When used without an argument, rounds a floating-point number to one decimal place
– but only if there's a decimal part to be displayed.

For example:
```sql
{{ 34.00000|floatformat }}
{{ 34.23234|floatformat }}
{{ 34.26000|floatformat }}
```

Will return:
```
34
34.2
34.3
```

If used with a numeric integer argument, floatformat rounds a number to that many
decimal places.

For example:
```sql
{{ 34.00000|floatformat:"3" }}
{{ 34.23234|floatformat:"3" }}
{{ 34.26000|floatformat:"3" }}
```

Will return:
```
34.000
34.232
34.260
```

If the argument passed to floatformat is negative, it will round a number to that
many decimal places – but only if there's a decimal part to be displayed.

For example:
```sql
{{ 34.00000|floatformat:"-3" }}
{{ 34.23234|floatformat:"-3" }}
{{ 34.26000|floatformat:"-3" }}
```

Will return:
```
34
34.232
34.260
```

## `intcomma`

Converts an integer or float (or a string representation of either) to a string
containing commas every three digits.

For example:
```sql
{{ 450000|intcomma }}
```

Will return:
```
450,000
```

## `ordinal`

Converts an integer to its ordinal as a string.

For example:
```sql
{{ 3|ordinal }}
```

Will return:
```
3rd
```
