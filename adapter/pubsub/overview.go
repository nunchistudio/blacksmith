/*
Package pubsub provides the development kit for working with Publish / Subscribe
systems.

This feature shall be used by the gateway as a publisher and by the scheduler as
a subscriber. This way, the gateway can publish events in realtime to the scheduler
that then will take care of forwarding the events' jobs to the desired destinations.
*/
package pubsub
