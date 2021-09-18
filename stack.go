package errs

import (
	"bytes"
	"fmt"
	"runtime"
)

const skip = 2

type Stack struct {
	Skip int
}

func (s *Stack) IncrSkip() *Stack {
	if s == nil {
		return nil
	}
	s.Skip++
	return s
}

func (s *Stack) String() string {
	if s == nil {
		return ""
	}
	return CurrentStack(s.Skip + 1)
}

// CurrentStack return the current stack with the deepest n stack skipped.
func CurrentStack(skip int) string {
	return FullStack(skip + 3)
}

// FullStack return the full stack with the deepest n stack skipped.
func FullStack(skip int) string {
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
	if err == nil {
		return "<nil>"
	}
	if e, ok := err.(interface {
		Stack() string
	}); ok && e.Stack() != "" {
		return err.Error() + "\n" + e.Stack()
	}
	return err.Error()
}
