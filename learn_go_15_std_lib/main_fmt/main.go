package main

import (
	"fmt"
)

// verbs, places holders for value format
// %v is the default format
// %T is for the type of the value
// %t is for boolean
// %d is for base 10 number
// %e is for scientific notation
// %f is for decimals with no exponents
// %W.Pf where W is the width and P the precision
// %g uses %e for large exponent else %f
// %s for string
// %p for pointers

func main() {
	fmt.Println("Hello there, fmt!!")

	/* // verbs
	name, age, gpa, hasGraduated := "John Doe", 30, 2.56, false

	// using %v
	// any type we parse as %v works (kind of like a generic)
	fmt.Printf("Name: %v, Age: %v, GPA: %v, Graduated: %v\n", name, age, gpa, hasGraduated)

	// we can get the types of the data using the %T
	fmt.Printf("Name: %v is of Type, %T\n", name, name)
	fmt.Printf("Age: %v is of Type, %T\n", age, age)
	fmt.Printf("GPA: %v is of Type, %T\n", gpa, gpa)
	fmt.Printf("Graduated: %v is of Type, %T\n", hasGraduated, hasGraduated)

	// we can also use the specifics
	// %s, %d, %f, %t for string, int, float and bool
	// when the wrong format is passed, there is a warning
	fmt.Printf("Name: %s, Age: %d, GPA: %f, Graduated: %t\n", name, age, gpa, hasGraduated)
	// NB: %v and %f have different decimal places. %v, is as passed but %f is to 6dp
	// try try the below to see what happens when we pass the incorrect format verb
	// fmt.Printf("Name: %p, Age: %d, GPA: %s, Graduated: %f\n", name, age, gpa, hasGraduated)

	// using %w.pf on a PI=22.0/7.0 to print the precision
	// 22/7 is an int (int/int = int)
	fmt.Printf("%f\n", 22.0/7) //.6 dp
	fmt.Printf("%.51f\n", 22.0/7)
	fmt.Printf("%e\n", 22.0/7)
	fmt.Printf("%g\n", 22.0/7)
	fmt.Printf("%g\n", 22.0/7)
	fmt.Printf("%p\n", &name)

	//
	// other flags
	fmt.Printf("%v %-3v\n", name, name) */

	//
	// functions
	/* // Errorf: format error
	// lets assume, we want to return some error about some invalid data
	// name is invalid, error code, n
	statusMessage, statusCode := "Please provide a valid username", 422
	errorMessage := fmt.Errorf("%s, %d", statusMessage, statusCode)
	// fmt.Println(errorMessage)
	fmt.Println(errorMessage.Error()) */

	//
	// Fprint, Fprintf, Fprintln
	/* name := "John Doe"
	fmt.Fprint(os.Stdout, "Hello there", name, "\n")
	fmt.Fprintf(os.Stdout, "Hello there %s\n", name)
	fmt.Fprintln(os.Stdout, "Hello there", name)
	// is the same as below but we have to specify where
	// the data (output) is written to

	fmt.Print("Hello there", name, "\n")
	fmt.Printf("Hello there %s\n", name)
	fmt.Println("Hello there", name) */

	//
	/* // writing into a file using Fprint...
	// filePtr, errorMessage := os.OpenFile("sample.txt", os.O_CREATE, 0664)
	// file was create we couldn't write into
	filePtr, errorMessage := os.Create("sample1.txt")
	//

	if errorMessage != nil {
		fmt.Println(errorMessage)
	} else {
		fmt.Println("File create or already exists")
		fmt.Fprintln(filePtr, "Hello there, I am a new data")
		filePtr.Close()
	} */

	//
	/* // read from file after writing
	fileContent, errorMessage := os.ReadFile("sample1.txt")

	if errorMessage != nil {
		fmt.Println(errorMessage)
	} else {
		fmt.Println(string(fileContent))
	} */

	//
	// Scan, Scanf, Scanln
	/* var name string
	var age int
	var isGopher bool

	fmt.Print("Enter you first name, age and true/false for gopher status: ")
	// n, scanError := fmt.Scan(&name, &age, &isGopher)
	// n, scanError := fmt.Scanf("%s %d %t", &name, &age, &isGopher)
	n, scanError := fmt.Scanln(&name, &age, &isGopher)

	if scanError != nil {
		fmt.Println(scanError)
	} else if n < 3 {
		fmt.Println("Enter you first name, age and true/false for gopher status")
	} else {
		fmt.Printf("Name: %s, age: %d, Is Gopher: %t\n", name, age, isGopher)
		fmt.Printf("Name: %T, age: %T, Is Gopher: %T\n", name, age, isGopher)
	} */

	/* var name string
	fmt.Scanln(&name) // don't use Scanln, it a weird one
	fmt.Println(name) */

	//
	// Sprint, Sprintf, Sprintln

	var name string
	var age int
	var isGopher bool

	fmt.Print("Enter you first name, age and true/false for gopher status: ")
	n, scanError := fmt.Scan(&name, &age, &isGopher)

	if scanError != nil {
		fmt.Println(scanError)
	} else if n < 3 {
		fmt.Println("Enter you first name, age and true/false for gopher status")
	} else {
		// sentence := fmt.Sprint("Name:", name, "Age:", age, "IsGopher:", isGopher)
		// sentence := fmt.Sprintf("Name: %s, Age: %d, IsGopher: %t", name, age, isGopher)
		sentence := fmt.Sprintln("Name:", name, "Age:", age, "IsGopher:", isGopher)

		fmt.Println(sentence)

		// we can break down the sentence
		// fmt.Sscanln(sentence, "Name:", &name, "Age:", &age, "IsGopher:", &isGopher)
		// fmt.Println(name, age, isGopher)
		// fmt.Sscanln("My name is Michael", "My name is %s", &name)
		// fmt.Println("The name in the above sentence is", name)

		// some of this things don't work I expect, I don't know much

	}

}
