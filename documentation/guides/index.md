---
title: Guides & Tutorials
enterprise: false
---

# Guides & Tutorials

Following is a cheatsheet of the command lines to generate files in a few seconds.
Please refer to each section on the left navigation for details and examples.

## Data Extraction

Generate a source:
```bash
$ blacksmith generate source --name <name> --path ./relative/path --migrations
```

Generate a trigger using HTTP mode:
```bash
$ blacksmith generate trigger --name <name> --mode http --path ./relative/path --migrations
```

Generate a trigger using CRON mode:
```bash
$ blacksmith generate trigger --name <name> --mode cron --path ./relative/path --migrations
```

Generate a trigger using CDC mode:
```bash
$ blacksmith generate trigger --name <name> --mode cdc --path ./relative/path --migrations
```

Generate a trigger using subscription mode:
```bash
$ blacksmith generate trigger --name <name> --mode sub --path ./relative/path --migrations
```

## Data Transformation

Generate a flow:
```bash
$ blacksmith generate flow --name <name> --path ./relative/path
```

## Data Load

Generate a destination:
```bash
$ blacksmith generate destination --name <name> --path ./relative/path --migrations
```

Generate an action:
```bash
$ blacksmith generate action --name <name> --path ./relative/path --migrations
```

## Database migrations

Generate a migration:
```bash
$ blacksmith generate migration --name <name> --path ./relative/path
```
