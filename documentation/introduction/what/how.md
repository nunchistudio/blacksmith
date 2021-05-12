---
title: How it works
enterprise: false
---

# How it works

Blacksmith is a low-code ecosystem offering a complete and consistent approach for
building self-managed data engineering solutions. It allows organizations to process,
trust, and visualize all their data flowing from end-to-end in a consistent way.

Blacksmith is available in two Editions:
- **Blacksmith Standard Edition** addresses the technical complexity of data
  engineering. It is and will always be free. This Edition focuses on building
  ETL platforms for better data quality and stronger data reliability across systems.
- **Blacksmith Enterprise Edition** addresses the complexity of collaboration and
  governance across multi-team and multi-scope data solutions. This Edition adds
  advanced features for the enterprises, such as migrations management, distributed
  semaphore, and a built-in dashboard.

Any team that is building — or think about building — such a platform knows the
tremendous amount of work needed to properly accomplish this mission. Think of
Blacksmith as the central piece of your data engineering workflow, leading you to
save months of customized and professional work.

![Data engineering with Blacksmith](/images/blacksmith/approach.png)

![Blacksmith Dashboard](/images/blacksmith/dashboard.002.png)

## Sources and destinations

The main role of Blacksmith is to act as a data pipeline. This role takes care of
Extracting data from *sources*, Transforming it when necessary, and Loading it
to *destinations*. Blacksmith can handle transformations both before and after
saving the jobs' payload, allowing to keep the original events and still having
optimized data for each job.

Each source has *triggers* and each destination has *actions*. In other words, a
Blacksmith application is able to receive events from sources' triggers, transform
data of these events, and run them against the appropriate destinations' actions
by creating jobs.

Sources can be websites, mobile applications, third-party services, databases, etc.
Destinations can be databases, data warehouses, or third-party services for analytics
or marketing.

A system can be treated both as a source *and* destination, allowing bidirectional
data flows.

![Step 01](/images/blacksmith/how.001.png)

Triggers and actions inherit properties of their parent source or destination — such
as retry logic — but can override it in special cases.

Both triggers and actions have a business logic for handling enrichment, transformation,
and automating the data flow.

## The gateway and scheduler

Instead of using a lot of micro-services and over-engineer the whole process,
Blacksmith keep data pipelines as simple as possible with only two services: the
*gateway* and the *scheduler*.

The gateway is in charge of extracting the events from sources' triggers. It can
happen on:
- [HTTP requests](/blacksmith/guides/extraction/triggers-http);
- [CRON schedules](/blacksmith/guides/extraction/triggers-cron);
- [CDC notifications](/blacksmith/guides/extraction/triggers-cdc);
- [Pub / Sub messages](/blacksmith/guides/extraction/triggers-sub).

The scheduler is in charge of loading the events into actions of destinations,
creating *jobs*. There is an infinity of possibilities and workarounds for loading
data to destinations. Blacksmith is unopiniated and does not enforce the use of
any schema or business logic. 

Instead of locking users into a few limited patterns and still not covering every
needs you may have, we offer a collection of [production-ready modules](/blacksmith/modules).
It simplifies development, enforces best practices, and still allows a complete
freedom on how data is loaded when this is necessary.

![Step 02](/images/blacksmith/how.002.png)

These services can run on a single machine but should be splitted for production
use. This way, you gain more flexibility, security, and scalability.

## Database-as-a-Queue

The gateway needs a way to keep track of received events. The scheduler needs a
way to keep track of jobs to execute onto destinations, and their status (also
called *transitions*). So it can be aware of successes, failures, and discards.

The best way to achieve this is to have a persistent *store*. Every events received
by the gateway are stored along the desired jobs. The scheduler knows how and when
to load the data to appropriate destinations and can keep a record of jobs'
transitions. 

![Step 03](/images/blacksmith/how.003.png)

The `store` adapter is the only one required along the services.

Available drivers for the `store` adapter:
- [PostgreSQL](/blacksmith/options/store/postgres) (`postgres`)

## Enabling realtime

The store is perfect to persist events, jobs, and status. But we often need to
send data from sources to destinations in realtime.

We accomplish this by adding a pub / sub mechanism between the gateway and scheduler.
Once the event is received and stored, it is then automatically *published* by the
gateway to the scheduler that *subscribed* to the events.

![Step 04](/images/blacksmith/how.004.png)

The `pubsub` adapter is optional. When no adapter is provided or the data doesn't
need to be loaded in realtime into the destination, the scheduler will load it
on a given schedule (order matters):
- the destination's action schedule; or
- the destination's schedule; or
- the default schedule.

Once configured, the `pubsub` adapter allows realtime message extraction from
queues / topics / subscriptions.

Available drivers for the `pubsub` adapter:
- [AWS SNS / SQS](/blacksmith/options/pubsub/aws) (`aws/snssqs`)
- [Azure Service Bus](/blacksmith/options/pubsub/azure) (`azure/servicebus`)
- [Google Pub / Sub](/blacksmith/options/pubsub/google) (`google/pubsub`)
- [Apache Kafka](/blacksmith/options/pubsub/kafka) (`kafka`)
- [NATS](/blacksmith/options/pubsub/nats) (`nats`)
- [RabbitMQ](/blacksmith/options/pubsub/rabbitmq) (`rabbitmq`)

## Versioning migrations

Data solutions often — to not say always — need to communicate with databases.
Managing and versioning database migrations is a difficult task to achieve,
especially across distributed teams in an organization. Also, the data schemas
of your data solution and the one of the databases need to be synchronized to
avoid schema incompatibilities as much as possible.

To resolve this challenge we add a *wanderer*. The wanderer does not handle the
migrations logic but keeps track of migration runs across every sources, triggers,
destinations, and actions. It is therefore aware of what migrations need to run
and which ones to rollback.

Migrations have a *up* and *down* logic allowing rollbacks.

When working with migrations, the `supervisor` adapter as described below should
be enabled so only one migration can run at a time for a given scope.

The `wanderer` adapter is optional. It is only needed when managing migrations
within a Blacksmith application.

Available drivers for the `wanderer` adapter:
- [PostgreSQL](/blacksmith/options/wanderer/postgres) (`postgres`)

## Distributed environments

Most companies need to have zero downtime to maximize their service availability.
Blacksmith can run in a high availability (HA) to protect against outages by running
multiple Blacksmith instances.

Blacksmith applications can optionally leverage a *supervisor*, which is in charge
of acquiring and releasing lock accesses in distributed environments. This is also
known as a *distributed semaphore*.

The `supervisor` adapter ensures resources in a multi-nodes are accessed by a single
instance of the gateway, scheduler, and CLI to avoid collision when listening for
events, executing jobs, or running migrations.

The `supervisor` adapter is optional. It is highly recommended when running
Blacksmith applications in distributed environments or when several engineers can
work on database migrations at the same time.

Available drivers for the `supervisor` adapter:
- [Consul](/blacksmith/options/supervisor/consul) (`consul`)
- [PostgreSQL](/blacksmith/options/supervisor/postgres) (`postgres`)

## Conclusion

By splitting Blacksmith into two services only, we ensure simplicity and consistency.
This also adds a great separation of concerns for better security and scalability.
While the gateway may be exposed to the outside world, the scheduler can live
inside a private network with no public address.

The store keeps track of everything happening in the data pipeline. The pub / sub
allows realtime events forwarding between the gateway and scheduler.

The wanderer keeps track of migrations across every sources and destinations.

The supervisor provides coordination to run Blacksmith applications in distributed
environments. It also allows several engineers to work at the same time on the
application with no conflicts when running migrations.
