package main

import (
	"fmt"
	"strings"
)

func main() {
	//Pointer
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber

	//Pointer (memory address)
	var nomorSatu int = 4
	var nomorDua *int = &nomorSatu
	fmt.Println("Nilai nomorSatu : ", &nomorSatu)
	fmt.Println("Alamat memori nomorSatu : ", &nomorSatu)

	fmt.Println("Nilai nomorDua : ", &nomorDua)
	fmt.Println("Alamat memori nomorSatu : ", &nomorDua)
	fmt.Println("==========================================")

	//Pointer (changing value through pointer)
	var firstOrang string = "Jaehyun"
	var secondOrang *string = &firstOrang

	fmt.Println("firstOrang (value) : ", firstOrang)
	fmt.Println("firstOrang (memory address) : ", &firstOrang)

	fmt.Println("secondOrang (value) : ", *&secondOrang)
	fmt.Println("secondOrang (memory address) : ", secondOrang)

	fmt.Println("\n", strings.Repeat("#", 50), "\n")

	*secondOrang = "Doe"

	fmt.Println("firstOrang (value) : ", firstOrang)
	fmt.Println("firstOrang (memory address) : ", &firstOrang)

	fmt.Println("secondOrang (value) : ", *&secondOrang)
	fmt.Println("secondOrang (memory address) : ", secondOrang)
	fmt.Println("==========================================")

	//Pointer (pointer as a parameter)
	var k int = 5
	fmt.Println("Before:", k)

	changeValue(&k)

	fmt.Print("After:", k)
}
func changeValue(number *int) {
	*number = 10
}
