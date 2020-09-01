The PostgreSQL adapter is compatible with any PostgreSQL wire compatible database
and can work with any kind of extensions.

#### Configuration

```go
&blacksmith.Options{
  Wanderer: &wanderer.Options{
    From:       "postgres",
    Connection: "postgres://user:password@127.0.0.1/database",
  },
}
```

#### Environment variables

- `POSTGRES_WANDERER_URL`: The PostgreSQL URL to use for the wanderer adapter.
  If `Options.Wanderer.Connection` is set, it will override and be used in replacement
  of this environment variable.

  Example: `postgres://user:password@127.0.0.1/database`

#### SQL migration

Before using the adapter, you first need to run the following migration:

```sql
CREATE SCHEMA IF NOT EXISTS blacksmith_wanderer;

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.migrations (
  id VARCHAR(27) PRIMARY KEY,
  version VARCHAR(14) NOT NULL,
  interface_kind TEXT NOT NULL,
  interface_string TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.directions (
  id VARCHAR(27) PRIMARY KEY,
  filename TEXT NOT NULL,
  direction TEXT NOT NULL,
  sha256 BYTEA NOT NULL,
  migration_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.transitions (
  id VARCHAR(27) PRIMARY KEY,
  attempt INTEGER NOT NULL,
  state_before TEXT,
  state_after TEXT NOT NULL,
  error JSONB,
  migration_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  direction_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.directions (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX migrations_version
  ON blacksmith_wanderer.migrations (version, interface_kind, interface_string);
```
