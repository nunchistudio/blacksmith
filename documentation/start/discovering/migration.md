---
title: Your first migration
enterprise: true
---

# Your first migration

For simplicity, the data warehouse is the same PostgreSQL database created in
the `Docker-compose.yml` file, which is also used as `store`, `supervisor`, and
`wanderer` adapters. **It should not be the case when running in production.**

You should access it with the following details:
- **Host:** `localhost`
- **Port:** `5432`
- **User:** `<appname>`
- **Password:** `<appname>`
- **Database:** `<appname>`

Before inserting data coming from a source, the warehouse must have the desired
schema. The migration file `init_users.up.sql` contains the the appropriate SQL
statement to create the table for storing our users.

If you do not use Blacksmith Enterprise Edition, you can run the content of the
migration file using the tool of your choice.

First, we need to acknowledge (_a.k.a._ ack) the migration:
```bash
$ blacksmith migrations ack

1 new migration within the scope acknowledged and awaiting to run:
  - destination:sqlike(warehouse): 20210617112246.init_users.up.sql

You can now run these migrations by running:

  $ blacksmith migrations run [--scope <scope>]


Documentation:

  https://nunchi.studio/blacksmith/practices/migrations

```

This allows the `wanderer` adapter to keep track of the new migration files.

As described in the output of the previous command, we can run the migration for
executing the `up` logic:
```bash
$ blacksmith migrations run

1 migration within the scope will be run:
  - destination:sqlike(warehouse): 20210617112246.init_users.up.sql (status: acknowledged)

Blacksmith will run the migrations as shown above. Do you confirm?
Only 'yes' will be accepted to confirm.
> yes

Executing migrations:

  -> Executing 20210617112246.init_users.up.sql...
     Success!

INFO[2021-06-17T13:27:51+02:00] Migrations successfully run!       

```

Using the software of your choice, you should now see a table `users` in your
PostgreSQL database (schema `public`).
