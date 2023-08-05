package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there variables")
	fmt.Println("")

	// creating variables of different types
	var stationName string
	var nextTrainTime int8
	var isUptownTrain bool
	var babyWeight float32

	stationName = "Union Square"
	nextTrainTime = 12
	isUptownTrain = false
	babyWeight = 34.5

	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)
	fmt.Println("Baby weighs:", babyWeight)

	stationName = "Grand Central"
	nextTrainTime = 3
	isUptownTrain = true

	fmt.Println("")
	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	// operations: +, -, *, /, %
	fmt.Println("")
	fmt.Println(20 * 3)    // Prints: 60
	fmt.Println(55.21 / 5) // Prints: 11.042
	fmt.Println(9 % 2)     // Prints: 1

	// constant
	const funFact = "Hummingbirds' wings can beat up to 200 times a second."

	fmt.Println("Did you know?")
	fmt.Println(funFact)

	// reference the value
	fullName := "John Doe"
	fmt.Println("Full name", fullName)

	// printf with verbs, %v, %T, %d, %f, %.nf where n is number of decimal places
	fmt.Printf("Full name is %v, which is of type %T\n", fullName, fullName)
	fmt.Printf(
		"Current station: %v, Next train: %v, Is uptown: %v, Baby weighs: %v\n",
		stationName, nextTrainTime, isUptownTrain, babyWeight)

	// Sprintf
	sentence := fmt.Sprintf("My name is %v and I %v years old", fullName, 45)
	fmt.Println(sentence)

	// Scanf: taking inputs
	var courseName, courseCode string
	var passMark int

	fmt.Print("Enter the course name: ")
	fmt.Scan(&courseName)

	fmt.Print("Enter the course code: ")
	fmt.Scan(&courseCode)

	fmt.Print("Enter the pass mark: ")
	fmt.Scan(&passMark)

	finalStatement := fmt.Sprintf("The course, %v (%v) has a passmark of %v",
		courseName, courseCode, passMark)
	fmt.Println(finalStatement)

}
