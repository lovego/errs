package errs

import (
	"fmt"
)

type Error struct {
	err           error // original error
	code, message string
	stack         string
	data          interface{}
}

func New(code, message string) *Error {
	return &Error{code: code, message: message}
}

func Newf(code, message string, args ...interface{}) *Error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	return &Error{code: code, message: message}
}

// Wrap return an wrapped error.
// Choose "error" return type instead of "*Error", to avoid "nil pointer" become a "non nil error".
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &Error{err: err}
}

// Trace return an traced error with stack.
// Choose "error" return type instead of "*Error", to avoid "nil pointer" become a "non nil error".
func Trace(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		if e.Stack() == "" {
			e.stack = CurrentStack(1)
		}
		return e
	} else {
		return &Error{err: err, stack: CurrentStack(1)}
	}
}

func Tracef(format string, args ...interface{}) *Error {
	return &Error{err: fmt.Errorf(format, args...), stack: CurrentStack(1)}
}

// Verbose print err verbosely
func Verbose(err error) string {
	if e, ok := err.(*Error); ok {
		return fmt.Sprintf("%s\n%s\n%v", e.Error(), e.Stack(), e.Data())
	} else {
		return err.Error()
	}
}

func (err *Error) Error() string {
	if err.err != nil {
		return err.err.Error()
	} else {
		return err.code + `: ` + err.message
	}
}

func (err *Error) Stack() string {
	return err.stack
}

func (err *Error) Trace() *Error {
	err.stack = CurrentStack(1)
	return err
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) Message() string {
	return err.message
}

func (err *Error) SetCodeMessage(code, message string) *Error {
	err.code, err.message = code, message
	return err
}

func (err *Error) Data() interface{} {
	return err.data
}

func (err *Error) SetData(data interface{}) *Error {
	err.data = data
	return err
}

func (err *Error) GetError() error {
	return err.err
}

func (err *Error) SetError(e error) *Error {
	err.err = e
	return err
}
