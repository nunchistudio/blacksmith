---
title: Blacksmith configuration
enterprise: false
---

# Blacksmith configuration

## Common configuration

The Blacksmith configuration reference includes details, notes, and environment
variables for configuring an application.

The first step is to configure the `gateway` and `scheduler` services. Then, the
`store` adapter since this is the only one mandatory.

The `pubsub` adapter must be configured for realtime data loading and extracting
data from Pub / Sub messages.

[As described in the introduction](/blacksmith/introduction/start/create),
Blacksmith options is of type
[`*blacksmith.Options`](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc#Options).
The options must be passed in the application `Init` function, as follow:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

  }

  return options
}

```

## Enterprise Edition configuration

For leveraging Blacksmith Enterprise Edition you must pass your subscription
details. The first way is to set the environment variables `BLACKSMITH_LICENSE_KEY`
and `BLACKSMITH_LICENSE_TOKEN`. The second one is to pass those values into the
`License` property of the Blacksmith options, like this:

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    License: &blacksmith.License{
      Key:   "key",
      Token: "token",
    },
  }

  return options
}

```

> We strongly advise to pass these values using the environment variables
  `BLACKSMITH_LICENSE_KEY` and `BLACKSMITH_LICENSE_TOKEN` to avoid exposing
  credentials in your code.

The CLI will verify your credentials when running some commands. It does so by
making a request to the Nunchi Checkpoint API. If the subscription returned is
not valid, Blacksmith will stop the desired task.

Also, every 4 hours the `gateway` and `scheduler` services will make the same
request as detailed before. If the subscription has expired or if it is not valid
anymore for some reason, Blacksmith will prompt a warning until the subscription
is valid again, for up to 72 hours. After that, the running service(s) will
gracefully be interrupted.

You can update your subscription, billing information, and payment details by
running:
```bash
$ blacksmith login

```

It will prompt a link to open a new session within the Nunchi Customer Portal,
powered by [Stripe](https://stripe.com/) for simplified billing:

![Nunchi Customer Portal](/images/blacksmith/portal.png)

To avoid any network error, make sure you can make a request against the following
URL from your different environments:
```bash
$ curl --request GET --url 'https://nunchi.studio/api/checkpoint'

```
