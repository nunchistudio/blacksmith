---
title: F.A.Q.
enterprise: false
---

# F.A.Q.

## Is Blacksmith a SaaS?

No. SaaS live on servers and domains you do not control. They usually expose HTTP
APIs so you can consume them by making HTTP requests to the endpoints they offer.

You could compare Blacksmith to a SaaS because organizations consume the public
APIs it exposes. However, unlike SaaS, developers mainly consume a [Go
API](https://pkg.go.dev/github.com/nunchistudio/blacksmith) in addition to the [HTTP
one](/blacksmith/http). This API lets you define your own data engineering
strategy as-code. In this way, it acts like a framework to build your own data
engineering SaaS. It runs on your servers and infrastructure.

Furthermore, a Blacksmith application also [exposes HTTP endpoints](/blacksmith/http).
Developers can therefore consume these to embed any kind of data in third-party
services and build custom dashboards on top of it.

## Can I use Blacksmith with an existing data stack?

Yes. Blacksmith is flexible and allows you to have external dependencies or pieces
of software in addition to the application built on top of it.

We are well aware analytics and marketing teams already use third-party services
such as Segment, Zapier, Fivetran, or dbt. Blacksmith can act both as an addition
or a substitute to these services, depending on your needs and the responsibilities
of each one.

## Why should I use Blacksmith instead of a no-code solution?

We think this wonderful strip from [CommitStrip](https://www.commitstrip.com/)
is appropriate.

You can [learn more about our product approach](/about).

![Low-code](/images/blacksmith/commitstrip.jpg)

## What are the prerequisites to learn Blacksmith?

Blacksmith is developed with the Go language. Therefore it is necessary to have
some experience with it.

The Blacksmith interfaces make it very easy to dive quickly into development.
Even if you just experienced Go for a few days it should be enough to understand
how Blacksmith works and create your first simple data engineering solution.

## Why did you choose Go?

Go has the right level of simplicity and abstraction we desired to build such a
product. Its performances and design choices lead us to pick it from the start,
without any regret.

Also, Go has become "*the language of the cloud*" in the past years. A major part
of cloud infrastructures and tools rely on Go such as Docker, Kubernetes, and
Terraform.

## What is the difference between an adapter and a driver?

An adapter adapts or extends the behavior of a Blacksmith application. For example,
the `supervisor` adapter allows data engineering in distributed systems and teams.

A driver is an *implementation* of an adapter. For example the PostgreSQL driver
for the `store` adapter let you use PostgreSQL as a store for Blacksmith.

## What is the license?

The use of Blacksmith is governed by the [Blacksmith Terms and Conditions](/legal/terms).

**Public repositories:**
- The Go public API is [available on GitHub](https://github.com/nunchistudio/blacksmith),
  and licensed under the [Apache License 2.0](https://github.com/nunchistudio/blacksmith/blob/main/LICENSE).
- The Go modules are [available on GitHub](https://github.com/nunchistudio/blacksmith-modules),
  and licensed under the [Apache License 2.0](https://github.com/nunchistudio/blacksmith-modules/blob/main/LICENSE).
- The UI kit is [available on GitHub](https://github.com/nunchistudio/blacksmith-ui),
  and licensed under the [Apache License 2.0](https://github.com/nunchistudio/blacksmith-ui/blob/main/LICENSE).
- The built-in dashboard is [available on GitHub](https://github.com/nunchistudio/blacksmith-dashboard),
  and licensed under the [Apache License 2.0](https://github.com/nunchistudio/blacksmith-dashboard/blob/main/LICENSE).
