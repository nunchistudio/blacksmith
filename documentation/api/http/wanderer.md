---
title: Wanderer resources
enterprise: true
---

# Wanderer resources

The HTTP API exposes endpoints to retrieve migrations, and migrations' status (also
known as migrations' transitions) of a Blacksmith application.

When retrieving a collection of migrations, the request can have query params for
searching, filtering, grouping, and paginating objects.

List of available query params:

- **Name:** `migrations.scopes_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  scope present in the slice.

- **Name:** `migrations.scopes_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the scope present in the slice.

- **Name:** `migrations.versioned_before`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to a
  migration versioned before this instant.

- **Name:** `migrations.versioned_after`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to a
  migration versioned after this instant.

- **Name:** `offset`

  **Type:** `uint`

  **Description:** Specifies the number of entries to skip before starting to
  return entries from the query.

  **Default value:** `0`

- **Name:** `limit`

  **Type:** `uint`

  **Description:** Specifies the number of entries to return after the offset clause
  has been processed.

  **Default value:** `100`

## Retrieve all migrations

This endpoint exposes all the migrations registered in the wanderer given the
filters passed as query parameters.

- **Method:** `GET`
- **Path:** `/admin/wanderer/migrations`
- **Query params:** As listed at the top of this document.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/wanderer/migrations' \
    -d migrations.status_in=acknowledged \
    -d migrations.scopes_in=destination:my-destination
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 1,
      "pagination": {
        "current": 1,
        "previous": null,
        "next": null,
        "first": 1,
        "last": 1
      },
      "where": {
        "scope_in": [
          "destination:my-destination"
        ],
        "transitions": {
          "status_in": [
            "acknowledged"
          ]
        },
        "offset": 0,
        "limit": 100
      }
    },
    "data": [
      {
        "id": "1jbTWtayiztjlceyq9OiZiudL84",
        "version": "2020-10-30T15:23:01Z",
        "scope": "destination:my-destination",
        "name": "init",
        "transitions": [
          {
            "id": "1jbTX0m6lUDs2ns4TUyathJaDgm",
            "state_before": null,
            "state_after": "acknowledged",
            "error": null,
            "created_at": "2020-10-30T15:30:26.831789Z",
            "migration_id": "1jbTWtayiztjlceyq9OiZiudL84"
          }
        ],
        "created_at": "2020-10-30T15:30:26.804643Z"
      },

      [...]

    ]
  }
  ```

## Retrieve a specific migration

This endpoint exposes details about a single migration registered in the wanderer,
including its current status, which is its latest transition.

- **Method:** `GET`
- **Path:** `/admin/wanderer/migrations/:migration_id`
- **Route params:**
  - `migration_id`: ID of the migration to retrieve.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/wanderer/migrations/1jbTWtayiztjlceyq9OiZiudL84'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "id": "1jbTWtayiztjlceyq9OiZiudL84",
      "version": "2020-10-30T15:23:01Z",
      "scope": "destination:my-destination",
      "name": "init",
      "transitions": [
        {
          "id": "1jbTX0m6lUDs2ns4TUyathJaDgm",
          "state_before": null,
          "state_after": "acknowledged",
          "error": null,
          "created_at": "2020-10-30T15:30:26.831789Z",
          "migration_id": "1jbTWtayiztjlceyq9OiZiudL84"
        }
      ],
      "created_at": "2020-10-30T15:30:26.804643Z"
    }
  }
  ```

## Retrieve a migration's transitions

This endpoint exposes all the transitions registered in the wanderer for a given
migration.

- **Method:** `GET`
- **Path:** `/admin/wanderer/migrations/:migration_id/transitions`
- **Route params:**
  - `migration_id`: ID of the migration to retrieve.

- **Query params:**
  - **Name:** `offset`

    **Type:** `uint`

    **Description:** Specifies the number of entries to skip before starting to
    return entries from the query.

    **Default value:** `0`

  - **Name:** `limit`

    **Type:** `uint`

    **Description:** Specifies the number of entries to return after the offset clause
    has been processed.

    **Default value:** `100`

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/wanderer/migrations/1jbTWtayiztjlceyq9OiZiudL84/transitions' \
    -d limit=50
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 2,
      "pagination": {
        "current": 1,
        "previous": null,
        "next": null,
        "first": 1,
        "last": 1
      },
      "where": {
        "transitions": {
          "migrations.id": "1jbTWtayiztjlceyq9OiZiudL84"
        },
        "offset": 0,
        "limit": 50
      }
    },
    "data": [
      {
        "id": "1jbTX0m6lUDs2ns4TUyathJaDgm",
        "state_before": null,
        "state_after": "acknowledged",
        "error": null,
        "created_at": "2020-10-30T15:30:26.831789Z",
        "migration_id": "1jbTWtayiztjlceyq9OiZiudL84"
      },
      
      [...]
      
    ]
  }
  ```
