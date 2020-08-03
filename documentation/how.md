---
title: How it works
enterprise: false
---

# How it works

Wether it is your first time doing data engineering or not, this guide is important
since it describes some specifities of how Blacksmith works.

Let's say that, for now, Blacksmith should look like this to you: a blackbox.

![Step 01](/images/blacksmith/how.001.png)

Some things happen, but you don't know why or how.

## Sources and destinations

The role of a data pipeline is mainly to take care of Extracting data from *sources*,
Transforming it when necessary, and Loading it to *destinations*. Each source and
destination have *events*.

In other words, the blackbox that we imagined earlier is now able to receive events
from sources, transform data of these events, and sent it to appropriate events of
destinations.

Sources can be websites, mobile applications, third-party services, databases, etc.
Destinations can be databases, data warehouses, or third-party services for analytics
or marketing.


![Step 02](/images/blacksmith/how.002.png)

Events inherit properties of their parent source or destination — such as retry
logic — but can override it in special cases.

Both sources' and destinations' events have a business logic for handling enrichment,
transformation, and the data flow with other events.

## The gateway and scheduler

Instead of using micro-services and over-engineer the whole process, Blacksmith keep
your data pipeline simple with only two services: the *gateway* and the *scheduler*.

The gateway is in charge of receiving the events from sources. The scheduler is in
charge of sending the events to destinations.

![Step 03](/images/blacksmith/how.003.png)

These services can run on a single machine but should be splitted for production
use. This way, you gain more flexibility about security and scalability.

Nunchi offers `standard` adapters for both the `gateway` and `scheduler` interfaces.
They can be replaced by any in-house adapters if desired.

## Database-as-a-queue

The gateway needs a way to keep track of received events and *jobs* to execute onto
destinations. The scheduler needs a way to keep track of job status (also called
*transitions*) so it can be aware of successes, failures, and discards.

The best way to achieve this is to have a persistent *store*. Every events received
by the gateway are stored along the desired jobs. The scheduler now knows how and
when to load the data to appropriate destinations and can keep a record of jobs' transitions. 

![Step 04](/images/blacksmith/how.004.png)

Nunchi offers `store` adapters for:
- PostgreSQL (`postgres`)

## Enabling realtime

The store is perfect to persist events, jobs, and status. But we often need to
send data from sources to destinations in realtime.

We accomplish this by adding a pub / sub feature between the gateway and scheduler.
Once the event is received and stored, it is then *published* to the scheduler
that *subscribed* to the events.

![Step 05](/images/blacksmith/how.005.png)

The `pubsub` adapter is optional. When no adapter is provided, the scheduler will
send data to destinations and schedule retries given (order matters):
- the event's schedule; or
- the destination's schedule; or
- the default schedule.

Nunchi offers `pubsub` adapters for:
- Kafka (`kafka`)
- NATS (`nats`)
- RabbitMQ (`rabbitmq`)

## Versioning migrations

Data pipelines often — to not say always — need to communicate with databases.
Managing and versioning database migrations is a difficult task to achieve,
especially across distributed teams in an organization. Also, the data schemas
of the data pipeline and the one of the databases need to be synchronized to
avoid schema incompatibilities as much as possible.

To resolve this challenge we add a *wanderer*. The wanderer does not handle the
migrations logic but keeps track of migration runs across every adapters and
events. It is therefore aware of what migrations need to run and which ones to
rollback.

For example, a destination can have migrations and each of its events can also
have specific migrations. Sources and destinations have a `Migrate` function
which defines the migration logic to run for the adapter and will be executed
when running a migration against the adapter.

Migrations have a *up* and *down* logic allowing rollbacks. The wanderer leverages
the supervisor as described below so only one migration can run at a time on a
given target. This allows concurrency controls and avoid race conditions.

Nunchi offers `wanderer` adapters for:
- PostgreSQL (`postgres`)

> The *wanderer* is only available using Blacksmith Enterprise Edition.

## Distributed environments

Most companies need to have zero downtime to maximize their service availability.

Blacksmith applications can optionnaly leverage a *supervisor*, which is in charge
of acquiring and releasing lock accesses in distributed environments. This ensures
resources in a multi-nodes cluster are accessed by a single instance of the
gateway and the scheduler to avoid collision when listening for events, executing
jobs, or running migrations.

Nunchi offers `supervisor` adapters for:
- Consul (`consul`)

> The *supervisor* is only available using Blacksmith Enterprise Edition.

## Conclusion

By splitting Blacksmith into two services only, we ensure simplicity and consistency.
This also add a great separation of concerns for better security and scalability.
While the gateway might be exposed to the outside world, the scheduler can live
inside a private network with no public address.

The store keeps track of everything happening in the data pipeline. The pub / sub
allows realtime events forwarding between the gateway and scheduler.

The wanderer keeps track of migrations across every adapters and their events.
