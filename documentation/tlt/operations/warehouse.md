---
title: Operations on top of data warehouse
enterprise: false
---

# Operations on top of data warehouse

It's possible to run *operations* and *queries* manually from the Blacksmith CLI
outside the ETL process.

This can be achieved with the `run operation` command, which needs at least two
information:
- `scope` is the database to connect to.
- `file` is the SQL file to compile and execute.

In the following example, we run the file located at `./operations/demo.sql`
directly against the destination registered using the `sqlike` under the name
`mypostgres`.
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.sql"

Executing operations:

  -> Executing ./operations/demo.sql...
     Success!

```

**Running a template directly against your database without knowing the compiled
statement can be dangerous.** We strongly advise to first use the `--dryrun` flag,
which compiles the SQL file under a new file named `<operation>.compiled.sql`.
This file is located at the same path than the template one.

If we want to make the same run as the previous one but much more safely, we first
compile the SQL file using the `--dryrun` flag:
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.sql" \
  --dryrun

Compiling operations:

  -> Compiling ./operations/demo.sql...
     Writing SQL at ./operations/demo.compiled.sql...
     Success!

```

After making sure the output SQL is correct, we can then run the compiled statement
instead of the template one:
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.compiled.sql"

Executing operations:

  -> Executing ./operations/demo.compiled.sql...
     Success!

```
