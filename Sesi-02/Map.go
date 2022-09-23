package main

import (
	"fmt"
)

func main() {
	//MAP
	var orang = map[string]string{
		"nama":   "Mark",
		"umur":   "24",
		"alamat": "Canada",
	}

	fmt.Println("nama:", orang["nama"])
	fmt.Println("umur:", orang["umur"])
	fmt.Println("alamat:", orang["alamat"])
	fmt.Println("==========================================")

	//Looping with MAP
	var person = map[string]string{
		"nama":   "Mark",
		"umur":   "24",
		"alamat": "Canada",
	}
	for key, value := range person {
		fmt.Println(key, ":", value)
	}
	fmt.Println("==========================================")

	//Deleting Value
	var person1 = map[string]string{
		"nama":   "Jaehyun",
		"umur":   "26",
		"alamat": "Seoul",
	}

	value, exist := person1["umur"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}
	delete(person1, "umur")

	value, exist = person1["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}
	fmt.Println("==========================================")

	//Combining slice with MAP
	var orang1 = []map[string]string{
		{"nama": "Seungyoon", "age": "29"},
		{"nama": "Seunghoon", "age": "31"},
		{"nama": "Seungmoon", "age": "35"},
	}

	for i, orang2 := range orang1 {
		fmt.Printf("Index: %d, nama: %s, umur: %s\n", i, orang2["nama"], orang2["age"])
	}
	fmt.Println("==========================================")
}
