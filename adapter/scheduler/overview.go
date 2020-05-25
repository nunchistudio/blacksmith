/*
Package scheduler provides the development kit for working with a scheduler that
receives events from the pubsub package and, used in tandem with the store package,
is in charge of the reliability of the event delivery to destinations.

Whereas the gateway takes care of incoming events, the scheduler is in charge of
handling jobs to destinations in an asynchronous way.

A scheduler adapter can be generated using the Blacksmith CLI:

  $ blacksmith generate scheduler

Note: Adapter generation using the Blacksmith CLI is a feature only available in
Blacksmith Enterprise.
*/
package scheduler
