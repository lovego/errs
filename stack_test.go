package errs

import (
	"errors"
	"fmt"
	"regexp"
)

func ExampleFullStack() {
	stack := FullStack(0)
	re := regexp.MustCompile(`^runtime\.Callers
	/.+/src/runtime/extern\.go:(\d+) \(0x[a-f0-9]+\)
github.com/lovego/errs.FullStack
	.*/stack\.go:(\d+) \(0x[a-f0-9]+\)
github.com/lovego/errs.ExampleFullStack
	/.+/errs/stack_test\.go:10 \(0x[a-f0-9]+\)
`)
	if !re.MatchString(stack) {
		fmt.Println(stack)
	}
	// Output:
}

func ExampleCurrentStack() {
	stack := CurrentStack(0)
	re := regexp.MustCompile(`^github.com/lovego/errs.ExampleCurrentStack
	/.+/errs/stack_test\.go:25 \(0x[a-f0-9]+\)
`)
	if !re.MatchString(stack) {
		fmt.Println(stack)
	}
	// Output:
}

func ExampleWithStack() {
	err := errors.New("the error")
	fmt.Println(WithStack(err))
	stack := WithStack(Trace(err))
	re := regexp.MustCompile(`the error
github.com/lovego/errs.ExampleWithStack
	/.+/errs/stack_test\.go:38 \(0x[a-z0-9]+\)
`)
	if !re.MatchString(stack) {
		fmt.Println(stack)
	}
	// Output: the error
}
