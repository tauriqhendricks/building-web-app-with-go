package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// a function that handles requests needs theses params
// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page!")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValues adds to integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listen on port 8080
	// this returns a error
	// the _ is set the error value
	_ = http.ListenAndServe(portNumber, nil)
}
