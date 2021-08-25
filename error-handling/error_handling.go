// Package erratum solves the 'error-handling' exercism problem.
package erratum

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Use performs error handling on a fictitious resource opener, which can
// return transient errors when attempting to open a resource.  It also
// handles panics which may occur when the resource itself is used.
func Use(o ResourceOpener, input string) (err error) {
	rnd := rand.New(rand.NewSource(42))
	maxTries := 5
	wait := time.Duration(50+rnd.Intn(50)) * time.Millisecond
	r, err := open(o, maxTries, wait)
	if err != nil {
		return err
	}

	defer r.Close()
	defer func() {
		p := recover()
		if p == nil {
			return
		}
		switch p := p.(type) {
		case FrobError:
			r.Defrob(p.defrobTag)
			err = p.inner
		case error:
			err = p
		default:
			panic(p)
		}
	}()

	r.Frob(input)
	return nil
}

// open attempts to obtain a resource from o.  It will try up to maxTries times,
// and sleep between attempts.  The first delay will be given by wait, and
// then will double with each failed attempt.
// If after maxTries attempts it still cannot obtain a Resource from o, then
// this function returns the most recent error from o.
// It is recommended to use a variable delay across consumers of o, so that
// retry attempts are staggered.
func open(o ResourceOpener, maxTries int, wait time.Duration) (r Resource, err error) {
	for n := 0; n < maxTries; n, wait = n+1, wait*2 {
		r, err = o()
		if err == nil {
			return
		}

		if !isTransient(err) {
			return
		}

		log.Printf("transient error %s ; retrying in %v\n", err, wait)
		time.Sleep(wait)
	}
	return nil, fmt.Errorf("tried opening resource %d times; last error: %w",
		maxTries, err)
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
