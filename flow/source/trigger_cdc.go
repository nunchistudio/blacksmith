package source

/*
ModeCDC is used to indicate the event is a forever loop. It is used for ongoing
listeners such as databases notifications.
*/
var ModeCDC = "cdc"

/*
TriggerCDC is the interface used for triggers using a Change-Data-Capture logic.
This can be used for listening to databases notifications.

A new CDC trigger can be generated using the Blacksmith CLI:

  $ blacksmith generate trigger --name <name> --mode cdc [--path <path>] [--migrations]
*/
type TriggerCDC interface {

	// Extract in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Extract(*Toolkit, *Notifier)
}

/*
Notifier includes channels the trigger can listen or write to.

Since this mode is asynchronous, there is no way for the gateway to know when the
trigger is done. To gracefully shutdown like in a synchronous mode, the function
receives a message on `IsShuttingDown` and must write to `Done` whenever the function
is ready to exit. Otherwise, the gateway will block until `true` is received on
`Done`.

Example:

  func (t MyTrigger) Extract(tk *source.Toolkit, notifier *source.Notifier) {
    for {
      select {
      case <- notification:
        tk.Logger.Info("A notification arrived")
        notifier.Payload <- &source.Payload{}
      case <-notifier.IsShuttingDown:
        tk.Logger.Warn("gateway/enterprise: Instance is shutting down, what to do?")
        time.Sleep(5 * time.Second)
        notifier.Done <- true
      }
    }
  }
*/
type Notifier struct {

	// Payload allows you to write the payload and send it to the gateway.
	Payload chan<- *Payload

	// Error allows you to write an error (if any occurred) and send it to the
	// gateway.
	Error chan<- error

	// IsShuttingDown receives a notification when the gateway instance is shutting
	// down. This allows the function to quit (using Done) when it is ready. This
	// way, the gateway can gracefully shutdown without stopping any active work.
	IsShuttingDown <-chan bool

	// Done lets you indicate when the function is ready to gracefully exit. This
	// shall be used once IsShuttingDown is received and unblock the shutdown of
	// the gateway.
	Done chan<- bool
}
