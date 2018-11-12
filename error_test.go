package errs

import (
	"errors"
	"fmt"
	"strings"
)

func ExampleNew() {
	err := New(`no-login`, `please login first.`)
	fmt.Println("Error:", err.Error())
	fmt.Println("Code:", err.Code())
	fmt.Println("Message:", err.Message())
	fmt.Println("Stack:" + err.Stack())
	fmt.Println("Data:", err.Data())
	fmt.Println("GetError:", err.GetError())
	err.Trace()
	fmt.Println("Stack:", strings.HasPrefix(err.Stack(), "github.com/lovego/errs.ExampleNew"))
	err.SetData("data")
	fmt.Println("Data:", err.Data())
	err.SetError(errors.New("the error"))
	fmt.Println("GetError:", err.GetError())
	// Output:
	// Error: no-login: please login first.
	// Code: no-login
	// Message: please login first.
	// Stack:
	// Data: <nil>
	// GetError: <nil>
	// Stack: true
	// Data: data
	// GetError: the error
}

func ExampleTrace() {
	err := Trace(errors.New(`connection timeout`))
	fmt.Println("Error:", err.Error())
	fmt.Println("Code:" + err.Code())
	fmt.Println("Message:" + err.Message())
	fmt.Println("Stack:", strings.HasPrefix(err.Stack(), "github.com/lovego/errs.ExampleTrace"))
	fmt.Println("Data:", err.Data())
	fmt.Println("GetError:", err.GetError())
	fmt.Println("Trace Again:", Trace(err) == err)
	err.SetCodeMessage("code", "message")
	fmt.Println("Code:", err.Code())
	fmt.Println("Message:", err.Message())

	// Output:
	// Error: connection timeout
	// Code:
	// Message:
	// Stack: true
	// Data: <nil>
	// GetError: connection timeout
	// Trace Again: true
	// Code: code
	// Message: message
}

func ExampleTracef() {
	err := Tracef("connection timeout: %d", 3)
	fmt.Println("Error:", err.Error())
	fmt.Println("Code:" + err.Code())
	fmt.Println("Message:" + err.Message())
	fmt.Println("Stack:", strings.HasPrefix(err.Stack(), "github.com/lovego/errs.ExampleTrace"))
	fmt.Println("Data:", err.Data())
	fmt.Println("GetError:", err.GetError())
	fmt.Println("Trace Again:", Trace(err) == err)

	// Output:
	// Error: connection timeout: 3
	// Code:
	// Message:
	// Stack: true
	// Data: <nil>
	// GetError: connection timeout: 3
	// Trace Again: true
}
