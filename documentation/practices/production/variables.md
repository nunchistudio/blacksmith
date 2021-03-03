---
title: Environment variables
enterprise: false
---

# Environment variables

An application's config is everything that is likely to vary between environments:
staging, production, developer environments, etc. To avoid writing and versioning
private secrets such as database credentials in code, some config can be passed
using environment variables.

Environment variables can be set by loading a `.env` file when the application is
bootstrapped.  To simplify this best practice, Blacksmith generates `.env` and
`.env.example` files when creating a new application. The `.env` file is added
to the `.gitignore` to ensure it is not versioned alongside your code.

A `.env` file contains a collection of variables following the format:
```txt
KEY=value

```

A `.env` file for a Blacksmith application can contain configuration for the
adapters as well as the license details for the Enterprise Edition:
```txt
BLACKSMITH_LICENSE_KEY=omSYn27xhL6OKwvpuHrf
BLACKSMITH_LICENSE_TOKEN=Li5eVI3RJ6Chi9w3QSHf
NATS_SERVER_URL=nats://localhost:4222
POSTGRES_STORE_URL=postgres://app:app@localhost:5432/app
POSTGRES_WANDERER_URL=postgres://app:app@localhost:5432/app

```
