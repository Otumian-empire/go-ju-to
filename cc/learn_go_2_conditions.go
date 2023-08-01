package main

import (
	"fmt"
	"math/rand"
	// "time"
	// "math/rand"
)

func main() {
	// bool: true, false
	isBig := true
	fmt.Print("Is this big?")

	// if else statement
	if isBig {
		fmt.Print(" Yes")
	} else {
		fmt.Print(" No")
	}
	// we can have an if {} else if {} else {}

	fmt.Println("")

	// switch
	var grade string
	fmt.Print("Enter your grade: ")
	fmt.Scan(&grade)

	switch grade {
	case "A":
		fmt.Println("Your score was > 75, Pass")
	case "B":
		fmt.Println("Your score was > 65, Pass")
	case "C":
		fmt.Println("Your score was > 55, Pass")
	default:
		fmt.Println("Your score was < 55, Fail")
	}

	fmt.Println("")

	var username, password, authStatement string
	fmt.Print("Enter username and password: ")
	fmt.Scan(&username, &password)

	authStatement = fmt.Sprintf("Auth User(%v):Pass(%v)", username, password)
	fmt.Println(authStatement)

	fmt.Println("")

	// &&, ||, !
	// ==, !=
	// <, >, <=, >=
	// same as everywhere

	// random numbers
	// rand.Seed(time.Now().UnixNano())
	// we will use this to change the randomness of the seed
	fmt.Println("A simple random number is ", rand.Intn(10)) // 0 - 10

}
