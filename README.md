# Blacksmith

After a few years we decided to crystallize all of our data engineering best
practices into a product, on top of which organizations can layer cloud-assisted
data solutions.

![Data engineering with Blacksmith](https://nunchi.studio/images/blacksmith/approach.png)

Blacksmith is a programmable low-code ecosystem, offering a complete and consistent
data engineering solution. It allows you to design, build, and deploy reliable data
engineering platforms in a consistent way. The goal of Blacksmith is to address as
many pain points as possible data engineering teams encounter.

Any team that is building — or think about building — a complete data engineering
platform knows the tremendous amount of work needed to properly accomplish this
mission. Think of Blacksmith as the central piece of your data engineering workflow,
leading you to save months of customized and professional work.

By leveraging Blacksmith, organizations benefit a single source of truth for all
their data with a unique developer experience.

Powerful REST API ([Enterprise Edition](https://nunchi.studio/blacksmith/editions)):
```bash
$ curl --request GET --url 'https://example.com/admin/api/store/jobs' \
  -d events.sources_in=cms \
  -d events.sources_in=crm \
  -d jobs.destinations_in=warehouse \
  -d jobs.actions_in=register \
  -d jobs.status_in=discarded \
  -d offset=0 -d limit=100
```

Built-in dashboard ([Enterprise Edition](https://nunchi.studio/blacksmith/editions)):
![Blacksmith Dashboard](https://nunchi.studio/images/blacksmith/dashboard.002.png)

## Product offerings

**Blacksmith is not an open-source software.** This repository only holds the
public Go APIs, allowing organizations to build reliable data engineering solutions
on top of Blacksmith using Go. Blacksmith itself is [built and distributed in a
Docker image](https://github.com/nunchistudio/blacksmith-docker).

Blacksmith is available in two Editions:
- **Blacksmith Standard Edition** addresses the technical complexity of data
  engineering. It is and will always be free.
- **Blacksmith Enterprise Edition** addresses the complexity of collaboration
  and governance across multi-team and multi-scope data solutions.

- [Compare Editions](https://nunchi.studio/blacksmith/editions)

## Links

- [Learn more on Nunchi website](https://nunchi.studio/blacksmith)

## Professional services

Along consulting and training, we provide different product offerings as well as
different levels of support.

- [Discover our services](https://nunchi.studio/support)

## License

Repository licensed under the [Apache License, Version 2.0](./LICENSE).

By downloading, installing, and using Blacksmith, you agree to the
[Blacksmith Terms and Conditions](https://nunchi.studio/legal/terms).
