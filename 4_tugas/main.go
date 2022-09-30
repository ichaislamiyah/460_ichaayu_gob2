package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer catchErr()

	var username string
	fmt.Println("Input Username:")
	fmt.Scanln(&username)
	fmt.Println("Input Password:")
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validateAccount(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if a := recover(); a != nil {
		fmt.Println("Error Occured:", a)
	} else {
		fmt.Println("Application Running Perfectly")
	}
}

func validateAccount(username string, password []byte) (string, error) {
	p := string(password[:])
	if username == "Mark" && p != "1234" {
		return "", errors.New("Username/Password is wrong!")
	}
	return "Succesfully Logged In!", nil
}
