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
        "from": "consul"
      },
      "wanderer": {
        "from": "postgres"
      },
      "store": {
        "from": "postgres",
        "purge": [
          {
            "where": {
              "jobs": {
                "transitions": {
                  "jobs.status_in": [
                    "succeeded"
                  ],
                  "jobs.status_notin": [
                    "acknowledged",
                    "awaiting",
                    "executing",
                    "failed",
                    "discarded",
                    "unknown"
                  ]
                }
              },
              "offset": 0,
              "limit": 0
            },
            "interval": "@weekly"
          }
        ]
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
        "version": "0.17.0"
      },
      "go": {
        "version": "1.16",
        "environment": "darwin/amd64"
      }
    }
  }

  ```
