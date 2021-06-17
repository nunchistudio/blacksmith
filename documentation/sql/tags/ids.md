---
title: Unique IDs
enterprise: false
---

# Unique IDs

## `ksuid`

Generates a new KSUID, the unique identifiers used in Blacksmith.

```sql
INSERT INTO orders (id) VALUES
  ('{% ksuid %}');

```
