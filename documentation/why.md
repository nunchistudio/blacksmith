---
title: Why Blacksmith
enterprise: false
---

# Why Blacksmith

Blacksmith is a Software Development Kit specifically designed for data engineering
teams. It allows you to build reliable data pipelines in a consistent way. Whether
you are collecting data from HTTP APIs, CRON tasks or ongoing event listeners and
whether you are loading this data in realtime to third-party services or using a
specific schedule to data warehouses.

Any team that is building — or think about building — a complete data pipeline knows
the tremendous amount of work needed to properly accomplish this mission. Think
of Blacksmith as the central piece of your data engineering workflow, leading you
to save months of customized and professional ETL work.

## Approach and philosophy

The goal of Blacksmith is to address as many pain points as possible that data
engineering teams encounter while working on data solutions.

Blacksmith is a SDK, not a framework, which means it gives the flexibility needed
when working on data solutions. You can use all the features Blacksmith offers or
just a set of the ones needed to solve your problems.

Applications built on top of Blacksmith can embed or be embedded by existing Go
applications. Also, any HTTP handlers can run side-by-side with your data pipeline
so you can have traditional REST APIs sharing the same code base and middlewares.

Blacksmith is a lightweight suite of *interfaces* designed to solve data engineering
challenges without carrying any undesired features. It is up to engineering teams
to bring additional utilities if needed.

The *adapters* we offer are implementations of those interfaces and can replaced
by custom ones. They are compiled as Go plugins and are not open-source.

## Features and benefits

The **Blacksmith Standard** version addresses the technical complexity of data
engineering.

- **Synchrounous & asynchronous events:** Whether you are collecting data from HTTP
  APIs, CRON tasks or ongoing event listeners, Blacksmith handles them in a consistent
  way with a unique event handling interface.
- **Architecture reliability:** With a state-of-the-art queue and retry management
  system, Blacksmith makes it very difficult to lose data between an event and the
  finale data destination.
- **Flow automation:** Whenever an event happens or whenever data is received by a
  destination, other events can automatically be triggered using original or transformed
  data. Each event can have its own scheduling options. This let you have a central
  and complete control over the data flows across your stack. 
- **Universal error handling:** Blacksmith offers a universal way of handling errors,
  failures, retries, and discards across events regardless of their origin, making
  their management easier. 
- **Data management:**
  - **Data governance:** With Blacksmith as the single source of truth for all your
    data, it is easy for organizations to protect customers' data against modern
    threats. This can also maintain customer trust and automate compliance with
    GDPR, CCPA, and whatever comes next.
  - **Data collection:** Blacksmith allows to extract any kind of data whether you
    are collecting it from HTTP APIs, CRON tasks or ongoing event listeners with a
    consistent interface.
  - **Data transformation:** Because each and every system has a specific design,
    you often need to validate and transform the data before sending it to destinations.
    Any events received from various sources can be transformed to match any events
    of any destinations.
  - **Data enrichment:** Sometimes the data you collect through an event is not
    enough and you need to enrich it by adding some properties or metadata. Blacksmith
    is flexible to let any kind of business logic act as a middleware between the
    sources and the destinations. This allows to centrally store external data from
    third-party services into a data warehouse for example.
  - **Data distribution:** Once the data is transformed and enriched, you can send
    it to any destinations. When at any point in time, destinations will be in a
    state of failure, Blacksmith will handle perfectly the need for retries or not.
  - **Data warehousing:** A data warehouse centrally stores all of your data and
    can be used by Blacksmith as a destination for any kind of events. Since Blacksmith
    considers a data warehouse as a traditional destination, you can use time-series
    or realtime analytical database as a data warehouse.
  - **Data ownership:** By not relying on a cloud provider for running your data
    pipeline, you have complete control over your data and manage its entire
    lifecycle across your stack.
- **Simplicity, clarity, performance:** Like written before, Blacksmith is very
  light. Even though the data pipelines can sometimes be complicated, Blacksmith
  makes it really easy to understand the desired behavior of each event.
- **Adapter-based:** The flexibility brought by Blacksmith allows organizations to
  use the technologies they already love and adopted in their technical stack.
  Any piece of technology can be plugged into a Blacksmith application.
- **Cloud-native & multi-cloud:** Data pipelines built on top of Blacksmith can
  be deployed on any infrastructure or cloud-provider. There is no lock-in.

The **Blacksmith Enterprise** version addresses the complexity of collaboration
and governance across multi-team and multi-scope data solutions.

- **REST API:** The Enterprise version comes with a REST API so organizations can
  bring data pipeline information as well as historical and real time events / jobs
  into third-party applications.
- **Database migrations:** Versioning and migrating database schema is difficult,
  especially if there is more than one engineer working on the code base. By using
  the Blacksmith Enterprise version, organizations can run and version database
  migrations smoothly with no conflicts across teams.
- **Command-Line Interface:** The Blacksmith CLI lets you generate any kind of
  adapter in a simple command-line as well as manage database migrations.
- **Server health-checks:** Instead of using the standard adapters, organizations
  can switch to the Enterprise version of the gateway and scheduler to add these
  services to their service mesh. More details about adapters in the next guide.
- **Dashboard:** We created a dashboard serving as a reference implementation of
  the REST API. It allows you to visualize everything that happened or is happening
  in a Blacksmith application. It can be customized in any way needed to fit the
  requirements of businesses.
