---
title: Guides for TLT with SQL
enterprise: false
---

# Guides for TLT with SQL

Please refer to the navigation on the left sidebar for TLT guides.

These guides focus on accomplishing TLT with Blacksmith. This is usually the work
of **data** or **analytics engineers**. Given the schema in the "Getting started",
we only focus on the **TLT** part with SQL here:
![TLT with Blacksmith](/images/blacksmith/guides-tlt.png)

It's possible to run two kinds of work on top of your data warehouse: *operations*
and *queries*.

An **operation** executes a statement without returning any rows. When running
an operation, it is automatically wrapped inside a transaction to ensure it is
either entirely commited or rolled back if any error occured. Operations should
be used for:
- Creating or refreshing views (materialized or not).
- Inserting, updating, or deleting data.

An operation should never be used for evolving the database schema, which could
impact software engineers. In this scenario, you should take advantage of [database
migrations](/blacksmith/practices/management/migrations).

A **query** executes a statement that returns rows, typically a `SELECT`. Queries
should be used for:
- `SELECT`ing samples of data.

By running these using Blacksmith, you are able to leverage templating SQL with
a Pythonistic approach. SQL files can follow the [Django built-in template tags
and filters](/blacksmith/sql):
```sql
INSERT INTO users (id, first_name, last_name, username, email, created_at)
  VALUES (
    '{{ user.id }}',
    '{{ user.first_name | capfirst }}',
    '{{ user.last_name | upper }}',
    '{{ user.username | slugify }}',
    '{{ user.email | lower }}',

    {% if user.created_at %}
      '{{ user.created_at }}'
    {% else %}
      '{% now "2006-01-02 15:04:05" %}'
    {% endif %}
  );

{% if order %}
  INSERT INTO orders (id, amount, user_id, created_at)
    VALUES (
      '{% ksuid %}',
      {{ order.amount | floatformat:2 }},
      '{{ user.id }}',

      {% if order.created_at %}
        '{{ order.created_at }}'
      {% else %}
        '{% now "2006-01-02 15:04:05" %}'
      {% endif %}
    );
{% endif %}

```

Any source and destination can leverage the `warehouse` package to benefit templating
SQL. The `sqlike` module — which offers a unique and consistent approach for Loading
data into any and every SQL databases — makes use of this package.

This allows data engineers to follow best practices already used in software
engineering, such as code versioning and the DRY (Don't Repeat Yourself) principle.
You can version original *and* compiled files within your source code management,
depending on your needs and workflow.
