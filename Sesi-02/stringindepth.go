package main

import (
	"fmt"
)

func main() {

	//Looping over string (byte-by-byte)
	city := "Surabaya"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
	fmt.Println("==========================================")

	//Looping over string (rune-by-rune)
	str1 := "makan"
	str2 := "mangan"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))
	fmt.Println("==========================================")

	//
	str := "mangan"

	for b, s := range str {
		fmt.Printf("indez => %d, string => %s\n", b, string(s))
	}
	fmt.Println("==========================================")
}
