---
title: Configuring an application
enterprise: false
---

# Configuring an application

As we learned in the previous guide, an application must be configured in the
`Init` function by returning options, of type
[`*blacksmith.Options`](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc#Options).

When generating an application, all the options are already set to work in a
development environment. You can skip this section if you are working on your
local machine with Docker.

Otherwise, please refer to [the configuration reference](/blacksmith/options) to
properly configure your application for a non-local environment, depending on your
needs. This should only takes a few minutes.
