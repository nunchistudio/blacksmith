---
title: Purge policies
enterprise: false
---

# Purge policies

Every day your application can handle millions or billions of incoming events and
outgoing jobs. As you might be aware of, everything is saved in the `store` adapter.
The more data you store, the more your application can become slower to ingest and
process said data. This impacts every components of the application, from the
events' ingestion by the `gateway` to the jobs' responsiveness handled by the
`scheduler`.

This is why it is very important for a data engineering platform to purge unused
data. This frees up space by deleting obsolete entries that are not required by
your solution.

Blacksmith offers advanced *purge policies* allowing data retention at a high or
granular level.

## Usage with Go API

To implement these, you simply need to add the desired policies to the `store`
adapter's options. It it of type
[`*store.PurgePolicy`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/adapter/store?tab=doc#PurgePolicy)
and can accept very granular inclusion / exclusion conditions.

The following example contains two purge policies:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/store"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Store: &store.Options{
      From: "postgres",
      PurgePolicies: []*store.PurgePolicy{
        {
          Interval: "@weekly",
          WhereEvents: &store.WhereEvents{
            AndWhereJobs: &store.WhereJobs{
              AndWhereTransitions: &store.WhereTransitions{
                StatusIn: []string{
                  store.StatusSucceeded,
                },
                StatusNotIn: []string{
                  store.StatusAcknowledged,
                  store.StatusAwaiting,
                  store.StatusExecuting,
                  store.StatusFailed,
                  store.StatusDiscarded,
                  store.StatusUnknown,
                },
              },
            },
          },
        },

        {
          Interval: "@daily",
          WhereEvents: &store.WhereEvents{
            SourcesIn: []string{
              "my-source-one",
              "my-source-two",
            },
            AndWhereJobs: &store.WhereJobs{
              AndWhereTransitions: &store.WhereTransitions{
                StatusIn: []string{
                  store.StatusSucceeded,
                },
                StatusNotIn: []string{
                  store.StatusAcknowledged,
                  store.StatusAwaiting,
                  store.StatusExecuting,
                  store.StatusFailed,
                  store.StatusDiscarded,
                  store.StatusUnknown,
                },
              },
            },
          },
        },
      },
    },
  }

  return options
}

```

The first purge policy is the only one present when generating a new application.
It runs weekly (at midnight between Saturday and Sunday). It purges every entries
related to jobs marked as `succeeded` but exclude the ones related to jobs marked
as `acknowledged`, `awaiting`, `executing`, `failed`, `discarded`, or `unknown`. 
Because a job can be related to both an event *and* a parent job, it is critical
to strictly exclude the entries related to jobs' status that we wish to retain.
In this case, we only want to purge all entries related to *only* successful jobs.

As an example, an event can create two jobs. One is marked as `succeeded` and the
other one as `failed`. If we simply set `succeeded` in `StatusIn`, the event will
be purged from the store because the event is in fact related to a successful job.
However, if we add `failed` in `StatusNotIn`, the event will not be purged because
it is related to a job status explicitly listed in the exclusion conditions.

The second policy acts almost like the first one. It runs daily (at midnight) but
is only applied for the sources `my-source-one` and `my-source-two`.

## Usage with HTTP API

The Go API is useful for defining policies directly inside the application for
regular intervals and fixed policies. You might sometimes need to manually purge
data from the `store` adapter.

This can be done with the HTTP API. It can accept an incoming request and then
purge the data based on the request's params.
[Read how it works.](/blacksmith/api/http/store)
