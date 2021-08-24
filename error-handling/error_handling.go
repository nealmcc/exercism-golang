// Package erratum solves the 'error-handling' exercism problem.
package erratum

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Use performs error handling on a fictitious resource opener, which can
// return transient errors when attempting to open a resource.  It also
// handles panics which may occur when the resource itself is used.
func Use(o ResourceOpener, input string) error {
	r, err := open(o, 5, time.Duration(100*time.Millisecond))
	if err != nil {
		return err
	}
	defer r.Close()

	return nil
}

func open(o ResourceOpener, retry int, wait time.Duration) (Resource, error) {
	var (
		r   Resource
		err error
	)

	for n := 0; n < retry; n, wait = n+1, wait*2 {
		r, err = o()
		if err == nil {
			break
		}

		if isTransient(err) {
			log.Printf("transient error %s ; waiting %v\n", err, wait)
			time.Sleep(wait)
			continue
		}

		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("retried opening resource %d times; last error: %w",
			retry, err)
	}
	return r, nil
}

// isTransient returns true if err, or any error it wraps is transient.
// An error is considered to be transient if it has a method IsTransient()
// that returns true.
func isTransient(err error) bool {
	var t interface{ IsTransient() bool }
	if errors.As(err, &t) {
		return t.IsTransient()
	}
	return false
}
