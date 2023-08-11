// https://dev.to/dsysd_dev/demystifying-error-handling-in-go-best-practices-and-beyond-272f
package main

import (
	"fmt"
)

// explicitly handling the error
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Can not divide by zero")
	}

	return x / y, nil
}

// propagating the error
func wrapDiv(x, y int) error {
	if res, err := div(x, y); err != nil {
		return err
	} else {
		fmt.Printf("%v / %v = %v\n", x, y, res)
	}

	return nil
}

func recoverDiv() {
	if err := recover(); err != nil {
		fmt.Println("Recover from an error")
		fmt.Println(err)
	} else {
		fmt.Println("There is no issue")
	}
}

func directDiv(x, y int) int {
	defer recoverDiv()
	return x / y
}

func main() {
	fmt.Println("Error Handling")

	// In Go, error handling is approached differently
	// emphasizing explicit error handling and avoiding exceptions

	// 1. The Basics of Error Handling in Go
	// NB: you can check out the std lib
	// https://pkg.go.dev/errors@go1.20.7

	// By convention, functions that can return an error have a (result, error)
	// return signature, where the error value is the last returned value
	// https://github.com/Otumian-empire/go-simple-calculator/blob/main/main.go#L64
	// in this project error wasn't handle in the sense of error handling
	// rather we checked if the second operand is zero
	// x, y := 2, 0
	// res, err := x / y
	// assignment mismatch: 2 variables but 1 value
	// it will be better to have something like this in a function and then handle it
	// we have to be aware of instances where an error could occur and handle it

	// when using divide, we have to check that error is nil for there to be a result
	/* x, y := 2, 0
	result, err := div(2, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v / %v = %v", x, y, result)
	}
	fmt.Println("") */

	// 2. Error Propagation and Centralized Handling

	// 3. Error Wrapping and Context
	// By wrapping an error with additional context,
	// you can provide more detailed information about
	// the error’s origin and the surrounding circumstances.
	/* if err := wrapDiv(10, 3); err != nil {
		fmt.Println(err)
	} */

	// 4. Error Handling Libraries and Tools

	// 5. Handling Panics and Recovering
	fmt.Println(directDiv(2, 0))
	fmt.Println(directDiv(5, 3))

}
