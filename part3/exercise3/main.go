package main

import (
	"fmt"
)

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}

func main() {
	c1 := customError{
		info: "This is an error message",
	}

	foo(c1)
}

/*
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint,
*/
