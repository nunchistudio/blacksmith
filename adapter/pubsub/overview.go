/*
Package pubsub provides the development kit for working with Publish / Subscribe
mechanism between the gateway and scheduler services.

This feature is used by the gateway as a publisher and by the scheduler as a
subscriber. This way, the gateway can publish jobs in realtime to the scheduler
that then will take care of forwarding them to the appropriate actions.
*/
package pubsub
