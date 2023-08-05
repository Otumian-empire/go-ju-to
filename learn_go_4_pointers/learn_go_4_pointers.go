package main // main entry package

import (
	"fmt"
)

/*
addresses
pointers
dereferencing
*/
func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"

	fmt.Println("greeting is:", greeting)

	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
