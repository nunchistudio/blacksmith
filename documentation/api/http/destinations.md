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
- **Path:** `/admin/api/destinations`

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/destinations'

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
including its options and some details about its actions.

- **Method:** `GET`
- **Path:** `/admin/api/destinations/:destination_name`
- **Route params:**
  - `destination_name`: Name of the destination to retrieve.

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/destinations/my-destination'

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

## Retrieve a specific action

This endpoint exposes every details about an action, including its semaphore status
given by the `supervisor` adapter (if enabled).

- **Method:** `GET`
- **Path:** `/admin/api/destinations/:destination_name/actions/:action_name`
- **Route params:**
  - `destination_name`: Name of the destination to retrieve.
  - `action_name`: Name of the action to retrieve.

- **Example request:**
  ```bash
  $ curl --request GET --url 'http://localhost:9091/admin/api/destinations/my-destination/actions/action-a'

  ```
- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "name": "action-a",
      "schedule": null,
      "semaphore": {
        "key": "actions/demo-destination-one/action-a",
        "is_applicable": true,
        "is_acquired": true,
        "acquirer_name": "blacksmith-scheduler",
        "acquirer_address": ":9091",
        "session_id": "1p1RzXlka08MaE2ht3jRWW36isZ"
      }
    }
  }

  ```
