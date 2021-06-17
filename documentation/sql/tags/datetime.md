---
title: Date & Time
enterprise: false
---

# Date & Time

## `now`

Displays the current date and/or time, using a format according to the [Go's time
format](https://golang.org/pkg/time/#pkg-constants). Such string can contain format
specifiers characters as described in the Go documentation.

```sql
INSERT INTO orders (amount, created_at) VALUES
  ({{ amount }}, '{% now "Mon Jan 2 15:04:05 -0700 MST 2006" %}');

```

In order to harmonize date and time formats in an organization and across systems,
one could create a helper template containing a single format. For example, a
file `./helpers/now.helper` could look like:
```sql
{% now "Mon Jan 2 15:04:05 -0700 MST 2006" %}

```

Now, you can include this template wherever a current date with time is required:
```sql
INSERT INTO orders (amount, created_at) VALUES
  ({{ amount }}, '{% include "./helpers/now.helper" %}');

```

Which outputs:
```sql
INSERT INTO orders (amount, created_at) VALUES
  (12, 'Mon Jun 14 16:59:20 +0200 CEST 2021');

```

This avoids confusion around date and time formats but also simplies the writing
of SQL queries.
