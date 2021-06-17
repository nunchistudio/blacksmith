---
title: How Blacksmith works
enterprise: false
---

# Welcome to the Blacksmith documentation!

Before jumping right into the installation and creating your first data engineering
platform, it's important to understand how Blacksmith works. In this document
you will discover concepts and internal mechanisms useful for both **software**
and **data** engineers.

Blacksmith allows **software engineers** to write low-code ETL using the Go
language. It also allows **data engineers** to write templated SQL for TLT and
database migrations on top of one or multiple databases.

This approach lets engineering teams collaborate on managing a unique data
engineering platform from end-to-end:
![ETLT with Blacksmith](/images/blacksmith/overview.png)

## ETL with Go

### Sources and destinations

The ETL part takes care of Extracting data from *sources*, Transforming it when
necessary, and Loading it to *destinations*. Blacksmith can handle transformations
both before and after saving the jobs' payload, allowing to keep the original
events and still having optimized data for each job.

Each source has *triggers* and each destination has *actions*. In other words, a
Blacksmith application is able to receive events from sources' triggers, transform
data of these events, and run them against the appropriate destinations' actions
by creating jobs.

Sources can be websites, mobile applications, third-party services, databases, etc.
Destinations can be databases, data warehouses, or third-party services for analytics
or marketing.

A system can be treated both as a source *and* destination, allowing bidirectional
data flows.

Knowing this, we can imagine the ETL like this:
![Step 01](/images/blacksmith/how.001.png)

Triggers and actions inherit properties of their parent source or destination — such
as retry logic — but can override it in special cases.

Both triggers and actions have a business logic for handling enrichment, transformation,
and automating the data flow.

### The gateway and scheduler

Instead of using a lot of micro-services and over-engineering the whole process,
Blacksmith keep infrastructure as simple as possible with only two services: the
*gateway* and the *scheduler*.

The gateway is in charge of Extracting the events from sources' triggers. It can
happen on:
- [HTTP requests](/blacksmith/etl/extraction/triggers-http);
- [CRON schedules](/blacksmith/etl/extraction/triggers-cron);
- [CDC notifications](/blacksmith/etl/extraction/triggers-cdc);
- [Pub / Sub messages](/blacksmith/etl/extraction/triggers-sub).

The scheduler is in charge of Loading the events to actions of destinations,
creating *jobs*. There is an infinity of possibilities and workarounds for Loading
data to destinations. Blacksmith is unopiniated and does not enforce the use of
any schema or business logic.

Instead of locking users into a few limited patterns and still not covering every
needs you may have, we offer a collection of [production-ready modules](/blacksmith/modules).
It simplifies development, enforces best practices, and still allows a complete
freedom on how data is loaded when this is necessary.

![Step 02](/images/blacksmith/how.002.png)

These services can run on a single machine but should be splitted for production
use. This way, you gain more flexibility, security, and scalability.

### Database-as-a-Queue

The gateway needs a way to keep track of received events. The scheduler needs a
way to keep track of jobs to execute onto destinations, and their status (also
called *transitions*). So it can be aware of successes, failures, and discards.

The best way to achieve this is to have a persistent *store*. Every events received
by the gateway [**1**] are stored along the desired jobs [**2**]. The scheduler
knows how and when [**3**] to Load the data to appropriate destinations [**4**]
and can keep a complete history of jobs' transitions [**3**].

![Step 03](/images/blacksmith/how.003.png)

The `store` adapter is the only one required along the services.

Available drivers for the `store` adapter:
- [PostgreSQL](/blacksmith/options/store/postgres) (`postgres`)

### Enabling realtime

The store is perfect to persist events, jobs, and transitions. But we often need
to send data from sources to destinations in realtime.

We accomplish this by adding a pub / sub mechanism between the gateway and scheduler.
Once the event is received [**1**] and stored [**2**], it is then automatically
*published* by the gateway [**3**] to the scheduler that *subscribed* to the events
[**4**].

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

### Distributed semaphore

