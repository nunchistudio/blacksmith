---
title: blacksmith start
enterprise: false
docker: true
---

# `blacksmith start`

This command starts the `gateway` and / or `scheduler` service(s) of your application.
Before starting, the CLI builds the application to validate it.

**Example:**
```bash
$ blacksmith start --bind 3000:9090 --bind 3001:9091

```

**Related ressources:**
- Getting started >
  [Docker environment](/blacksmith/introduction/start/docker)
- Getting started >
  [Running an instance](/blacksmith/introduction/start/run)

## Optional flags

- `--service [name]`: Name of the service to start. It must be one of `gateway`
  or `scheduler`. If the flag is not passed, both services will start in the
  same process. In production, it is highly recommended to run the gateway and
  scheduler services on separate *applications* for better control over security
  and scalability.

  **Example:**
  ```bash
  $ blacksmith start --service gateway

  ```

- `--bind [port:port]`: Publish a container's port(s) to the host. Without this
  flag, it is not possible for a host (like a workstation for local development)
  to access the ports exposed by the Docker container. In the following example,
  the port `9090` of the container exposed by the service `gateway` will be binded
  to the port `3000` of the host.

  **Example:**
  ```bash
  $ blacksmith start --service gateway --bind 3000:9090

  ```

- `--no-cache`: Do not use the Docker cache when building the application before
  starting the service(s).

  **Example:**
  ```bash
  $ blacksmith start --no-cache

  ```
