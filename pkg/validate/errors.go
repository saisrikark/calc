package validate

import "errors"

var (
	insufficientArgsErr = errors.New("Insufficient number of arguments")
	excessiveArgsErr    = errors.New("Excessive number of arguments")
)
