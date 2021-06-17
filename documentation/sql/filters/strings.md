---
title: Strings
enterprise: false
---

# Strings

## `add`

Adds the argument to the value.

For example:
```sql
{{ "Hello "|add:"John" }}

```

Will return:
```
Hello John
```

## `capfirst`

Capitalizes the first character of the value. If the first character is not a
letter, this filter has no effect.

For example:
```sql
{{ "hi there!"|capfirst }}

```

Will return:
```
Hi there!
```

## `cut`

Removes all values of the argument from the given string.

For example:
```sql
{{ "Hello world"|cut: " " }}

```

Will return:
```
Helloworld
```

## `default`

If value evaluates to `False`, uses the given default. Otherwise, uses the value.

For example:
```sql
{{ ""|default:"n/a" }}

```

Will return:
```
n/a
```

## `length`

Returns the length of the string.

For example:
```sql
{{ "Jane"|length }}

```

Will return:
```
4
```

## `length_is`

Returns `True` if the value's length is the argument, or `False` otherwise.

For example:
```sql
{{ "Jane"|length_is:"3" }}

```

Will return:
```
False
```

## `lower`

Converts a string into all lowercase.

For example:
```sql
{{ "This is AWESOME"|lower }}

```

Will return:
```
this is awesome
```

## `pluralize`

Returns a plural suffix if the value is not 1, "1", or an object of length 1. By
default, this suffix is "s".

For example:
```sql
customer{{ 4|pluralize }}
cherr{{ 0|pluralize:"y,ies" }}
cherr{{ 3|pluralize:"y,ies" }}

```

Will return:
```
customers
cherry
cherries
```

## `slugify`

Converts to ASCII. Converts spaces to hyphens. Removes characters that aren't
alphanumerics, underscores, or hyphens. Converts to lowercase. Also strips leading
and trailing whitespace.

For example:
```sql
{{ "Jane Doe"|slugify }}

```

Will return:
```
jane-doe
```

## `title`

Converts a string into titlecase by making words start with an uppercase character
and the remaining characters lowercase.

For example:
```sql
{{ "My first post"|title }}

```

Will return:
```
My First Post
```

## `upper`

Converts a string into all uppercase.

For example:
```sql
{{ "Hello world"|upper }}

```

Will return:
```
HELLO WORLD
```

## `wordcount`

Returns the number of words.

For example:
```sql
{{ "Hello Jane, how are you today?"|wordcount }}

```

Will return:
```
6
```
