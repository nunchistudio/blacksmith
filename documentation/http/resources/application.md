---
title: Application resources
enterprise: true
---

# Application resources

The HTTP API exposes endpoints to retrieve general information about a Blacksmith
application.

## Retrieve application's options

This endpoint exposes the options for the services and adapters passed when creating
the Blacksmith application. It does not include information about sources and
destinations.

- **Method:** `GET`
- **Path:** `/admin/api/options`

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/options'

  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "supervisor": {
        "from": "consul",
        "node": {
          "name": "node-1",
          "address": "https://consul-1.example.com",
          "tags": ["blacksmith"],
          "meta": {
            "go_version": "1.16.0",
            "blacksmith_version": "0.15.1"
          }
        }
      },
      "wanderer": {
        "from": "postgres"
      },
      "store": {
        "from": "postgres"
      },
      "pubsub": {
        "from": "nats",
        "topic": "blacksmith",
        "subscription": "blacksmith"
      },
      "gateway": {
        "address": ":9090",
        "admin": {
          "enabled": false,
          "dashboard": false
        }
      },
      "scheduler": {
        "address": ":9091",
        "admin": {
          "enabled": true,
          "dashboard": true
        }
      },
      "blacksmith": {
        "version": "0.15.1"
      },
      "go": {
        "version": "1.16",
        "environment": "darwin/amd64"
      }
    }
  }

  ```
