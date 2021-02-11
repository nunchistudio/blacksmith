---
title: Store resources
enterprise: true
---

# Store resources

The HTTP API exposes endpoints to retrieve events, jobs, and jobs' status (also
known as jobs' transitions) of a Blacksmith application.

When retrieving a collection of events or jobs, the request can have query params
for searching, filtering, grouping, and paginating objects. This is a very powerful
feature allowing you to have complete view of your data status at the dimension
you need.

List of available query params:

- **Name:** `events.sources_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  source name present in the slice.

- **Name:** `events.sources_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the source name present in the slice.

- **Name:** `events.triggers_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  source's trigger name present in the slice.

- **Name:** `events.triggers_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the source's trigger name present in the slice.

- **Name:** `events.versions_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  source's version present in the slice.

- **Name:** `events.versions_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the source's version present in the slice.

- **Name:** `events.created_before`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to an
  event created before this instant.

- **Name:** `events.created_after`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to an
  event created after this instant.

- **Name**: `jobs.destinations_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  destination name present in the slice.

- **Name**: `jobs.destinations_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the destination name present in the slice.

- **Name**: `jobs.actions_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  destination's action name present in the slice.

- **Name**: `jobs.actions_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the destination's action name present in the slice.

- **Name**: `jobs.versions_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  destination's version present in the slice.

- **Name**: `jobs.versions_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the destination's version present in the slice.

- **Name**: `jobs.created_before`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to a
  job created before this instant.

- **Name**: `jobs.created_after`

  **Type:** `time.Time`

  **Description:** Makes sure the entries returned by the query are related to a
  job created after this instant.

- **Name**: `jobs.status_in`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query have any of the
  status present in the slice.

- **Name**: `jobs.status_notin`

  **Type:** `[]string`

  **Description:** Makes sure the entries returned by the query do not have any
  of the status present in the slice.

- **Name**: `jobs.min_attempts`

  **Type:** `uint`

  **Description:** Makes sure the entries returned by the query have equal to or
  greater than this number of attempts.

- **Name**: `jobs.max_attempts`

  **Type:** `uint`

  **Description:** Makes sure the entries returned by the query have equal to or
  lesser than this number of attempts.

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

## Retrieve all events

This endpoint exposes all the events registered in the store given the filters passed
as query parameters. Events' details such as `context` and `data` keys are not
included in events and jobs information. Jobs only include their current state,
which is their lastest transition.

