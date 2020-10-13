CREATE SCHEMA IF NOT EXISTS blacksmith_wanderer;

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.migrations (
  id VARCHAR(27) PRIMARY KEY,
  version TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  scope TEXT NOT NULL,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.transitions (
  id VARCHAR(27) PRIMARY KEY,
  state_before TEXT,
  state_after TEXT NOT NULL,
  error JSONB,
  migration_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX migrations_version
  ON blacksmith_wanderer.migrations (version, scope);
