package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

func summation(x, y int) (int, int) {
	return add(x, y), sub(x, y)
}
func getIntegerInput() int {
	var input int

	print("Enter input: ")
	// get user input as an integer
	fmt.Scan(&input)

	return input
}

func main() {
	fmt.Println("Functions")
	defer fmt.Println("We compute the sum and difference between two values")
	// above line will be executed at the end of the function

	firstInput := getIntegerInput()
	secondInput := getIntegerInput()

	fmt.Printf("The sum and difference between %v and %v is, %v and %v respectively",
		firstInput, secondInput, add(firstInput, secondInput), sub(firstInput, secondInput))
	fmt.Println("")
	fmt.Println(mult(firstInput, secondInput))

	fmt.Println("")
	fmt.Println("")
	_sum, _sub := summation(firstInput, secondInput)
	fmt.Printf("The sum and difference between %v and %v is, %v and %v respectively",
		firstInput, secondInput, _sum, _sub)
	fmt.Println("")
}
