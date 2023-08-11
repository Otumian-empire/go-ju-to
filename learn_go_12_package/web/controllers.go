package web

import (
	"fmt"
	"net/http"

	"rsc.io/quote/v4"
	"rsc.io/sampler"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Index page, %v</h1>", sampler.Glass())
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About us page, %v</h1>", quote.Hello())
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%v</p>", quote.Opt())
}
