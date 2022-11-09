package main

import (
	"fmt"
	"net/http"

	"github.com/tauriqhendricks/building-web-app-with-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listen on port 8080
	// this returns a error
	// the _ is set the error value
	_ = http.ListenAndServe(portNumber, nil)
}
