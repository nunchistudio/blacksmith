package mutex

import (
	"time"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
Mutex provides features for remote mutual exclusion locks. This can be used inside
wanderer adapters to ensure lock access when running database migrations.
*/
type Mutex interface {

	// Lock locks the mutex. If the lock is already in use, this shall return a nil
	// lock and an error.
	Lock(string) (*Lock, error)

	// Unlock unlocks the mutex. If the lock is already in use, this shall return a
	// nil lock and an error.
	Unlock(string) (*Lock, error)

	// Status returns the current status of the mutex. If the mutex is inlocked,
	// the lock will be nil.
	Status() (*Lock, error)
}

/*
Lock represents an exclusion lock, used by the mutex.
*/
type Lock struct {

	// ID is the unique identifier of the lock. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// AcquiredAt is a timestamp of the lock acquisition date.
	AcquiredAt time.Time `json:"acquired_at"`

	// ReleasedAt is a timestamp of the lock release date.
	ReleasedAt *time.Time `json:"released_at"`

	// CreatedAt is a timestamp of the lock creation date into the datastore.
	CreatedAt time.Time `json:"created_at"`
}

/*
Validate validates the lock info. It returns an error if the lock is not valid.
*/
func (lock *Lock) Validate() error {
	fail := &errors.Error{
		Message:     "mutex: Failed to validate lock",
		Validations: []errors.Validation{},
	}

	if lock == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Lock must not be nil",
			Path:    []string{"lock"},
		})

		return fail
	}

	if lock.ID == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "ID must not be empty",
			Path:    []string{"lock", "ID"},
		})
	}

	if lock.AcquiredAt.IsZero() {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "AcquiredAt must not be empty",
			Path:    []string{"lock", "AcquiredAt"},
		})
	}

	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
