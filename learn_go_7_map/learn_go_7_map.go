package main // main entry package

import (
	"fmt"
)

func main() {
	fmt.Println("Map")

	// A map is used to store a collection of unordered key-value pairs.
	// The pairs can be of the same type or of mixed types.
	// It is Go’s implementation of a hash table, which allows for
	// efficient access, insertion, and deletion.

	// An empty map can be created using the make() function and assigning it to a variable.
	newMap := make(map[string]string)
	newMap["name"] = "John doe"
	newMap["job"] = "Software developer"
	fmt.Println(newMap)
	fmt.Println("")

	// An empty map can also be created by assigning the variable to a map literal.
	someMap := map[string]string{ /* Could be empty */ "name": "Jane doe"}
	fmt.Println(someMap)
	fmt.Println("Name,", someMap["name"])
	fmt.Println("")

	// remove item
	delete(someMap, "name")
	fmt.Println(someMap)
	fmt.Println("")

	// assign a value from a map to a var
	lastName := someMap["lastName"]

	if len(lastName) > 0 {
		fmt.Println("Last name", lastName)
	} else {
		fmt.Println("someMap[\"lastName\"] doesn't exist")
	}
	fmt.Println("")

	// loop through the map
	for key, value := range newMap {
		fmt.Println(key + ": " + value)
	}
	fmt.Println("")

}
