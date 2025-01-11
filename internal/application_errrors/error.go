package application_errrors

import "errors"

var (
	ErrUnableToProcess = errors.New("unable to process, try again later")
)
