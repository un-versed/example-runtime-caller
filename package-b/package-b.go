package b

import (
	customerrors "github.com/un-versed/example-runtime-caller/errors"
)

func FromB(shouldError bool) error {
	if shouldError {
		return customerrors.New("error from package-b", "ERROR-0000001")
	}
	return nil
}
