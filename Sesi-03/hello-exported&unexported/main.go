package main

import (
	"helpers"
)

func main() {
	// Exported & Unexported

	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()

	// Init function
}
