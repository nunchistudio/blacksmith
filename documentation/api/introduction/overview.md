---
title: Overview of the Blacksmith API
enterprise: true
---

# Overview of the Blacksmith API

The Blacksmith API is organized around REST. The API has predictable resource-oriented
URLs, returns JSON-encoded responses, and uses standard HTTP response codes and
verbs.

The API exposes information about your data platform, such as configuration,
realtime and historical events, jobs' status, migrations, etc. Therefore, it
allows organizations to embed data in third-party services and build custom
dashboards on top of it.

## Status codes

Blacksmith uses conventional HTTP response codes to indicate the success or failure
of an API request. Codes in the `2xx` range indicate success. Codes in the `4xx`
range indicate an error that failed given the information provided. Codes in the
`5xx` range indicate an error with the Blacksmith framework.

> If you experience a `5xx` error and the error is not correctly handled, please
  open an issue on GitHub describing the expected behavior and the steps to
  reproduce the error encountered.

## Metadata

All HTTP responses follow the same design including the status code, a message,
and some data:
```json
{
  "statusCode": 200,
  "message": "Successful",
  "data": {
    
    [...]
    
  }
}
```

When retrieving lists, a `meta` object is also included to add metadata about the
request and / or the response, such as the total count of objects found, the query
applied, and pagination details.

For example, the following request allows to retrieve every events in the `store`
adapter coming from the source `my-source` and where produced jobs are either
`failed` or `discarded`, with a limit of `50` events per page:
```bash
curl -G 'http://localhost:9091/admin/store/events' \
  -d events.sources_in=my-source \
  -d jobs.status_in=failed \
  -d jobs.status_in=discarded \
  -d limit=50
```

This will produce the following response:
```json
{
  "statusCode": 200,
  "message": "Successful",
  "meta": {
    "count":11,
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
      "jobs": {
        "transitions": {
          "jobs.status_in": [
            "failed",
            "discarded"
          ]
        }
      },
      "offset": 0,
      "limit": 50
    }
  },
  "data": [

    [...]

  ]
}
```
