---
title: Sources resources
enterprise: true
---

# Sources resources

The HTTP API exposes endpoints to retrieve information about sources of a Blacksmith
application.

## Retrieve all sources

This endpoint exposes all the sources registered in an application, including
the options for each one.

- **Method:** `GET`
- **Path:** `/admin/api/sources`

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/sources'

  ```
- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 5
    },
    "data": [
      {
        "name": "my-source",
        "options": {
          "versions": {
            "2020-10-27": "0001-01-01T00:00:00Z"
          },
          "default_version": "2020-10-27",
          "cron": {
            "interval": "@every 1h"
          }
        }
      },
      
      [...]

    ]
  }

  ```

## Retrieve a specific source

This endpoint exposes details about a single source registered in an application,
including its options and triggers.

- **Method:** `GET`
- **Path:** `/admin/api/sources/:source_name`
- **Route params:**
  - `source_name`: Name of the source to retrieve.

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/sources/my-source'

  ```
- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "name": "my-source",
      "options": {
        "versions": {
          "2020-10-27": "0001-01-01T00:00:00Z"
        },
        "default_version": "2020-10-27",
        "cron": {
          "interval": "@every 1h"
        }
      },
      "triggers": [
        {
          "name": "trigger-a",
          "mode": {
            "mode": "cron",
            "cron": {
              "interval": "@every 40s"
            }
          }
        },
        {
          "name": "trigger-b",
          "mode": {
            "mode": "http",
            "http": {
              "methods": [
                "POST"
              ],
              "path": "/endpoint",
              "show_meta": true,
              "show_data": true
            }
          }
        },
        
        [...]
        
      ]
    }
  }

  ```
