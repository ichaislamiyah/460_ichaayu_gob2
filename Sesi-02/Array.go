package main

import (
	"fmt"
)

func main() {
	//Array
	var numbers [5]int
	numbers = [5]int{1, 2, 3, 4, 5}

	var strings = [4]string{"Jennie", "Jisoo", "Lisa", "Rose"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
	fmt.Println("====================================")

	//Array (Modify Element Through Index)
	var fruits = [3]string{"Semangka", "Anggur", "Pisang"}
	fruits[0] = "Watermelon"
	fruits[1] = "Grapes"
	fruits[2] = "Banana"

	fmt.Printf("%#v\n", fruits)
	fmt.Println("====================================")

	//Array (Loop through elements)
	var buah = [2]string{"apple", "mango"}
	for i := 0; i < len(buah); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, buah[i])
		fmt.Println("====================================")
	}

	//Array (Multidimensional Array)
	balances := [2][3]int{{4, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
