/*
Package gateway provides the development kit for exposing a service receiving
incoming events from registered triggers. It receives events from sources such as
websites, mobile applications, or databases notifications.

The gateway can handle HTTP requests, CRON tasks, CDC notifications, and Pub / Sub
messages. It registers events using the store adapter, whereas the scheduler takes
care of distributing jobs asynchronously to handle failures and retries against
destinations.
*/
package gateway
