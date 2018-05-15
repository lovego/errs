package errs

import (
	"bytes"
	"fmt"
	"runtime"
)

type TraceErr struct {
	err   error
	stack string
}

func Trace(err error) TraceErr {
	if trace, ok := err.(TraceErr); ok {
		return trace
	} else {
		return TraceErr{err: err, stack: Stack(3)}
	}
}

func Tracef(format string, args ...interface{}) TraceErr {
	return TraceErr{err: fmt.Errorf(format, args...), stack: Stack(3)}
}

func (s TraceErr) Stack() string {
	return s.stack
}

func (s TraceErr) Error() string {
	return s.err.Error()
}

func (s TraceErr) Err() interface{} {
	return s.err
}

func Stack(skip int) string {
	buf := new(bytes.Buffer)

	callers := make([]uintptr, 32)
	n := runtime.Callers(skip, callers)
	frames := runtime.CallersFrames(callers[:n])
	for {
		if f, ok := frames.Next(); ok {
			fmt.Fprintf(buf, "%s\n\t%s:%d (0x%x)\n", f.Function, f.File, f.Line, f.PC)
		} else {
			break
		}
	}
	return buf.String()
}

func WithStack(err error) string {
	if e, ok := err.(interface {
		Stack() string
	}); ok {
		return err.Error() + "\n" + e.Stack()
	}
	return err.Error()
}
