---
title: Downloading data into JSON file
enterprise: false
---

# Downloading data into JSON file

When running a query, you can download the rows returned as a JSON file with the
`json` flag. It creates or overwrites the file `<query>.json`. This file is located
at the same path than the SQL one.
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --json

Compiling & Executing queries:

  -> Compiling & Executing ./queries/demo.sql...
     Writing JSON at ./queries/demo.json...
     Success!

```

The rows are not exported as a top-level array but inside a `rows` key alongside
the `count` of rows returned. This makes the JSON reusable as a data source when
running operations!
```json
{
  "count": 12,
  "rows": [
    {
      "id": "1234567890",
      "first_name": "John",
      "last_name": "DOE",
      "username": "john-doe",
      "email": "johndoe@example.com",
      "created_at": "2021-06-02T16:59:30Z",
    },

    [...]

  ]
}

```
