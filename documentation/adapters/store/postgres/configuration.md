The PostgreSQL adapter is compatible with any PostgreSQL wire compatible database
and can work with any kind of extensions.

#### Usage

```go
&blacksmith.Options{
  Store: &store.Options{
    From:       "postgres",
    Connection: "postgres://user:password@127.0.0.1/database",
  },
}
```

#### Environment variables

- `POSTGRES_STORE_URL`: The PostgreSQL URL to use for the store adapter. If
  `Options.Store.Connection` is set, it will override and be used in replacement
  of this environment variable.

  Example: `postgres://user:password@127.0.0.1/database`

#### SQL migration

```sql
CREATE SCHEMA IF NOT EXISTS blacksmith_store;

CREATE TABLE IF NOT EXISTS blacksmith_store.events (
  id VARCHAR(27) PRIMARY KEY,
  source TEXT NOT NULL,
  trigger TEXT NOT NULL,
  context JSONB,
  data JSONB,
  sent_at TIMESTAMP WITHOUT TIME ZONE,
  received_at TIMESTAMP WITHOUT TIME ZONE,
  ingested_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_store.jobs (
  id VARCHAR(27) PRIMARY KEY,
  destination TEXT NOT NULL,
  action TEXT NOT NULL,
  context JSONB,
  data JSONB,
  parent_job_id VARCHAR(27) REFERENCES blacksmith_store.jobs (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  event_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.events (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_store.transitions (
  id VARCHAR(27) PRIMARY KEY,
  attempt INT4 NOT NULL,
  state_before TEXT,
  state_after TEXT NOT NULL,
  error JSONB,
  event_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.events (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  job_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.jobs (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
```
