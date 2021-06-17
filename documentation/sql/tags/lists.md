---
title: Lists
enterprise: false
---

# Lists

## `cycle`

Produces one of its arguments each time this tag is encountered. The first argument
is produced on the first encounter, the second argument on the second encounter,
and so forth. Once all arguments are exhausted, the tag cycles to the first argument
and produces it again.

In the following example, a player is either in the `odd` or `even` side:
```sql
INSERT INTO players (side) VALUES
  {% for player in game.players %}
    ('{% cycle "odd" "even" %}') {% if forloop.Last %};{% else %},{% endif %}
  {% endfor %}

```

This outputs:
```sql
INSERT INTO players (side) VALUES
  ('odd'),
  ('even'),
  ('odd'),
  ('even');

```

You can use the `{% resetcycle %}` tag to make a `{% cycle %}` tag restart from
its first value when it's next encountered.

## `firstof`

Outputs the first value that is not "false" (*i.e.* exists, is not empty, is not
a false boolean value, and is not a zero numeric value). Outputs nothing if all
the passed variables are "false".

Given the following data:
```json
{
  "choice1": "",
  "choice2": "none",
  "choice3": "",
}

```

And given this template SQL:
```sql
INSERT INTO settings (notifications_mobile) VALUES
  ('{% firstof choice1 choice2 choice3 %}');

```

The output is:
```sql
INSERT INTO settings (notifications_mobile) VALUES
  ('none');

```

The template equivalent only with conditions tags would be:
```sql
INSERT INTO settings (notifications_mobile) VALUES
  {% if choice1 %}
    ('{{ choice1 }}');
  {% elif choice2 %}
    ('{{ choice2 }}');
  {% elif choice3 %}
    ('{{ choice3 }}');
  {% endif %}

```

You can also use a literal string as a fallback value in case all passed variables
are "false":
```sql
INSERT INTO settings (notifications_mobile) VALUES
  ('{% firstof choice1 choice2 choice3 "default" %}');

```

## `for`

Loops over each item in an slice, array or map, making the item available in a
scoped variable.

Accessing the value for each iteration inside a slice, array, map:
```sql
INSERT INTO players (username) VALUES
{% for player in game.players %}
  ('{{ player.username }}'){% if forloop.Last %};{% else %},{% endif %}
{% endfor %}

```

Accessing the key and the value for each iteration inside a map:
```sql
INSERT INTO kv_store (key, value) VALUES
{% for k, v in context %}
  ('{{ k }}', '{{ v }}'){% if forloop.Last %};{% else %},{% endif %}
{% endfor %}

```

### `reversed`

You can loop over a list in reverse by adding `reversed` to a `{% for %}` tag:
```sql
INSERT INTO players (username) VALUES
{% for player in game.players reversed %}
  ('{{ player.username }}'){% if forloop.Last %};{% else %},{% endif %}
{% endfor %}

```

### `forloop.` variables

The for loop sets a number of variables available within the loop:
- `forloop.Counter` is the current iteration of the loop, 1-indexed.
- `forloop.Counter0` is current iteration of the loop, 0-indexed.
- `forloop.Revcounter` is the number of iterations from the end of the loop,
  1-indexed.
- `forloop.Revcounter0` is the number of iterations from the end of the loop,
  0-indexed.
- `forloop.First` is `True` if this is the first time through the loop.
- `forloop.Last` is `True` if this is the last time through the loop.
- `forloop.Parentloop` is for nested loops. This is the loop surrounding the current
  one. It also gives access to the `forloop.` variables of the parent one.
