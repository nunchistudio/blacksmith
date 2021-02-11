---
title: Tour of the dashboard
enterprise: true
---

# Tour of the dashboard

The Blacksmith Dashboard enables organizations to have a complete view of their
data engineering solution. It is built on top of the Blacksmith UI kit, which
leverages the REST API only available in the Enterprise Edition.

To make forks and customization as easy as possible, the dashboard is available as
an [open-source "template" repository on GitHub](https://github.com/nunchistudio/blacksmith-dashboard).

## Enabling the dashboard

When generating an application, the admin REST API and the built-in dashboard are
enabled by default and attached to the `scheduler` service.

Options for enabling the dashboard should look like this:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/service"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Scheduler: &service.Options{
      Address:  ":9091",
      Admin: &service.Admin{
        Enabled:       true,
        WithDashboard: true,
      },
    },
  }

  return options
}

```

With an application up and running, you can access the built-in dashboard at
<http://localhost:9091/admin/>.

## Main pages

### Events & Jobs

The home page gives an overview of the `store` adapter. One can retrieve events and
jobs, and may filter results to a very narrow scope.

![Blacksmith Dashboard](/images/blacksmith/dashboard.001.png)

You can then click on a specific event or job to access its details.

As an example, this job succeeded after 2 attempts, knowing the first one failed
because of a HTTP / API error:

![Blacksmith Dashboard](/images/blacksmith/dashboard.002.png)

This view also gives you details about the `context` and `data` keys of the job,
which can be specific to the destination and therefore different from the parent
event.

### Sources & Triggers

When selecting a source, you can view its options as well as its triggers.

![Blacksmith Dashboard](/images/blacksmith/dashboard.003.png)

### Destinations & Actions

When selecting a destination, you can view its options as well as its actions.

![Blacksmith Dashboard](/images/blacksmith/dashboard.004.png)
