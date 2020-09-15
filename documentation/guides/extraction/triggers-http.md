---
title: HTTP requests
enterprise: false
---

# Triggers: HTTP requests

Triggers of mode `http` allow data extraction from a source when an API route is
requested. The most interesting use case is for capturing webhooks from third-party
services. This way, whenever a condition is met in one of your applications, it
can automatically make a HTTP request to one of the triggers registered in `http`
mode.

## Configuration

To make a source's trigger handles HTTP request, you first need to specifiy the
mode to `http` in the trigger's options.

The following route will trigger the `Extract` function every time a HTTP request
comes in `/identify` using a `POST` method:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeHTTP,
    UsingHTTP: &source.Route{
      Methods: []string{"POST"},
      Path:    "/identify",
    },
  }
}
```

## Usage

If the trigger mode is `http`, it must respect the interface
[`source.TriggerHTTP`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerHTTP).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *http.Request) (*source.Payload, error)
```

The gateway ensures the content type returned is `application/json`. It can also
includes information inside the response body such as the jobs created by the
flows called.

## Example

```go
/*
Extract is the function being executed when the HTTP route is
called. The function gives access to the original HTTP request.
*/
func (t MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

  // Create an empty payload, catch unwanted fields, and
  // unmarshal it. Return an error if any occured.
  var payload MyTrigger
  decoder := json.NewDecoder(req.Body)
  decoder.DisallowUnknownFields()
  err := decoder.Decode(&payload)
  if err != nil {
    return nil, err
  }

  // Try to marshal the context from the request payload.
  ctx, err := json.Marshal(&payload.Context)
  if err != nil {
    return nil, err
  }

  // Try to marshal the data from the request payload.
  data, err := json.Marshal(&payload.Data)
  if err != nil {
    return nil, err
  }

  // Return the context, data, and a collection of flows to execute
  // against destinations with transformed data. More details about
  // data flowing in the "Flows" guide.
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```
