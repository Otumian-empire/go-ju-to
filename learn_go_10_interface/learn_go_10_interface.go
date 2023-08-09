package main // main entry package

import (
	"fmt"
)

type Extension interface {
	IncreaseByOne()
	DecreaseByOne()
	Double()
}

type Int struct {
	val int
}

// Int has the implementations for the methods in the Extension
func (this *Int) IncreaseByOne() {
	fmt.Println("Increase by One")
	fmt.Println("Before:", this.val)
	this.val += 1
	fmt.Println("After:", this.val)
	fmt.Println("")
}

func (this *Int) DecreaseByOne() {
	fmt.Println("Decrease by One")
	fmt.Println("Before:", this.val)
	this.val -= 1
	fmt.Println("After:", this.val)
	fmt.Println("")
}

func (this *Int) Double() {
	fmt.Println("Double")
	fmt.Println("Before:", this.val)
	this.val *= 2
	fmt.Println("After:", this.val)
	fmt.Println("")
}

// function constructor
func NewInt(i Extension) {
	i.IncreaseByOne()
	i.DecreaseByOne()
	i.DecreaseByOne()
	i.Double()
	i.IncreaseByOne()
	i.DecreaseByOne()
}

func main() {
	fmt.Println("Interfaces")
	fmt.Println("")
	//

	NewInt(&Int{val: 10})
}

/*
type marriage interface {
	love()
	arranged()
}

type john struct {}

func (a *john) love() {
	fmt.Println("love with john")
}

func (a *john) arranged() {
	fmt.Println("arranged marriage with john")
}

func newLife(guy marriage) {
	guy.love()
	guy.arranged()
}

func main() {
	newLife(&john{})
}

*/
