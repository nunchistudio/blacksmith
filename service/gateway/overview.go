/*
Package gateway provides the development kit for working with incoming events.
It receives events from sources such as websites, mobile applications, or databases
notifications.

The gateway can handle HTTP requests, CRON tasks, and CDC notifications for
registering events, whereas the scheduler takes care of distributing jobs
asynchronously to handle failures and retries against destinations.

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
*/
package gateway
