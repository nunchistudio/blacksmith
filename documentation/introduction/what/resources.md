---
title: Resources
enterprise: false
---

# Resources

Blacksmith is composed of several components, each acting within its own scope
for better separation of concerns.

## Go API

The **Blacksmith Go API** is a public-facing collection of packages. It is written
in [Go](https://golang.org/), "*the language of the cloud*". It allows you to
define your data solution as-code for complete control and versioning. As such,
it is used for configuring a Blacksmith application, or creating in-house sources
and destinations.

**Related resources:**
- [Blacksmith repository on GitHub](https://github.com/nunchistudio/blacksmith)
- [Go reference on Go Developer Portal](https://pkg.go.dev/github.com/nunchistudio/blacksmith)

## Go Modules

The **Go Modules for Blacksmith** is a collection of open-source third-party
integrations allowing to quickly integrate your Blacksmith application with the
tools you already use and love. Each module leverages one or several packages
from the Go API, letting you use a source or a destination within a few lines of
code.

**Related resources:**
- [Integrations catalog on Nunchi website](/blacksmith/modules)
- [Integrations repository on GitHub](https://github.com/nunchistudio/blacksmith-integrations)

## Command line

The **Blacksmith CLI** allows to build, run, and manage your application which is
leveraging the Go API.

**Related resources:**
- [Download the CLI](/blacksmith/downloads)
- [CLI reference](/blacksmith/cli)

## Docker images

**Blacksmith images for Docker** ensures environment parity and make deployments
a breeze. When running a command that needs to build and / or run an application,
the CLI will communicate with the Docker instance installed on the machine, and
will run itself inside a container created based on a `Dockerfile`.

By forwarding most of the command of the local Blacksmith CLI to a container, we
make sure your application can run on different machines, regardless the environment,
as long as a Docker daemon is running.

The Docker images contain all the non-public logic for running a Blacksmith
application. Therefore, running Blacksmith outside one of these images is not
officially supported.

**Related resources:**
- [Docker images on GitHub](https://github.com/nunchistudio/blacksmith-docker)
- [Docker images on Docker Hub](https://hub.docker.com/r/nunchistudio)

## UI kit

The **Blacksmith UI kit** is a collection of open-source, reusable, front-end
components. It lets you embed any kind of information from your Blacksmith
application in a custom dashboard within a few lines of code. This is particularly
useful for front-end developers when creating custom dashboards.

**Related resources:**
- [Blacksmith UI kit on GitHub](https://github.com/nunchistudio/blacksmith-ui)
- [Storybook of Blacksmith UI](/storybook/blacksmith-eui)

## Dashboard

The **Blacksmith Dashboard** is the dashboard built-in within any application using
the Enterprise Edition. It leverages the Blacksmith UI kit to simplify custom work
on top of it.

![Blacksmith Dashboard](/images/blacksmith/dashboard.002.png)

**Related resources:**
- ["Template" repository on GitHub](https://github.com/nunchistudio/blacksmith-dashboard)
