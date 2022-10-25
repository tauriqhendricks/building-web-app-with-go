package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/" pathname
	// listening for a request by a web browser
	// we didnt start any process to say "listen for a requesr"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}

		// Sprintf allows you to return any data type as a string
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// listen on port 8080
	// this return a error
	// the _ is set the error value
	_ = http.ListenAndServe(":8000", nil)
}
