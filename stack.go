package errs

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

var panicStackDepth int

func init() {
	defer setupPanicStackDepth()
	panic(nil)
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

func PanicStackDepth() int {
	return panicStackDepth
}

func setupPanicStackDepth() {
	recover()
	callers := make([]uintptr, 32)
	n := runtime.Callers(2, callers)
	frames := runtime.CallersFrames(callers[:n])

	depth := 0
	for {
		if f, ok := frames.Next(); ok && strings.HasPrefix(f.Function, "runtime.") {
			depth++
		} else {
			break
		}
	}
	panicStackDepth = depth
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
