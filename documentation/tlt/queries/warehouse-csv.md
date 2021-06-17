---
title: Downloading data into CSV file
enterprise: false
---

# Downloading data into CSV file

When running a query, you can download the rows returned as a CSV file with the
`csv` flag. It creates or overwrites the file `<query>.csv`. This file is located
at the same path than the SQL one.
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --csv

Compiling & Executing queries:

  -> Compiling & Executing ./queries/demo.sql...
     Writing CSV at ./queries/demo.csv...
     Success!

```

The CSV follows the format desired when passing a `data` flag. This makes the CSV
reusable as a data source when running operations!