- **Method:** `GET`
- **Path:** `/admin/store/events`
- **Query params:** As listed at the top of this document.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/store/events' \
    -d events.sources_in=my-source \
    -d offset=50 \
    -d limit=50
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 73,
      "pagination": {
        "current": 2,
        "previous": 1,
        "next": null,
        "first": 1,
        "last": 2
      },
      "where": {
        "events.sources_in": [
          "my-source"
        ],
        "offset": 50,
        "limit": 50
      }
    },
    "data": [
      {
        "id": "1jbDehaSRN1whZAYNqRawEUrI7g",
        "source": "my-source",
        "trigger": "trigger-a",
        "version": "2020-10-27",
        "jobs": [
          {
            "id": "1jbDei2Hfzb4poa4JjouoMk87ZA",
            "destination": "my-destination",
            "action": "action-a",
            "version": "2020-10-27",
            "transitions": [
              {
                "id": "1jbDelDtJbfiMXphT2srUe7H8QK",
                "attempt": 1,
                "state_before": "executing",
                "state_after": "discarded",
                "error": {
                  "statusCode": 401,
                  "message": "Not authorized",
                  "validations": [
                    {
                      "message": "Email address not authorized",
                      "path": [
                        "request",
                        "payload",
                        "data",
                        "email"
                      ]
                    }
                  ]
                },
                "created_at": "2020-10-30T13:19:54.032384Z",
                "event_id": "1jbDehaSRN1whZAYNqRawEUrI7g",
                "job_id": "1jbDei2Hfzb4poa4JjouoMk87ZA"
              }
            ],
            "created_at": "2020-10-30T13:19:54.005406Z",
            "event_id": "1jbDehaSRN1whZAYNqRawEUrI7g",
            "parent_job_id": ""
          }
        ],
        "received_at": "2020-10-30T13:19:54.001208Z",
        "ingested_at": "2020-10-30T13:19:54.006898Z"
      },
      
      [...]

    ]
  }
  ```

## Retrieve a specific event

This endpoint exposes details about a single event registered in the store,
including their `context` and `data` keys. Jobs only include their current state,
which is their lastest transition.

- **Method:** `GET`
- **Path:** `/admin/store/events/:event_id`
- **Route params:**
  - `event_id`: ID of the event to retrieve.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/store/events/1jbDyotE3aB7qYNOaSQRlLa3sRK'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "id": "1jbDyotE3aB7qYNOaSQRlLa3sRK",
      "source": "my-source",
      "trigger": "trigger-a",
      "version": "2020-10-27",
      "context": { [...] },
      "data": { [...] },
      "jobs": [
        {
          "id": "1jbDynjIuCBqcDAR5PkhwVjvzZ2",
          "destination": "my-destination",
          "action": "action-a",
          "version": "2020-10-27",
          "context": { [...] },
          "data": { [...] },
          "transitions": [
            {
              "id": "1jbDytaDvFFBUKr3uDg1C5M6nL0",
              "attempt": 1,
              "state_before": "executing",
              "state_after": "discarded",
              "error": {
                "statusCode": 401,
                "message": "Not authorized",
                "validations": [
                  {
                    "message": "Email address not authorized",
                    "path": [
                      "request",
                      "payload",
                      "data",
                      "email"
                    ]
                  }
                ]
              },
              "created_at": "2020-10-30T13:22:34.033247Z",
              "event_id": "1jbDyotE3aB7qYNOaSQRlLa3sRK",
              "job_id": "1jbDynjIuCBqcDAR5PkhwVjvzZ2"
            }
          ],
          "created_at": "2020-10-30T13:22:34.004662Z",
          "event_id": "1jbDyotE3aB7qYNOaSQRlLa3sRK",
          "parent_job_id": ""
        }
      ],
      "received_at": "2020-10-30T13:22:34.001514Z",
      "ingested_at": "2020-10-30T13:22:34.006282Z"
    }
  }
  ```

## Retrieve all jobs

This endpoint exposes all the jobs registered in the store given the filters passed
as query parameters.

