package draining

import (
	"github.com/giantswarm/microerror"
)

var podNotRunningError = microerror.New("pod not running")

// IsPodNotRunning asserts podNotRunningError.
func IsPodNotRunning(err error) bool {
	return microerror.Cause(err) == podNotRunningError
}

var waitTimeoutError = microerror.New("waitTimeout")

// IsWaitTimeout asserts invalidConfigError.
func IsWaitTimeout(err error) bool {
	return microerror.Cause(err) == waitTimeoutError
}
