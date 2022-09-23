package main

import (
	"fmt"
)

func main() {
	//Loopings (first way of loopings)
	for i := 0; i < 4; i++ {
		fmt.Println("Angka", i)
	}

	//Loopings (second way of loopings)
	var b = 0

	for b < 3 {
		fmt.Println("Angka", b)
		b++
	}

	//Loopings (third way of loopings)
	var a = 0
	for {
		fmt.Println("Angka", a)

		a++
		if a == 3 {
			break
		}
	}

	//Loopings (Break and Continue Keywords)
	for c := 1; c <= 10; c++ {
		if c%2 == 1 {
			continue
		}
		if c > 7 {
			break
		}
		fmt.Println("Angka", c)
	}

	//Loopings (Nested Looping)
	for d := 0; d < 5; d++ {
		for e := d; e < 5; e++ {
			fmt.Println(e, " ")
		}

		fmt.Println()
	}

	//Loopings (Label)
outerLoop:
	for f := 0; f < 3; f++ {
		fmt.Println("Perulangan ke - ", f+1)
		for g := 0; g < 3; g++ {
			if f == 2 {
				break outerLoop
			}
			fmt.Print(g, " ")
		}
		fmt.Print("\n")
	}
}
