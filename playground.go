package main

import (
	"fmt"
	"log"

	"github.com/MarcelGeo/goplayground/greetings"
)

// Hello returns a greeting for the named person.
func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, error := greetings.Hello("Marcel")
	messageEmpty, errorME := greetings.Hello("Fuck")

	if error != nil || errorME != nil {
		log.Fatal(error)
	}

	fmt.Println(messageEmpty)
	fmt.Println(message)
	greetings.PrintQuote()
}