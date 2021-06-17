---
title: Conditions
enterprise: false
---

# Conditions

## `if`, `elif`, `else`

The `{% if %}` tag evaluates a variable, and if that variable is "true" (i.e.
exists, is not empty, and is not a false boolean value) the contents of the block
are output:
```sql
{% if order %}
  INSERT INTO orders (amount, user_id)
    VALUES (
      {{ order.amount | floatformat:2 }},
      '{{ user.id }}'
    );
{% endif %}

```

The `{% if %}` tag may take one or several `{% elif %}` clauses, as well as an
`{% else %}` clause that will be displayed if all previous conditions fail.

### Boolean operators

`{% if %}` tags may use `and`, `or`, or `not` to test a number of variables or to
negate a given variable:

Example if there is an `order` and a `payment`:
```sql
{% if order and payment %}
  INSERT INTO orders (amount, user_id)
    VALUES (
      {{ order.amount | floatformat:2 }},
      '{{ user.id }}'
    );
{% endif %}

```

Example if there is a `user` or a `payment`:
```sql
{% if user or payment %}
  INSERT INTO payments (amount, user_id)
    VALUES (
      {{ payment.amount | floatformat:2 }},
      '{{ user.id }}'
    );
{% endif %}

```

Example if there is no user and a `payment`:
```sql
{% if not user and payment %}
  INSERT INTO payments (amount)
    VALUES (
      {{ payment.amount | floatformat:2 }}
    );
{% endif %}

```

Use of both `and` and `or` clauses within the same tag is allowed, with `and` having
higher precedence than `or`.

`{% if %}` and `{% elif %}` tags may also use the operators `==`, `!=`, `<`, `>`,
`<=`, `>=`, `in`, `not in`, `is`, and `is not`.

### Filters

You can also use filters in the `{% if %}` expression. For example:
```sql
{% if game.players|length >= 100 %}
  REFRESH MATERIALIZED VIEW games_critical;
{% endif %}
```

## `ifchanged`

Check if a value has changed from the last iteration of a loop.

The `{% ifchanged %}` tag is used within a loop. It has two possible uses.

1. Checks its own rendered contents against its previous state if it has changed.
  For example, this `INSERT`s a list of `sessions`, only `INSERT`ing it if it has
  changed:
  ```sql
  {% for session in batch.sessions %}
    {% ifchanged %}
      INSERT INTO sessions (id, user_id, ip, device) VALUES
        ('{{ session.id }}', '{{ session.user_id }}', '{{ session.ip }}', '{{ session.device }}');
    {% endifchanged %}
  {% endfor %}
  ```

2. If given one or more variables, check whether any variable has changed. For
  example, the following `INSERT`s the session only if the `ip` or the `device`
  has changed from the previous iteration:
  ```sql
  {% for session in batch.sessions %}
    {% ifchanged session.ip session.device %}
      INSERT INTO sessions (id, user_id, ip, device) VALUES
        ('{{ session.id }}', '{{ session.user_id }}', '{{ session.ip }}', '{{ session.device }}');
    {% endifchanged %}
  {% endfor %}
  ```

The `{% ifchanged %}` tag can also take an optional `{% else %}` clause that will
be displayed if the value has not changed.
