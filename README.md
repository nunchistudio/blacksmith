# Blacksmith

After a few years we decided to crystallize all of our data engineering best
practices into a product, on top of which organizations can layer cloud-assisted
data solutions.

Blacksmith is a Go framework specifically designed for data engineering teams. It
allows you to design, build, and deploy reliable data platforms in a consistent
way. The goal of Blacksmith is to address as many pain points as possible engineering
teams encounter while working on data solutions.

Any team that is building — or think about building — a complete data platform knows
the tremendous amount of work needed to properly accomplish this mission. Think
of Blacksmith as the central piece of your data engineering workflow, leading you
to save months of customized and professional work.

By leveraging Blacksmith, organizations benefit a single source of truth for all
their data with a unique developer experience:
```bash
curl -X GET 'https://example.com/admin/store/jobs' \
  -d events.sources_in=cms \
  -d events.sources_in=crm \
  -d jobs.destinations_in=warehouse \
  -d jobs.actions_in=register \
  -d jobs.status_in=discarded \
  -d offset=0 -d limit=100
```

## Product offerings

Blacksmith is available in two editions:

- **Blacksmith Standard Edition** addresses the technical complexity of data
  engineering. It is and will always be free.
- **Blacksmith Enterprise Edition** addresses the complexity of collaboration
  and governance across multi-team and multi-scope data solutions.

- [Compare Editions](https://nunchi.studio/blacksmith/compare)

## Links

- [Learn more on Nunchi website](https://nunchi.studio/blacksmith)
- [Reference on Go developer portal](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc)

## Professional services

Along consulting and training, we provide different product offerings as well as
different levels of support.

- [Discover our services](https://nunchi.studio/blacksmith/support)

## License

Repository licensed under the [Apache License, Version 2.0](./LICENSE).

By downloading, installing, and using Blacksmith, you agree to the
[Blacksmith Terms and Conditions](https://nunchi.studio/blacksmith/legal/terms).
