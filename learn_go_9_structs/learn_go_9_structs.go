 package main // main entry package

import (
	"fmt"
)

// A struct is a user-defined type that combines several fields of different
// data types, normally related, forming a collection.

/*
	type <struct_name> struct {
		<field_one> <data-type>
			<field_two> <data-type>
		...
	}
*/

// create struct
type Human struct {
	name         string
	height       float32
	isMarried    bool
	numberOfKids int
}

// we can add a function that will behave like a construct
// so that we don't have to  be doing the field assignment all the time
func NewHuman(name string, height float32, isMarried bool, numberOfKids int) *Human {
	return &Human{name, height, isMarried, numberOfKids}
}

// we can provide default by
/*
type Human struct {
	name         string  `default: "Human"`
	height       float32 `default: 4.6`
	isMarried    bool    `default: false`
	numberOfKids int     `default: 0`
}
*/

// add methods
func (this Human) SayHello() {
	fmt.Println("Hello there my name is", this.name)
}

func (this Human) RelationshipStatus() {
	if this.isMarried {
		fmt.Println("I am married")
	} else {
		fmt.Println("Well, I am opened to all opportunities and ready to mingle")
	}
}

func main() {
	fmt.Println("Structs")
	fmt.Println("")

	// create an instance of a struct
	john := Human{
		name:         "John Doe",
		height:       5.4,
		isMarried:    true,
		numberOfKids: 3,
	}

	fmt.Println(john)
	john.SayHello()
	john.RelationshipStatus()
	fmt.Println("")

	// we can also do it this way
	john.name = "John 'The coder' Doe"
	john.height = 5.6
	john.isMarried = false
	john.numberOfKids = 6
	fmt.Println(john)
	fmt.Println("")

	// use the NewHuman
	bob := NewHuman("Daniel Paddy", 6.2, true, 4)
	fmt.Println(bob)
	fmt.Println(*bob)
	bob.name = bob.name + ", Phd"
	fmt.Println(bob)
	fmt.Println(*bob)
	fmt.Println("")

	// Another way to initialize a struct
	var nancy Human
	nancy.name = "Nancy Miles"
	nancy.numberOfKids = 1
	nancy.height = 4.9
	nancy.isMarried = true
	fmt.Println(nancy)
	fmt.Println("")

	peter := Human{
		name:         "Peter Packer",
		height:       6.6,
		isMarried:    false,
		numberOfKids: 0,
	}

	fmt.Println(peter)
	peter.SayHello()
	peter.RelationshipStatus()
	fmt.Println("")

}
