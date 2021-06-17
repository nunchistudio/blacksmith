---
title: Lists
enterprise: false
---

# Lists

The examples in this document share `players` for the variable `value`, which is:
```json
{
  "players": [
    "John",
    "Jane",
    "Alex",
    "Max"
  ]
}

```

## `first`

Returns the first item in a list. The list must be a slice or array of basic types
such as strings, numbers, or booleans.

For example:
```sql
{{ players|first }}

```

Will return:
```
John
```

## `join`

Returns a stringified list with items seperated by the param. The list must be a
slice or array of basic types such as strings, numbers, or booleans.

For example:
```sql
{{ players|join:", " }}

```

Will return:
```
John, Jane, Alex, Max
```

## `last`

Returns the last item in a list. The list must be a slice or array of basic types
such as strings, numbers, or booleans.

For example:
```sql
{{ players|last }}

```

Will return:
```
Max
```

## `length`

Returns the length of the list.

For example:
```sql
{{ players|length }}

```

Will return:
```
4
```

## `length_is`

Returns `True` if the value's length is the argument, or `False` otherwise.

For example:
```sql
{{ players|length_is:"4" }}

```

Will return:
```
True
```

## `random`

Returns a random item from the given list. The list must be a slice or array of
basic types such as strings, numbers, or booleans.

For example:
```sql
{{ players|random }}

```

*Could* return:
```
Alex
```
