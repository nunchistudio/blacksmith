---
title: Starting an instance
enterprise: false
---

# Starting an instance

> If this is not already the case, we invite you to download the Blacksmith CLI
  in the [downloads section](https://nunchi.studio/blacksmith/downloads).

## Validation

You can compile and validate an application without starting any service by running:
```bash
$ blacksmith validate
```

This will create a `.blacksmith` directory at the top of your application, including
your application compiled as a Go plugin.

## Running an instance

Based on what we learned in the previous guide, we can start the Docker stack with:
```bash
$ docker-compose up
```

Before starting an instance, Blacksmith will automatically compile and validate
your application.
