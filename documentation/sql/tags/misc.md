---
title: Misc.
enterprise: false
---

# Misc.

## `with`

Caches a complex variable under a simpler name. This is useful when accessing an
"expensive" object.

For example:
```sql
{% with total=game.players|length %}
  SELECT id FROM games
  WHERE number_players >= {{ total }};
{% endwith %}

```

The populated variable (in the example above, `total`) is only available between
the `{% with %}` and `{% endwith %}` tags.
