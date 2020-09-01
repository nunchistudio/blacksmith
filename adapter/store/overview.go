/*
Package store provides the development kit for working with a database-as-a-queue,
so the jobs queue can be persisted in a datastore.

At any point in time, destinations will be in a state of failure. By using a store,
in tandem with a scheduler, Blacksmith applications are sure to build a reliable
system for delivering jobs at scale and keep track of successes, failures, and
discards.
*/
package store
