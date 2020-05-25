The PostgreSQL adapter is compatible with any PostgreSQL-like database and can
work with any kind of extensions.

#### Usage

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

```sql
CREATE SCHEMA IF NOT EXISTS blacksmith_wanderer;

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.locks (
  id VARCHAR(27) PRIMARY KEY,
  acquired_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  released_at TIMESTAMP WITHOUT TIME ZONE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.migrations (
  id VARCHAR(27) PRIMARY KEY,
  version VARCHAR(14) NOT NULL,
  interface_kind TEXT NOT NULL,
  interface_string TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.jobs (
  id VARCHAR(27) PRIMARY KEY,
  filename TEXT NOT NULL,
  direction TEXT NOT NULL,
  sha256 BYTEA NOT NULL,
  migration_ID VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
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
  lock_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.locks (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  migration_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  job_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.jobs (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX migrations_version
  ON blacksmith_wanderer.migrations (version, interface_kind, interface_string);
```
