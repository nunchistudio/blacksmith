---
title: Queries on top of data warehouse
enterprise: false
---

# Queries on top of data warehouse

Queries can be used to `SELECT` rows and download the result inside a CSV or JSON
file.

This can be achieved with the `run query` command, which needs at least two
information:
- `scope` is the database to connect to.
- `file` is the SQL file to compile and execute.

In the following example, we run the file located at `./queries/demo.sql` directly
against the destination registered using the `sqlike` under the name `mypostgres`.
We download the rows both as CSV and JSON files.
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --csv --json

Compiling & Executing queries:

  -> Compiling & Executing ./queries/demo.sql...
     Writing CSV at ./queries/demo.csv...
     Writing JSON at ./queries/demo.json...
     Success!

```

**Running a template directly against your database without knowing the compiled
statement can be tedious.** We strongly advise to first use the `--dryrun` flag,
which compiles the SQL file under a new file named `<query>.compiled.sql`. This
file is located at the same path than the template one.

If we want to make the same run as the previous one but much more safely, we first
compile the SQL file using the `--dryrun` flag:
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --dryrun

Compiling queries:

  -> Compiling ./queries/demo.sql...
     Writing SQL at ./queries/demo.compiled.sql...
     Success!

```

After making sure the output SQL is correct, we can then run the compiled statement
instead of the template one, and download the result both as CSV and JSON:
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.compiled.sql" \
  --csv --json

Compiling & Executing queries:

  -> Compiling & Executing ./queries/demo.compiled.sql...
     Writing CSV at ./queries/demo.csv...
     Writing JSON at ./queries/demo.json...
     Success!

```
