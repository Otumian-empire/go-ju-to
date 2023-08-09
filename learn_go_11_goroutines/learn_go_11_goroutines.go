package main // main entry package

import (
	"fmt"
	"time"
)

// Goroutines are functions and methods that run concurrently in a Go program.
// they don't depend on the hardware and are efficient
// they don't use thread-local storage and, thus, do not have a unique ID.
// they are light-weight versions of threads that operate within the context
// of the Go runtime

/*
func myFunction(parameter) returnType {
  // function body
}

go myFunction()
A goroutine is started by invoking a previously defined function or method,
	myFunction(), with the go keyword.
*/
/*
	type User struct {
		id   int
		name string
	}

	type Users struct {
		users []User
	}

	func (obj *Users) addUser(user User) {
		fmt.Println("Adding", user.name)
		obj.users = append(obj.users, user)

	}

	func main() {

		fmt.Println("Goroutines")
		fmt.Println("")

		users := Users{}
		fmt.Println(users)

		go users.addUser(User{1, "Dan"})
		fmt.Println("Added 1")
		time.Sleep(1)

		go users.addUser(User{2, "John"})
		fmt.Println("Added 2")
		time.Sleep(1)

		go users.addUser(User{3, "Jane"})
		fmt.Println("Added 3")
		time.Sleep(1)

		fmt.Println(users)
		time.Sleep(5)
		fmt.Println(users)
	}
*/

// define a function
func print(text string) {
	fmt.Println(text)
}

func main() {
	// call goroutine
	go print("This text is from the goroutine.")

	// call function
	print("This text is from the main function.")

	//
	time.Sleep(time.Second * 1)
}
