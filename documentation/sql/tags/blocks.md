---
title: Blocks
enterprise: false
---

# Blocks

## `block`

Defines a block that can be overridden by child templates. See the `{% extends %}`
tag for more details usage.

## `extends`

Signals that this template extends a parent template.

A template file must contain at most one `{% extends %}` tag. This shall be the first
line of the template file.

Template inheritance allows you to build a base "skeleton" template that contains
all the common elements of your SQL query and defines blocks that child templates
can override.

The following file is named `view.sql` and is located in the `base` sub-directory.
It creates a view and exposes three blocks:
- `temp` which defaults to `TEMPORARY`;
- `name` with no default value;
- `query` with no default value.

```sql
CREATE {% block temp %}TEMPORARY{% endblock %} VIEW
  IF NOT EXISTS {% block name %}{% endblock %}
  AS {% block query %}{% endblock %}
  WITH DATA;

```

In the following template, the `view.sql` is used as base file. We set a `name`,
override the default value of `temp`, and define a `query` using the `{% block %}`
tag.
```sql
{% extends "./base/view.sql" %}

{% block name %}myview{% endblock %}
{% block temp %}MATERIALIZED{% endblock %}

{% block query %}
  SELECT ...
{% endblock %}

```

This outputs:
```sql
CREATE MATERIALIZED VIEW
  IF NOT EXISTS myview
  AS SELECT ...
  WITH DATA;

```

Normally the template name is relative to the template query's directory. A string
argument may also be a relative path starting with `./` or `../`.

## `include`

Loads a template and renders it with the current context. This is a way of
"including" other templates within a template.

This example includes the contents of the template located at `./includes/user.sql`:
```sql
{% include "./includes/user.sql" %}

```

Normally the template name is relative to the template query's directory. A string
argument may also be a relative path starting with `./` or `../` as described in
the `{% extends %}` tag.

### `with`

You can pass additional context to the template using keyword arguments:
```sql
{% include "./includes/user.sql" with username="janedoe" %}

```

The template at `./includes/user.sql` could therefore call a variable named
`username`, as follow:
```sql
SELECT id, username, first_name, last_name
  FROM users
  WHERE username = '{{ username }}';

```

Which compiles to:
```sql
SELECT id, username, first_name, last_name
  FROM users
  WHERE username = 'janedoe';

```

## `macro`

Macros are leveraged to encapsulate logic to perform repeatable actions. Macros
are great for creating reusable components when we find ourselves copy pasting
around same lines of text and code *within a same template*.

Macros can take arguments or be used without them.

In the following example, we define a new macro `insertUser` to `INSERT` some
`users` using the `{% macro %}` tag. Then we call it twice, for each user we want
to `INSERT`:
```sql
{% macro insertUser(username, is_admin="false") %}
INSERT INTO users (id, username, is_admin) VALUES
  ('{% ksuid %}', '{{ username }}', {{ is_admin }});
{% endmacro %}

{{ insertUser("johndoe") }}
{{ insertUser("janedoe", "true") }}

```

Which compiles to:
```sql
INSERT INTO users (id, username, is_admin) VALUES
  ('1twYPzTuIoMfwdzZfG1kMeZkVxJ', 'johndoe', false);

INSERT INTO users (id, username, is_admin) VALUES
  ('1twYQ19zLfbkK4AdW5NSR0NtGyy', 'janedoe', true);

```
