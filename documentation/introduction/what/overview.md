---
title: Introduction to Blacksmith
enterprise: false
---

# Introduction to Blacksmith

Blacksmith is a Go framework specifically designed for data engineering teams. It
allows you to design, build, and deploy reliable data platforms in a consistent
way. The goal of Blacksmith is to address as many pain points as possible engineering
teams encounter while working on data solutions.

Any team that is building — or think about building — a complete data platform knows
the tremendous amount of work needed to properly accomplish this mission. Think
of Blacksmith as the central piece of your data engineering workflow, leading you
to save months of customized and professional work.

## Features and benefits

The **Blacksmith Standard Edition** addresses the technical complexity of data
engineering.

- **Architecture reliability:** With a state-of-the-art queue and retry management
  system, Blacksmith takes all the job management pain away. Also, the services
  gracefully shutdown without interrupting any active connections and tasks.
- **Flow automation:** Whenever an event happens or whenever data is received by a
  destination, other events can automatically be triggered using original or transformed
  data. Each event can have its own scheduling options. This let you have a central
  and complete control over the data flows across your stack. 
- **Universal error handling:** Blacksmith offers a universal way of handling errors,
  failures, retries, and discards across events regardless of their origin, making
  their management easier. 
- **Data management:**
  - **Data collection:** Blacksmith allows to extract any kind of data whether you
    are collecting it from HTTP APIs, CRON schedules, CDC notifications, or Pub / Sub
    messages with a consistent interface.
  - **Data governance:** With Blacksmith as the single source of truth for all your
    data, it is easy for organizations to protect customers' data against modern
    threats. This can also maintain customer trust and automate compliance with
    GDPR, CCPA, and whatever comes next.
  - **Data transformation:** Because each and every system has a specific design,
    you often need to validate and transform the data before sending it to destinations.
    Any events received from various sources can be transformed to match any events
    of any destinations.
  - **Data enrichment:** Sometimes the data you collect through an event is not
    enough and you need to enrich it by adding some properties or metadata. Blacksmith
    is flexible to let any kind of business logic act as a middleware between the
    sources and the destinations.
  - **Data distribution:** Once the data is transformed and enriched, you can send
    it to any destinations. When at any point in time, destinations will be in a
    state of failure, Blacksmith will handle perfectly the need for retries or not.
  - **Data warehousing:** A data warehouse centrally stores all of your data and
    can be used by Blacksmith as a destination for any kind of events. Since Blacksmith
    considers a data warehouse as a traditional destination, you can use time-series
    or realtime analytical database as a data warehouse.
  - **Data ownership:** By not relying on a cloud provider for running your data
    platform, you have complete control over your data and manage its entire
    lifecycle across your stack.
- **Simplicity, clarity, performance:** Even though the data platforms can sometimes
  be complicated, Blacksmith makes it really easy to understand the desired behavior
  of each event.
- **Adapter-based:** The flexibility brought by Blacksmith allows organizations to
  use the technologies they already love and adopted in their technical stack.
  Any piece of technology can be plugged into a Blacksmith application.
- **Cloud-native & multi-cloud:** Data platforms built on top of Blacksmith can
  be deployed on any infrastructure or cloud-provider. There is no cloud lock-in.
- **Docker-based workflow:** Blacksmith takes environment parity to the next level
  by leveraging Docker for most of operations. If your application works on your
  machine, it works on any Docker host with no modification.
- **Scaffolding:** The Blacksmith CLI lets you generate any kind of resources
  in a simple command-line to extend your application as quickly and easily as
  possible.

The **Blacksmith Enterprise Edition** addresses the complexity of collaboration
and governance across multi-team and multi-scope data solutions.

- **REST API:** The Enterprise Edition comes with a REST API so organizations can
  bring data platform information as well as historical and real time events / jobs
  into third-party applications.
- **Server health-checks:** Instead of using the Standard adapters, organizations
  can switch to the Enterprise Edition of the gateway and scheduler to add these
  services to their service mesh. More details about adapters in the next guide.
- **Distributed environments:** Blacksmith applications can be deployed in distributed
  and high-available environments. By using a distributed lock mechanism, we ensure
  strong data consistency and stronger fault-tolerance across nodes.
- **Database migrations:** Versioning and migrating database schema is difficult,
  especially if there is more than one engineer working on the code base. By using
  the Blacksmith Enterprise Edition, organizations can run and version database
  migrations smoothly with no conflicts across teams.
