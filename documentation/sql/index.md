---
title: SQL template syntax
enterprise: false
---

# SQL template syntax

Blacksmith exposes a Python / Django-syntax like templating-language when working
with SQL files. This is a convenient way to generate SQL dynamically interpolating
the template with contextual data and returning the compiled SQL file.

The templating-language involves four constructs.

**Variables** are surrounded by `{{` and `}}` like this:
```sql
INSERT INTO orders (amount, user_id) VALUES
  ({{ order.amount }}, '{{ user.id }}');

```

**Tags** provide arbitrary logic in the rendering process. Tags are surrounded
by `{%` and `%}` like this:
```sql
{% if order %}
  INSERT INTO orders (amount, user_id) VALUES
    ({{ order.amount }}, '{{ user.id }}');
{% endif %}

```

**Filters** transform the values of variables with tag arguments. They look like
this:
```sql
INSERT INTO orders (amount, user_id) VALUES
  ({{ order.amount|floatformat:2 }}, '{{ user.id }}');

```

**Comments** look like this:
```django
{# This won't be compiled #}

```
