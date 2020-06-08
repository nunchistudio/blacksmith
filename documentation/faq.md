---
title: F.A.Q.
enterprise: false
---

# F.A.Q.

## Can I use Blacksmith with an existing data pipeline?

Yes. Blacksmith is flexible and allows you to have external dependencies or pieces
of software in addition to the one using the SDK.

We are well aware analytics and marketing teams already use third-party services
like Segment and Zapier. Blacksmith can act as an addition or a substitute to these
services depending on your needs.

## What are the prerequisites to learn Blacksmith?

Blacksmith is developed with the Go language. Therefore it is necessary to have
some experience with it.

The Blacksmith interfaces make it very easy to dive quickly into development.
Even if you just experienced Go for a few days it may be enough to understand
how Blacksmith works and create a first simple data pipeline.

## Why did you choose Go?

Go has the right level of simplicity and abstraction we desired to build such a
product. Its performances and design choices lead us to pick it from the start,
without any regret.

Also, Go has become "the language of the cloud" in the past years. A major part
of cloud infrastructures and tools rely on Go such as Docker, Kubernetes, and
Terraform.

## How adapters work?

An adapter is an *implementation* of an *interface*. For example the PostgreSQL
store adapter let you use PostgreSQL as a store for Blacksmith.

Nunchi offers production-ready adapters so you don't have to build everything
from scratch. The adapters we provide are automatically downloaded, installed,
and updated accordingly to the Blacksmith version your application depends on.
The source code is not open and each adapter is licensed under the
[Blacksmith Adapter License](/licenses/blacksmith-adapter).

## Why the adapters are not open-source?

By keeping the source code closed, we also keep a better flexibility for Enterprise
versions and features. To develop a large-scale, technically-complex product,
and ensure long-term sustainability, we need to have a strong focus on the company
behind this. This means having different kinds of revenue streams alongside services
and support.

In a near future, we will offer `enterprise` adapters in addition to the `standard`
ones for the `gateway` and `scheduler` interfaces. More details coming soon.
