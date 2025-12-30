package execx

import "os"

// ErrExec reports a failure to start or an explicit execution failure.
type ErrExec struct {
	Err      error
	ExitCode int
	Signal   os.Signal
	Stderr   string
}

// Error returns the wrapped error message when available.
// @group Errors
//
// Example: error string
//
//	err := execx.ErrExec{Err: fmt.Errorf("boom")}
//	fmt.Println(err.Error())
//	// #string boom
func (e ErrExec) Error() string {
	if e.Err == nil {
		return "execx: execution failed"
	}
	return e.Err.Error()
}

// Unwrap exposes the underlying error.
// @group Errors
//
// Example: unwrap
//
//	err := execx.ErrExec{Err: fmt.Errorf("boom")}
//	fmt.Println(err.Unwrap() != nil)
//	// #bool true
func (e ErrExec) Unwrap() error {
	return e.Err
}
