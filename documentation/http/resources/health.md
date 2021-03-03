---
title: Health check
enterprise: true
---

# Health check

Blacksmith exposes an endpoint both on the `gateway` and `scheduler` services so
they can be registered in a service mesh. This allows to be aware of the health
of each of them.

For example, the [`supervisor` adapter for Consul](/blacksmith/options/supervisor/consul)
automatically registers these services in the Consul catalog, using this endpoint
as the health check.

This endpoint is not related to the admin API and is exposed even if the admin API
is disabled.

- **Method:** `GET`
- **Path:** `/_health`

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/_health'

  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful"
  }

  ```
