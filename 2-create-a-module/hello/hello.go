package main

import (
	"fmt"
	"log"

	"github.com/gaurav-iitg/create-a-module/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gaurav", "Rahul", "Paras"}

	messages, error := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if error != nil {
		log.Fatal(error)
		return
	}
	// If no error was returned, print the returned message
	fmt.Print(messages)
}