Most companies need to have zero downtime to maximize their service availability.
Blacksmith can run in a high availability (HA) to protect against outages by running
multiple Blacksmith instances.

Blacksmith applications can optionally leverage a *supervisor*, which is in charge
of acquiring and releasing lock accesses in distributed environments. This is also
known as a *distributed semaphore*.

The `supervisor` adapter ensures resources in a multi-nodes are accessed by a single
instance of the gateway, scheduler, and CLI to avoid collision when listening for
events, executing jobs, or running migrations.

As an example, the gateway service tries to acquire the lock for a specific
trigger [**2**] before Extracting data. If the key is already in use by an other
instance, the Extraction is skipped. It is applicable for CRON tasks and CDC
notifications only.

The scheduler service tries to acquire the lock for the specific action [**6**]
before Loading data. If the key is already in use by an other instance, the Load
is skipped.

![Step 05](/images/blacksmith/how.005.png)

The `supervisor` adapter is optional. It is highly recommended when running
Blacksmith applications in distributed environments or when several engineers can
work on database migrations at the same time.

Available drivers for the `supervisor` adapter:
- [Consul](/blacksmith/options/supervisor/consul) (`consul`)
- [PostgreSQL](/blacksmith/options/supervisor/postgres) (`postgres`)

## TLT with SQL

Engineering teams commonly leverage SQL for manipulating data. Blacksmith allows
to Transform and Load using a Python / Django-syntax like templating-language
applied to SQL:
```sql
{% if order %}
  INSERT INTO orders (amount, user_id) VALUES
    ({{ order.amount|floatformat:2 }}, '{{ user.id }}');
{% endif %}

```

Any source and destination can leverage the `warehouse` package to benefit templating
SQL. The `sqlike` module — which offers a unique and consistent approach for Loading
data into any and every SQL databases — makes use of this package.

It's possible to run two kinds of work on top of your data warehouse: *operations*
and *queries*.

An **operation** executes a statement without returning any rows. When running
an operation, it is automatically wrapped inside a transaction to ensure it is
either entirely commited or rolled back if any error occured. Operations should
be used for:
- Creating or refreshing views (materialized or not).
- Inserting, updating, or deleting data.

An operation should never be used for evolving the database schema, which could
impact software engineers. In this scenario, you should take advantage of database
migrations as introduced below.

A **query** executes a statement that returns rows, typically a `SELECT`. Queries
should be used for:
- `SELECT`ing samples of data.

Returned rows from a query can be downloaded both as CSV and JSON files:
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --csv --json

Executing queries:

  -> Executing ./queries/demo.sql...
     Writing CSV at ./queries/demo.csv...
     Writing JSON at ./queries/demo.json...
     Success!

```

## Migrations management

Data solutions often — to not say always — need to communicate with databases.
Managing and versioning database migrations is a difficult task to achieve,
especially across distributed teams in an organization. Also, the data schemas
of your data solution and the one of the databases need to be synchronized to
avoid schema incompatibilities as much as possible.

To resolve this challenge we add a *wanderer*. The wanderer does not handle the
migrations logic but keeps track of migration runs across every sources, triggers,
destinations, and actions. It is therefore aware of what migrations need to run
and which ones to rollback.

Migrations have a *up* and *down* logic allowing rollbacks:
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope "destination:sqlike(snowflake)" \
  --scope "destination:sqlike(mypostgres)"

```

When working with migrations, the `supervisor` adapter as described before should
be enabled so only one migration can run at a time for a given scope.

The `wanderer` adapter is optional. It is only needed when managing migrations
within a Blacksmith application.

Available drivers for the `wanderer` adapter:
- [PostgreSQL](/blacksmith/options/wanderer/postgres) (`postgres`)

## From end-to-end

Blacksmith allows **software engineers** to write low-code ETL using the Go
language. It also allows **data engineers** to write templated SQL for TLT and
database migrations on top of one or multiple databases.

By offering ETL, TLT, and databases migrations, Blacksmith is the central piece
of your data engineering workflow leading you to save months of customized and
professional work.
