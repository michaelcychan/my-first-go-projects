package main

import (
	"fmt"
	"io"
	"log"

	// "os" // used only for original main
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Original main without Internet function
//
// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

// Greet function can be used with os.Stdout and http.ResponseWriter,
// because both of them are using io.Writer interface.

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
