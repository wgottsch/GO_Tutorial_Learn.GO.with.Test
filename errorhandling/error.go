package errorhandling

import (
	"github.com/pkg/errors"
)

func ShowErrors() error {
	err := d3()
	return err
}

// "Wrap" is used to wrap the underlying error, add contextual text information, and attach the call stack.
// Generally it is used to wrap calls to API from other people (standard library or third-party library).
func d1() error {
	//return errors.Wrap(ErrNoOne, "GCP service not reachable")
	//return errors.Wrap(nil, "")
	return nil
	//return errors.WithStack(ErrNoOne)
}

// "WithMessage" is used to add contextual text information to underlying error without attaching call stack.
// Apply this method for “wrapped error”only.
// Note: Do not repeat Wrap, it will record redundancy call stacks
func d2() error {
	return errors.WithMessage(d1(), string(ErrNoZwo))
}

func d3() error {
	return errors.WithMessage(d2(), "If I want tell a Story")
}
