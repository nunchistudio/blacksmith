---
title: Your first query
enterprise: false
---

# Your first query

Now that the data properly flowed from the source's webhook to the data warehouse,
we are able to query the data. In the following example, we run the query of the
`./warehouse/queries/all-users.sql` file and download the result as CSV
and JSON files:
```bash
$ blacksmith run query \
  --scope "destination:sqlike(warehouse)" \
  --file "./warehouse/queries/all-users.sql"
  --csv --json

Compiling & Executing queries:

  -> Compiling & Executing ./warehouse/queries/all-users.sql...
     Writing CSV at ./warehouse/queries/all-users.csv...
     Writing JSON at ./warehouse/queries/all-users.json...
     Success!

```

By default, every `*.csv` and `*.json` files in `./warehouse/queries` are ignored
by Git. If you wish to version these files, you need to remove the following lines
in the `.gitignore`:
```
warehouse/queries/*.csv
warehouse/queries/*.json
```

**Congratulations!** You just created and played with your self-hosted data
engineering platform from end-to-end.
