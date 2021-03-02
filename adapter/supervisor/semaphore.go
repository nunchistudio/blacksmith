package supervisor

/*
Lock holds information about a lock-key in the distributed system. This allows to
acquire and release access to resources.
*/
type Lock struct {

	// Key is the id / name / key to acquire and release.
	Key string `json:"key"`
}

/*
Semaphore holds details about a semaphore in the supervisor. This is returned by
the supervisor for a given Lock's status. These details are also included in some
of the admin API endpoints to inform clients about the semaphore of a given trigger,
action, or migration.
*/
type Semaphore struct {

	// Key is the id / name / key looked up. Even if IsApplicable is false, Key is
	// set to the appropriate key name as if it were applicable.
	Key string `json:"key"`

	// IsApplicable informs the client if a semaphore is needed for the given
	// resource. As an example, the supervisor is leveraged for CDC and CRON triggers
	// but not for HTTP or subscription ones.
	IsApplicable bool `json:"is_applicable"`

	// IsAcquired informs if the key is currently in use. It can nil if the supervisor
	// adapter encountered an error while looking up the key and therefore does not
	// know its status. It can also be nil if the key to look up does not need to
	// acquire a lock in the semaphore (in other words, if IsApplicable is false).
	IsAcquired *bool `json:"is_acquired,omitempty"`

	// AcquirerName is the name of the acquirer currently using the key. It is empty
	// if the key is not in use.
	AcquirerName string `json:"acquirer_name,omitempty"`

	// AcquirerAddress is the address of the acquirer currently using the key. It
	// is empty if the key is not in use.
	AcquirerAddress string `json:"acquirer_address,omitempty"`

	// SessionID is the ID of the session started by the running service or CLI
	// currently using the key. It is empty if the key is not in use.
	SessionID string `json:"session_id,omitempty"`
}
