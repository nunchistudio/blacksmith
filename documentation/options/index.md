---
title: Blacksmith configuration
enterprise: false
---

# Blacksmith configuration

The Blacksmith configuration reference includes details, notes, and environment
variables for configuring an application.

The first step is to configure the `gateway` and `scheduler` services. Then, the
`store` adapter since this is the only one mandatory.

The `pubsub` adapter must be configured for realtime data loading and extracting
data from Pub / Sub messages.
