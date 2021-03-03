---
title: Migrations management
enterprise: true
---

# Migrations management

Migrations add a convenient way to evolve your databases' schemas over time.
Blacksmith has a unique set of migration features making the experience as smooth
as possible for data engineers.

Migrations are isolated and not centrally managed into the Blacksmith application.
Each source and each destination has its own `Migrate` function, leveraging the
native dialect of each one. For example, you write SQL-ish for SQL-like databases.

Each source and each destination can therefore have a `Migrations` function,
returning a collection of migrations to manage.

Isolation can go further. Inside a source, each trigger can also have its own
`Migrations` function, inheriting the `Migrate` logic of its parent source,
allowing to run or rollback migrations for a specific trigger only.

This works as well for destinations and actions: an action can have its own
`Migrations` function isolating its migrations from its parent destination, but
inheriting the `Migrate` logic of its destination.

## Scope definition

Isolation is made possible thanks to a "scope". A scope is defined by the user
when working with migrations. Multiple scopes can be defined.

To work with migrations for a specific source, the scope format is
`source:<mysource>`, or `s:<mysource>`. The scope is specific to the source only,
and not its triggers.

To work with migrations for a specific trigger, the scope format is
`source/trigger:<mysource>/<mytrigger>`, or `s/t:<mysource>/<mytrigger>`

To work with migrations for a specific destination, the scope format is
`destination:<mydestination>`, or `d:<mydestination>`. The scope is specific to
the destination only, and not its actions.

To work with migrations for a specific action, the scope format is
`destination/action:<mydestination>/<myaction>` or `d/a:<mydestination>/<myaction>`.

> When working with migrations inside Blacksmith, we strongly advise to enable
  the supervisor adapter. When enabled, the CLI leverages the supervisor to make
  sure only one migration is being acknowledged, running, or rolling back, within
  the same scope. This allows distributed teams to safely work together with no
  access collisions.

## How migrations work

When generating a source, trigger, destination, or action you can add migration
management with the `--migrations` flag.

On a source and a destination, it adds the following function:
```go
Migrate(*wanderer.Toolkit, *wanderer.Migration) error

```

The `Migrate` function holds the migration logic for the source or the destination,
which will be inherited by their triggers or actions if they handle migrations
separately.

To manage migrations, the `Migrations` function is defined as follow:
```go
Migrations() ([]*wanderer.Migration, error)

```

## Migration format

A migration has several properties, including:
- `Version`: The version number is a 14 length character holding the current
  timestamp, formatted with `YYYYMMDDHHMISS`. Example `20201012150000`.
- `Name`: The slugify name of the migration. Example `add_user_level`.
- `Direction`: The direction of the run. It is either `up` or `down`.

## Generating migrations

A migration can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate migration --name mymigration

```

This will generate the recommended files for a migration, inside the working
directory.

If you prefer, you can generate a migration inside a directory with the `--path` flag:
```bash
$ blacksmith generate migration --name mymigration \
  --path ./relative/path/migrations

```

## The `sqlike` package

Writing the logic inside the `Migrate` and `Migrations` can be tedious. This is
why we provide the `sqlike` package, which provides helpers when working with
SQL-like databases.

Using the `sqlike` package, a `Migrate` function should look like this:
```go
func (s *MySource) Migrate(tk *wanderer.Toolkit, migration *wanderer.Migration) error {

  db, err := sql.Open("<driver-name>", "<driver-url>")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    os.Exit(1)
  }

  defer db.Close()
  return sqlike.RunMigration(db, filepath.Join("sources", "api", "migrations"), migration)
}

```

The `sqlike.RunMigration` function will run (or rollback) the migration in a
transaction using the standard SQL package.


Using the `sqlike` package, a `Migrations` function should look like this:
```go
func (s *MySource) Migrations(tk *wanderer.Toolkit) ([]*wanderer.Migration, error) {
  return sqlike.LoadMigrations(filepath.Join("sources", "postgres", "migrations"))
}

```

The `sqlike.LoadMigrations` function will load every SQL files from a directory
with a migration file name, which is `<version>.<name>.<direction>.sql`.

## Acknowledging migrations

Once a migration is written, you can acknowledge it so the wanderer can keep
track of its runs:
```bash
$ blacksmith migrations ack

```

If you wish to acknowledge for a given scope, you can add the `--scope` flag:
```bash
$ blacksmith migrations ack --scope source:crm

```

Since working with migrations can be multi-scoped, you can add multiple scopes
as follow:
```bash
$ blacksmith migrations ack --scope source:crm \
  --scope destination:warehouse

```

## Running migrations

Once migrations are acknowledged, they can can be run with:
```bash
$ blacksmith migrations run

```

To only run migrations within a scope, you need to add the desired scope:
```bash
$ blacksmith migrations run --scope source:crm

```

Or running migrations within a multi-scope:
```bash
$ blacksmith migrations run --scope source:crm \
  --scope destination:warehouse

```

You can also run migrations, within a single or multi-scope, until a given version
is reached:
```bash
$ blacksmith migrations run --scope source:crm \
  --version 20200930071321

```

## Rolling back migrations

Mistakes can happen. This is why you sometimes need to "rollback" a migration, or
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
  --scope source:crm

```

Or rolling back migrations within a multi-scope:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope source:crm \
  --scope destination:warehouse

```

When rolled back, a migration is not discarded. Which means you can still update
it and try to run it again. However, if you need to discard a migration once it
is successfully rolled back, you can add the `--discard` flag as follow:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope source:crm \
  --discard

```

This will mark the migration as "discarded" so it will not be run again.
