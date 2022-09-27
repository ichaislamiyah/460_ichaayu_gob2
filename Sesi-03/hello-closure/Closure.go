package main

import (
	"fmt"
	"strings"
)

type isOddNum func(int) bool

func main() {
	//Closure (declare closure in variable)
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...))
	fmt.Println("==============================================")

	//Closure (IIFE)
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)
	fmt.Println("==============================================")

	//Closure (closure as a return value)
	var studentLists = []string{"Jaehyun", "Mark", "Seungyoon", "Senghoon"}

	var find = findStudent(studentLists)

	fmt.Println(find("Jaehyun"))
	fmt.Println("==============================================")

	//Closure (callback)
	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}

	var find2 = findOddNumbers(numbers2, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd number:", find2)
	fmt.Println("==============================================")

	// Closure (callback) lagi
	var numbers3 = []int{2, 5, 8, 10, 3, 99, 23}

	var find3 = findOddNumbers2(numbers3, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd number:", find3)
	fmt.Println("==============================================")

}

// func findStudent
func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

// func findOddNumbers
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

// func findOddNumbers
func findOddNumbers2(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
