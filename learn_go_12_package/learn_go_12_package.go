package main

import (
	"fmt"
	"net/http"

	"github.com/otumian/go-ju-to/learn_go_12_package/bleed"
	"github.com/otumian/go-ju-to/learn_go_12_package/extension"
	"github.com/otumian/go-ju-to/learn_go_12_package/web"
)

func main() {

	x, y := 12, 24

	intValue := extension.Int{Value: x}
	fmt.Println(extension.Bool{Value: true}.And(intValue.IsEqualTo(4)))

	fmt.Printf("%v > %v = %v", x, y, intValue.IsGreaterThan(y))
	fmt.Println("")

	fmt.Println(intValue.IsLessThan(3))
	fmt.Println("")

	bleed.Hello()
	fmt.Println("")

	// import a net/http
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/about", web.AboutUs)
	http.HandleFunc("/dashboard", web.Dashboard)
	fmt.Println("Server running on port localhost:3000")
	http.ListenAndServe(":3000", nil)

}
