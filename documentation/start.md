---
title: Starting an instance
enterprise: false
---

# Starting an instance

## Validation

You can compile and validate an application without starting any service by running:
```bash
$ blacksmith validate
```

This will create a `.blacksmith` directory at the top of your application, including
your application compiled as a Go plugin.

## Running an instance

Based on what we learned in the previous guide, we can start the desired stack with:
```bash
$ docker-compose up
```

Before starting an instance, Blacksmith will automatically compile and validate
your application.
