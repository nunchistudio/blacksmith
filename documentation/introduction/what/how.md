---
title: How it works
enterprise: false
---

# How it works

Whether it is your first time doing data engineering or not, this guide is important
since it describes some specifities of how Blacksmith works.

## Layers

Blacksmith is composed of several layers, each acting within its own scope for
better separation of concerns.

### Go API

The **Blacksmith Go API** is a public-facing collection of packages. It is written
in [Go](https://golang.org/), "*the language of the cloud*". It allows you to
define your data stack as-code for complete control and versioning.

Related resources:
- [Blacksmith repository on GitHub](https://github.com/nunchistudio/blacksmith)
- [Go reference on Go Developer Portal](https://pkg.go.dev/github.com/nunchistudio/blacksmith)

### Docker environment

**Blacksmith on Docker** ensures environment parity and make deployments a breeze.
When running a command that needs to build and / or run an application, the CLI
will communicate with the Docker instance installed on the machine, and will run
itself inside a container created based on a `Dockerfile`.

By forwarding most of the command of the local Blacksmith CLI to a container, we
make sure your application can run on different machines, regardless the environment,
as long as a Docker daemon is running.

The Docker images contain all the non-public logic for running a Blacksmith
application. Therefore, running Blacksmith outside one of these images is not
officially supported.

Related resources:
- [Docker images on GitHub](https://github.com/nunchistudio/blacksmith-docker)
- [Docker images on Docker Hub](https://hub.docker.com/r/nunchistudio)

### UI kit

The **Blacksmith UI kit** is a collection of open-source, reusable, front-end
components. It allows to embed any kind of information from your Blacksmith
application in a custom dashboard within a few lines of code. This layer is
particularly useful for front-end developers when creating custom dashboards.

Related resources:
- [Blacksmith UI kit on GitHub](https://github.com/nunchistudio/blacksmith-ui)
- [Storybook of Blacksmith UI](/storybook/blacksmith-eui)

### Dashboard

The **Blacksmith Dashboard** is the dashboard built-in within any application using
the Enterprise Edition. It leverages the Blacksmith UI kit to simplify custom work
on top of it.

![Blacksmith Dashboard](/images/blacksmith/dashboard.002.png)

Related resources:
- ["Template" repository on GitHub](https://github.com/nunchistudio/blacksmith-dashboard)

## Concepts

### Sources and destinations

The main role of Blacksmith is to act as a data pipeline. This role takes care of
Extracting data from *sources*, Transforming it when necessary, and Loading it
to *destinations*. Blacksmith can handle transformations both before and after
saving the jobs' payload, allowing to keep the original events and still having
optimized data for each job.

Each source has *triggers* and each destination has *actions*.

In other words, a Blacksmith application is able to receive events from sources'
triggers, transform data of these events, and run it against the appropriate
destinations' actions.

Sources can be websites, mobile applications, third-party services, databases, etc.
Destinations can be databases, data warehouses, or third-party services for analytics
or marketing.

![Step 01](/images/blacksmith/how.001.png)

Triggers and actions inherit properties of their parent source or destination — such
as retry logic — but can override it in special cases.

Both triggers and actions have a business logic for handling enrichment, transformation,
and automating the data flow.

### The gateway and scheduler

Instead of using a lot of micro-services and over-engineer the whole process,
Blacksmith keep data pipelines as simple as possible with only two services: the
*gateway* and the *scheduler*.

The gateway is in charge of extracting the events from sources on triggers. It can
happen on:
- [HTTP requests](/blacksmith/guides/extraction/triggers-http);
- [CRON schedules](/blacksmith/guides/extraction/triggers-cron);
- [CDC notifications](/blacksmith/guides/extraction/triggers-cdc);
- [Pub / Sub messages](/blacksmith/guides/extraction/triggers-sub).

The scheduler is in charge of loading the events into actions of destinations,
creating *jobs*. There is an infinity of possibilities and workarounds for loading
data to destinations. Instead of locking users into a few limited patterns and
still not covering every needs you may have, we offer a few *starters*. It simplifies
development, enforces best practices, and still allows a complete freedom on how
data is loaded. Starters allow to smoothly load data to:
- [HTTP APIs](/blacksmith/guides/load/actions-http);
- [SQL databases](/blacksmith/guides/load/actions-sql);
- [NoSQL databases](/blacksmith/guides/load/actions-docstore);
- [Blob storages](/blacksmith/guides/load/actions-blob);
- [Pub / Sub mechanisms](/blacksmith/guides/load/actions-pub).

![Step 02](/images/blacksmith/how.002.png)

These services can run on a single machine but should be splitted for production
use. This way, you gain more flexibility, security, and scalability.

### Database-as-a-Queue

The gateway needs a way to keep track of received events. The scheduler needs a
way to keep track of jobs to execute onto destinations, and their status (also
called *transitions*). So it can be aware of successes, failures, and discards.

The best way to achieve this is to have a persistent *store*. Every events received
by the gateway are stored along the desired jobs. The scheduler knows how and when
to load the data to appropriate destinations and can keep a record of jobs'
transitions. 

![Step 03](/images/blacksmith/how.003.png)

The `store` adapter is the only one required along the services.

Available `store` adapters:
- [PostgreSQL](/blacksmith/options/store/postgres) (`postgres`)

### Enabling realtime

The store is perfect to persist events, jobs, and status. But we often need to
send data from sources to destinations in realtime.

We accomplish this by adding a pub / sub mechanism between the gateway and scheduler.
Once the event is received and stored, it is then automatically *published* by the
gateway to the scheduler that *subscribed* to the events.

![Step 04](/images/blacksmith/how.004.png)

The `pubsub` adapter is optional. When no adapter is provided or the action does not
handle realtime events, the scheduler will send data schedule retries given (order
matters):
- the destination's action schedule; or
- the destination's schedule; or
- the default schedule.

Once configured, the `pubsub` adapter allows realtime message extraction from
queues / topics / subscriptions.

Available `pubsub` adapters:
- [AWS SNS / SQS](/blacksmith/options/pubsub/aws) (`aws/snssqs`)
- [Azure Service Bus](/blacksmith/options/pubsub/azure) (`azure/servicebus`)
- [Google Pub / Sub](/blacksmith/options/pubsub/google) (`google/pubsub`)
- [Kafka](/blacksmith/options/pubsub/kafka) (`kafka`)
- [NATS](/blacksmith/options/pubsub/nats) (`nats`)
- [RabbitMQ](/blacksmith/options/pubsub/rabbitmq) (`rabbitmq`)

### Versioning migrations

Data pipelines often — to not say always — need to communicate with databases.
Managing and versioning database migrations is a difficult task to achieve,
especially across distributed teams in an organization. Also, the data schemas
of the data pipeline and the one of the databases need to be synchronized to
avoid schema incompatibilities as much as possible.

To resolve this challenge we add a *wanderer*. The wanderer does not handle the
migrations logic but keeps track of migration runs across every adapters, sources,
triggers, destinations, and actions. It is therefore aware of what migrations need
to run and which ones to rollback.

For example, a destination can have migrations and each of its actions can also
have specific migrations. Sources and destinations can have a `Migrate` function
which defines the migration logic to run for the adapter and will be executed
when running a migration against the adapter.

Migrations have a *up* and *down* logic allowing rollbacks. The wanderer leverages
the *supervisor* as described below so only one migration can run at a time on a
given target.

The `wanderer` adapter is optional. It is only needed when managing migrations from
a Blacksmith application.

Available `wanderer` adapters:
- [PostgreSQL](/blacksmith/options/wanderer/postgres) (`postgres`)

### Distributed environments

Most companies need to have zero downtime to maximize their service availability.

Blacksmith applications can optionally leverage a *supervisor*, which is in charge
of acquiring and releasing lock accesses in distributed environments. This ensures
resources in a multi-nodes cluster are accessed by a single instance of the
gateway, scheduler, and CLI to avoid collision when listening for events, executing
jobs, or running migrations.

The `supervisor` adapter is optional. It is only needed when running Blacksmith
applications in distributed environments.

Available `supervisor` adapters:
- [Consul](/blacksmith/options/supervisor/consul) (`consul`)

### Conclusion

By splitting Blacksmith into two services only, we ensure simplicity and consistency.
This also add a great separation of concerns for better security and scalability.
While the gateway may be exposed to the outside world, the scheduler can live
inside a private network with no public address.

The store keeps track of everything happening in the data pipeline. The pub / sub
allows realtime events forwarding between the gateway and scheduler.

The wanderer keeps track of migrations across every sources and destinations.

The supervisor allows to run Blacksmith applications in distributed environments.
