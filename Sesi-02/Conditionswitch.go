package main

import (
	"fmt"
)

func main() {

	//Conditions
	var currentyear = 2021

	if age := currentyear - 1998; age < 17 {
		fmt.Printf("Kamu belum boleh membuat kartu sim => %d", age)
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu SIM")
		fmt.Println("====================================")
	}

	// Switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
		fmt.Println("====================================")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	//Switch with Relational Operators
	var nilai = 6

	switch {
	case nilai == 8:
		fmt.Println("Perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("not bad")
		fmt.Println("====================================")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	//Switch (Fallthrough keyword)
	var nilai1 = 6

	switch {
	case nilai1 == 8:
		fmt.Println("Sempurna")
	case (nilai1 < 8) && (nilai1 > 3):
		fmt.Println("Tidak Buruk")
		fmt.Println("====================================")
		fallthrough
	case nilai1 < 5:
		fmt.Println("Gapapa, Belajar lagi yuk!")
		fmt.Println("====================================")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	//Nested Conditions
	var score1 = 5

	if score1 > 6 {
		switch score1 {
		case 10:
			fmt.Println("Syempurna")
		default:
			fmt.Println("Bagoezz")
		}
	} else {
		if score1 == 5 {
			fmt.Println("Tidak boeroek")
			fmt.Println("====================================")
		} else if score1 == 3 {
			fmt.Println("Cobhaa laghie ea")
		} else {
			fmt.Println("Kamoe pasthie bizzha")
			if score1 == 0 {
				fmt.Println("Cobha laghie dengan keraz")
			}
		}
	}
}
