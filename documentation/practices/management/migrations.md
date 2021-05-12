---
title: Migrations management
enterprise: true
---

# Migrations management

Migrations add a convenient way to evolve your databases' schemas over time.
Blacksmith has a unique set of migration features making the experience as smooth
as possible for data engineers.

Migrations are isolated and not centrally managed into the Blacksmith application.
Each source and each destination has its own migrations. By doing so, you can
leverage the native dialect of each one.

Isolation can go further. Inside a source, each trigger can also have its own
migrations, inheriting the migration logic of its parent source. This to run or
rollback migrations for a specific trigger only.

This works as well for destinations and actions: an action can have its own
migrations isolating these from its parent destination.

## Foreword

### The `sqlike` package

Blacksmith does not enforce the format of the migrations. It is up to the sources
and destinations to implement their own business logic. However, in the following
document we assume you are working with the `sqlike` package, managing migrations
for SQL database(s).

The `sqlike` package does not handle isolation at the trigger and action levels.
Migrations are handled per source and destination only.

As of now, the `sqlike` package is the only one implementing the required interfaces
for working with migrations. Others could exist later!

### Scope definition

Isolation is made possible thanks to a *scope*. A scope is defined by the user
when working with migrations. Multiple scopes can be defined.

To work with migrations for a specific source, the scope format is
`source:<mysource>`, or `s:<mysource>`. The scope is specific to the source only,
and does not apply to its triggers.

To work with migrations for a specific trigger, the scope format is
`source/trigger:<mysource>/<mytrigger>`, or `s/t:<mysource>/<mytrigger>`

To work with migrations for a specific destination, the scope format is
`destination:<mydestination>`, or `d:<mydestination>`. The scope is specific to
the destination only, and does not apply to its actions.

To work with migrations for a specific action, the scope format is
`destination/action:<mydestination>/<myaction>` or `d/a:<mydestination>/<myaction>`.

### Locking scopes

When working with migrations, we strongly advise to enable the `supervisor` adapter.
When enabled, the CLI leverages it to ensure only one migration is being acknowledged,
running, or rolling back within the same scope at the same time. This allows
distributed teams to safely work together with no access collisions.

If a user tries to run or rollback migrations for a scope already in use, the CLI
will return an error like this one:
```bash
ERRO[2021-04-29T18:09:17+02:00] cli: Failed to manage migrations
  validations="[{Migration lock not acquired by the supervisor [migrations destination:warehouse]}]"

```

## Generating migrations

A migration can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate migration --name add_user_level

```

This will generate the recommended files for a migration, inside the working
directory.

If you prefer, you can generate a migration inside a directory with the `--path`
flag:
```bash
$ blacksmith generate migration --name add_user_level \
  --path ./relative/path/migrations

```

The generated file names should look like this:
- `20201012150000.add_user_level.up.sql`
- `20201012150000.add_user_level.down.sql`

A migration file has several properties, including:
- `Version`: The version number is a 14 length character holding the current
  timestamp, formatted with `YYYYMMDDHHMISS`.
  
  **Example:** `20201012150000`.

- `Name`: The slugify name of the migration.

  **Example:** `add_user_level`.

- `Direction`: The direction of the run. It is either `up` or `down`.

## Acknowledging migrations

Once a migration is written, you can acknowledge it so the `wanderer` adapter can
keep track of its status:
```bash
$ blacksmith migrations ack

```

If you wish to acknowledge migrations for a given scope, you can add a `--scope`
flag:
```bash
$ blacksmith migrations ack --scope "destination:sqlike(crm)"

```

Since working with migrations can be multi-scoped, you can add multiple scopes
as follow:
```bash
$ blacksmith migrations ack --scope "destination:sqlike(crm)" \
  --scope "destination:sqlike(warehouse)"

```

In the example above, all new migrations written for the destinations `crm` and
`warehouse` will be acknowledged.

## Running migrations

Once migrations are acknowledged, they can be run with:
```bash
$ blacksmith migrations run

```

To only run migrations within a scope, you need to add the desired scope:
```bash
$ blacksmith migrations run --scope "destination:sqlike(crm)"

```

Or running migrations within a multi-scope:
```bash
$ blacksmith migrations run --scope "destination:sqlike(crm)" \
  --scope "destination:sqlike(warehouse)"

```

You can also run migrations, within a single or multi-scope, until a given version
is reached:
```bash
$ blacksmith migrations run --scope "destination:sqlike(crm)" \
  --version 20200930071321

```

Migrations are executed ordered by version from oldest to newest. If an error
occured while executing a migration, the others will not be executed. You first
need to fix the failing migration to be able to run it again along the others
awaiting to run.

## Rolling back migrations

Mistakes can happen. This is why you sometimes need to *rollback* a migration, or
a suite of migrations.

When rolling back, a `version` flag must be provided. All migrations will be
rolled back until the version is reached.

For example, to rollback every migrations down to a specific version:
```bash
$ blacksmith migrations rollback --version 20200930071321

```

Or rolling back migrations within a single scope:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope "destination:sqlike(crm)"

```

Or rolling back migrations within a multi-scope:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope "destination:sqlike(crm)" \
  --scope "destination:sqlike(warehouse)"

```

Migrations are rolled back ordered by version from newest to oldest. If an error
occured while rolling back a migration, the others will not be rolled back. You
first need to fix the failing migration to be able to rollback it again along the
others awaiting to be rolled back.

## Discarding migrations

When rolled back, a migration is not discarded. Which means you can still update
its `up` logic and try to run it again. However, if you need to discard a migration
once it is successfully rolled back, you can add the `--discard` flag as follow:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope "destination:sqlike(crm)" \
  --discard

```

This will mark the migration as "discarded" so it will not be run again.
