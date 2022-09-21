package main

import (
	"fmt"
)

func main() {
	/*
		Belajar print hello
	*/
	fmt.Println("halo")

	//Print halo halo
	fmt.Println("Icha Ayu")
	fmt.Println("Icha", "Ayu")
	fmt.Print("Icha Ayu\n")
	fmt.Print("Icha", "Ayu\n")

	// Belajar Variabel with data type
	var namaLengkap string
	var umur int

	namaLengkap = "Icha"
	umur = 15
	fmt.Println("Hello nama saya", namaLengkap)
	fmt.Println("Umur saya", umur)

	// Belajar Variabel tanpa data type
	var namakamu = "Icha"
	var umurkamu = 30
	fmt.Printf("%T , %T\n", namakamu, umurkamu)

	// Multiple variabel declaration with data type
	var one, two, three string = "1", "2", "3"
	var angkaone, angkatwo, angkathree int = 1, 2, 3
	fmt.Println("test data multiple => ", one, two, three)
	fmt.Println("test data multiple => ", angkaone, angkatwo, angkathree)

	// Multiple variabel Declaration tanpa data type
	var one1, two2, three3 = "1", "2", "3"
	var angka1, angka2, angka3 = 1, 2, 3
	fmt.Println("test data multiple => ", one1, two2, three3)
	fmt.Println("test multiple data => ", angka1, angka2, angka3)

	// Multiple variabel Declaration short
	datasatu, datadua, datatiga := "1", "2", "3"
	datasatu1, datadua2, datatiga3 := 1, 2, 3
	fmt.Println("test data multiple => ", datasatu, datadua, datatiga)
	fmt.Println("test data multiple => ", datasatu1, datadua2, datatiga3)

	// Underscore variable dan fmt printf usage
	var cobavariabel string
	var satunama, duanama, alamatkamu, dataumur = "Icha", "Ayu", "Gresik", 15
	_, _, _, _, _ = cobavariabel, satunama, duanama, alamatkamu, dataumur
	fmt.Printf("test Underscore variabel => %T \n", satunama)
	fmt.Printf("Halo nama saya %s, umur saya %d, saya berasal dari %s\n", satunama, dataumur, alamatkamu)
}
