package errs

import (
	"errors"
	"fmt"
	"regexp"
)

func ExampleStack() {
	stack := Stack(0)
	re := regexp.MustCompile(`^runtime\.Callers
	/.+/src/runtime/extern\.go:2(12|24) \(0x[a-z0-9]+\)
github.com/lovego/errs.Stack
	/.+/src/github\.com/lovego/errs/(stack\.go:13|_test/_obj_test/stack\.go:14) \(0x[a-z0-9]+\)
github.com/lovego/errs.ExampleStack
	/.+/src/github.com/lovego/errs/stack_test\.go:10 \(0x[a-z0-9]+\)
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
	/.+/src/github.com/lovego/errs/stack_test\.go:27 \(0x[a-z0-9]+\)
`)
	if !re.MatchString(stack) {
		fmt.Println(stack)
	}
	// Output: the error
}