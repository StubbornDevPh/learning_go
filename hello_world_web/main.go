package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "Hello Gab")
	default:
		fmt.Fprint(w, "big fat err")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello break it down yo!</h1>")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.HandleFunc("/breakitdown", htmlVsPlain)
	http.ListenAndServe("", nil)
}