- **Method:** `GET`
- **Path:** `/admin/store/jobs`
- **Query params:** As listed at the top of this document.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/store/jobs' \
    -d events.sources_in=my-source
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 97,
      "pagination": {
        "current": 1,
        "previous": null,
        "next": null,
        "first": 1,
        "last": 1
      },
      "where": {
        "events.sources_in": [
          "my-source"
        ],
        "offset": 0,
        "limit": 100
      }
    },
    "data": [
      {
        "id": "1jbF63OeSFBbuez4351lfZ4f2jL",
        "destination": "my-destination",
        "action": "2020-10-27",
        "version": "action-a",
        "context": { [...] },
        "data": { [...] },
        "transitions": [
          {
            "id": "1jbF65rimr9iEVCGbDHTgOgjM7x",
            "attempt": 1,
            "state_before": "executing",
            "state_after": "discarded",
            "error": {
              "statusCode": 401,
              "message": "Not authorized",
              "validations": [
                {
                  "message": "Email address not authorized",
                  "path": [
                    "request",
                    "payload",
                    "data",
                    "email"
                  ]
                }
              ]
            },
            "created_at": "2020-10-30T13:31:45.024292Z",
            "event_id": "1jbF681tRCyaleY0P5lcKpgejkg",
            "job_id": "1jbF63OeSFBbuez4351lfZ4f2jL"
          }
        ],
        "created_at": "2020-10-30T13:31:45.002117Z",
        "event_id": "1jbF681tRCyaleY0P5lcKpgejkg"
      },

      [...]

    ]
  }
  ```

## Retrieve a specific job

This endpoint exposes details about a single migration registered in the wanderer,
including its current status, which is its latest transition.

- **Method:** `GET`
- **Path:** `/admin/store/jobs/:job_id`
- **Route params:**
  - `job_id`: ID of the job to retrieve.

- **Example request:**
  ```bash
  curl -G 'http://localhost:9091/admin/store/jobs/1jbF63OeSFBbuez4351lfZ4f2jL'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "data": {
      "id": "1jbF63OeSFBbuez4351lfZ4f2jL",
      "destination": "my-destination",
      "action": "action-a",
      "version": "2020-10-27",
      "context": { [...] },
      "data": { [...] },
      "transitions": [
        {
          "id": "1jbF65rimr9iEVCGbDHTgOgjM7x",
          "attempt": 1,
          "state_before": "executing",
          "state_after": "discarded",
          "error": {
            "statusCode": 401,
            "message": "Not authorized",
            "validations": [
              {
                "message": "Email address not authorized",
                "path": [
                  "request",
                  "payload",
                  "data",
                  "email"
                ]
              }
            ]
          },
          "created_at": "2020-10-30T13:31:45.024292Z",
          "event_id": "1jbF681tRCyaleY0P5lcKpgejkg",
          "job_id": "1jbF63OeSFBbuez4351lfZ4f2jL"
        }
      ],
      "created_at": "2020-10-30T13:31:45.002117Z",
      "event_id": "1jbF681tRCyaleY0P5lcKpgejkg"
    }
  }
  ```

## Retrieve a job's transitions

This endpoint exposes all the transitions registered in the store for a given
job.

- **Method:** `GET`
- **Path:** `/admin/store/jobs/:job_id/transitions`
- **Route params:**
  - `job_id`: ID of the job to retrieve.

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
  curl -G 'http://localhost:9091/admin/store/jobs/1jbHsR1l5r10ozAZFY4D23o2uZr/transitions' \
    -d limit=100
  ```

- **Example response**:
  ```json
  {
    "statusCode": 200,
    "message": "Successful",
    "meta": {
      "count": 140,
      "pagination": {
        "current": 1,
        "previous": null,
        "next": 2,
        "first": 1,
        "last": 2
      },
      "where": {
        "jobs": {
          "transitions": {
            "job.id": "1jbHsR1l5r10ozAZFY4D23o2uZr"
          }
        },
        "offset": 0,
        "limit": 100
      }
    },
    "data": [
      {
        "id": "1jbHsXny1aWQ0YAbiHA7nTmkRTT",
        "attempt": 1,
        "state_before": "executing",
        "state_after": "discarded",
        "error": {
          "statusCode": 401,
          "message": "Not authorized",
          "validations": [
            {
              "message": "Email address not authorized",
              "path": [
                "request",
                "payload",
                "data",
                "email"
              ]
            }
          ]
        },
        "created_at": "2020-10-30T13:54:37.035843Z",
        "event_id": "1jbHsWY4x2jQpy4rlrNdYB3LMYu",
        "job_id": "1jbHsR1l5r10ozAZFY4D23o2uZr"
      },

      [...]
      
    ]
  }
  ```

## Purge entries from store

This endpoint allows to manually purge the store from specific entries. Because
this can take some time and there is no data returned, this will asynchronously
run the task in background and inform the client the request has been accepted.

Even though the request is accepted, this does not serve as a guarantee for the
task to succeed.

- **Method:** `POST`
- **Path:** `/admin/store/purge`
- **Query params:** As listed at the top of this document. The `offset` and `limit`
  params will not be applied.

- **Example request:**
  ```bash
  $ curl --request POST --url 'http://localhost:9091/admin/store/purge' \
    -d jobs.status_in=discarded \
    -d events.received_before='2021-02-09 15:23:00'
  ```

- **Example response**:
  ```json
  {
    "statusCode": 202,
    "message": "Accepted",
    "meta": {
      "where": {
        "events.received_before": "2021-02-09T15:23:00Z",
        "jobs": {
          "transitions": {
            "jobs.status_in": ["discarded"]
          }
        }
      }
    }
  }

  ```
