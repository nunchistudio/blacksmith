---
title: Schema versioning
enterprise: true
---

# Schema versioning

An API acts like a contract, between a producer and consumers. It can't be changed
without considerable cooperation and effort, on both sides. If the producer changes
the contract, the consumers relying on it need to make some changes to maintain
compatibility.

This is also true when working with Blacksmith. If a source or a destination
changes its data schema, some changes need to be made in other places.

> Schema versioning is not a mandatory practice to use Blacksmith in production.
  It is a best practice recommendation that should be applied when possible and
  practical.

## How it works

To avoid breaking changes in your data pipeline, the best approach is to follow
schema versioning, by using the built-in versioning feature. It allows to have a
collection of versions for each `source` and `destination` adapter. You can mark
a version as the default one if none was provided by the consumer, and also mark
a version as deprecated when needed.

In the `Options`, a source and destination can set a collection of supported versions.
The value of each version is its deprecation date. It must be set to an empty
`time.Time` when the version is still maintained.

When you want to sunset a version number, simply remove its key from `Versions`.

> The version numbering is free. You can use semantic versioning, calendar
  versioning, etc. We higly recommend to use the same numbering convention across
  every sources and destinations for consistency.

## Version a source

### Configuration

In the following example, two versions are supported:
- `2020-10-01` is the default version. When the `Version` key is not set by the
  trigger, this version number will automatically be applied.
- `2020-06-01` will be deprecated on 2020-12-01. It is still usable and supported,
  but a warning will be prompted so engineering teams are aware about an outdated
  consumer.

```go
func New() source.Source {

  return &Source{
    options: &source.Options{

      // ...

      DefaultVersion: "2020-10-01",
      Versions: map[string]time.Time{
        "2020-10-01": time.Time{},
        "2020-06-01": time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
      },
    },
  }
}
```

### Usage

Here, if the `Version` key is not in the returned `*source.Payload`, the
`DefaultVersion` of the source will be applied. Given the previous code example,
it will be the version number `2020-10-01`.

```go
func (t MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

  // ...

  switch payload.Version {
  case "2020-10-01":
    // ...

  case "2020-06-01":
    // ...

  }

  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

## Version a destination

### Configuration

In the following example, two versions are supported:
- `2020-10-01` is the default version. When the `Version` key is not set by the
  calling flows, this version number will automatically be applied.
- `2020-06-01` will be deprecated on 2020-12-01. It is still usable and supported,
  but a warning will be prompted so engineering teams are aware about an outdated
  flow.

```go
func New() destination.Destination {

  return &Destination{
    options: &destination.Options{

      // ...

      DefaultVersion: "2020-10-01",
      Versions: map[string]time.Time{
        "2020-10-01": time.Time{},
        "2020-06-01": time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
      },
    },
  }
}
```

### Usage

Here, the `Version` key of each job present in the queue of events will be set to
the `DefaultVersion` of the destination if none was provided by the flow. Given
the previous code example, it will be the version number `2020-10-01`.

```go
func (a MyAction) Load(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  for _, event := range queue.Events {
    for _, job := range event.Jobs {
    
      switch job.Version {
      case "2020-10-01":
        // ...

      case "2020-06-01":
        // ...

      }
    }
  }
}
```
