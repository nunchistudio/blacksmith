/*
Package gateway provides the development kit for working with CRON and HTTP events.
It receives events from sources such as websites, mobile applications, or databases
notifications.

The gateway takes care of HTTP requests, CRON tasks or forever listeners for
registering events in a synchronous way, whereas the scheduler takes care of jobs
asynchronously to handle failures and retries against destinations. The gateway and
scheduler must work in tandem to have a reliable data pipeline, both using the pubsub
and store packages.

The HTTP server of the gateway is designed to work over HTTP with REST specification.
Even though we really like gRPC and GraphQL, we decided to only focus on REST for
several reasons:
  - The tooling, ecosystem, and developer experience of gRPC and GraphQL are still
    frustrating to work with and require more development to meet success.
  - Sources and destinations can all communicate using HTTP REST. However, it is
    not uncommon that sources and destinations are unable to communicate with gRPC
    or GraphQL. The most notable example are webhooks. We could support gRPC and
    GraphQL but this will add a lot of complexity for something not needed that
    can still be done at the application layer using a custom HTTP handler.

Both gRPC and GraphQL have good designs that unfortunately don't fit our needs.
Instead of trying to work with these designs, we implemented interfaces to mimic
some of the features they provide and we love, such as resolvers, while still
providing a smooth developer experience.

A gateway adapter can be generated using the Blacksmith CLI:

  $ blacksmith generate gateway

Note: Adapter generation using the Blacksmith CLI is a feature only available in
Blacksmith Enterprise.
*/
package gateway
