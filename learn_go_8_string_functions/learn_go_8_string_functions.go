package main // main entry package

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String methods")

	var fullName string = "Hello there! I am John Doe"
	fmt.Println("Full name", fullName)
	fmt.Println("")

	// we use strings.MethodName(args..)

	// strings.Contains(str, subStr)
	if strings.Contains(fullName, " ") {
		fmt.Println("There is at least one space in the name so we can split it by space")

		// strings.Count(str, subStr)
		spaceCount := strings.Count(fullName, " ")
		fmt.Println("There are", spaceCount, "spaces in:", fullName)

		// strings.Split(str, subStr)
		words := strings.Split(fullName, " ")
		fmt.Println("There are", len(words), "words in:", fullName)

		// strings.Cut(str, subStr)
		// slices a string around a separator.  import strings
		// before, after, found := strings.Cut(fullName, " ")
		/* fmt.Println(before)
		fmt.Println(after)
		fmt.Println(found) */

		before, after, found := strings.Cut(fullName, " ")

		for found {
			fmt.Println(before, ": ", after)
			before, after, found = strings.Cut(after, " ")
		}

	}

	// strings.Fields(str)-> type[] (slice)
	fmt.Println(strings.Fields(fullName))

	// strings.Index(str, subStr)
	fmt.Println(strings.Index(fullName, "John"))

	// strings.Join(array string, separator)
	fmt.Println(strings.Join([]string{"Daniel", "Pauline"}, "-%-"))

	// strings.Replace(inOriginalStr, replaceOldStr, withNewString, nNumberOfOccurrence)
	fmt.Println(strings.Replace(fullName, " ", "_", -1))

	// trimmed := strings.Trim(str, strContainingCharsToRemove) => bool

}
