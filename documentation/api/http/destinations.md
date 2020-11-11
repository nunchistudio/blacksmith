---
title: Destinations resources
enterprise: true
---

# Destinations resources

The HTTP API exposes endpoints to retrieve information about destinations of a
Blacksmith application.

## Retrieve all destinations

This endpoint exposes all the destinations registered in an application, including
the options for each one.

- **Method:** `GET`
- **Path:** `/admin/destinations`

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/destinations'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 7
    },
    "data": [
      {
        "name": "my-destination",
        "options": {
          "versions": {
            "2020-10-27": "0001-01-01T00:00:00Z"
          },
          "default_version": "2020-10-27",
          "schedule": {
            "realtime": true,
            "interval": "@every 1h",
            "max_retries": 50
          }
        }
      },
      
      [...]
      
    ]
  }
  ```

## Retrieve a specific destination

This endpoint exposes details about a single destination registered in an application,
including its options and actions.

- **Method:** `GET`
- **Path:** `/admin/destinations/:destination_name`
- **Route params:**
  - `destination_name`: Name of the destination to retrieve.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/destinations/my-destination'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "name": "my-destination",
      "options": {
        "versions": {
          "2020-10-27": "0001-01-01T00:00:00Z"
        },
        "default_version": "2020-10-27",
        "schedule": {
          "realtime": true,
          "interval": "@every 1h",
          "max_retries": 50
        }
      },
      "actions": [
        {
          "name": "action-a",
          "schedule": null
        },

        [...]
        
      ]
    }
  }
  ```
