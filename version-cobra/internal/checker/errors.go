package checker

import "fmt"

type UnreachableUrlError struct {
	URL string
	Err error
}

func (e *UnreachableUrlError) Error() string {
	return fmt.Sprintf("URL inaccessible : %s (%v)", e.URL, e.Err)
}

func (e *UnreachableUrlError) Unwrap() error {
	return e.Err
}
