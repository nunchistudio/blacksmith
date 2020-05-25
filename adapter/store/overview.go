/*
Package store provides the development kit for working with a database-as-a-queue,
so the jobs queue can be persisted in a datastore.

At any point in time, destinations will be in a state of failure. By using a store,
in tandem with a scheduler, Blacksmith applications are sure to build a reliable
system for delivering events at scale and keep track of successes, failures, errors,
and jobs' transitions.

The store shall be immutable: no rows shall be updated or removed.

A store adapter can be generated using the Blacksmith CLI:

  $ blacksmith generate store

Note: Adapter generation using the Blacksmith CLI is a feature only available in
Blacksmith Enterprise.
*/
package store
