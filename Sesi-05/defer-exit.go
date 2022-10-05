package main

import (
	"fmt"
	"os"
)

func main() {
	// defer #1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")
	fmt.Println("======================================")

	// Defer #2
	callDeferFunc()
	fmt.Println("Hai everyone")
	fmt.Println("======================================")

	// Exit
	defer fmt.Println("Invoke with defer")

	fmt.Println("Before exiting")
	os.Exit(1)
}

// Func untuk defer
func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
