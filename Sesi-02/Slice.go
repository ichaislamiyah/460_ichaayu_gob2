package main

import (
	"fmt"
	"strings"
)

func main() {

	//Slice (make function)
	var vegetables = make([]string, 3)

	_ = vegetables

	fmt.Printf("%#v", vegetables)
	fmt.Println("==========================================")

	//Slice (append function)
	var buah = make([]string, 3)

	buah = append(buah, "apel", "pisang", "mangga")
	fmt.Printf("%#v", buah)
	fmt.Println("==========================================")

	//Slice (append function with ellipsis)
	var fruit1 = []string{"pisang", "mangga", "jeruk"}
	var fruit2 = []string{"semangka", "kiwi", "salak"}

	fruit1 = append(fruit1, fruit2...)

	fmt.Printf("%#v", fruit1)
	fmt.Println("==========================================")

	//Slice (copy function)
	var buah1 = []string{"pisang", "jambu", "melon"}
	var buah2 = []string{"durian", "belimbing", "nanas"}

	nn := copy(buah1, buah2)
	fmt.Println("Buah1 =>", buah1)
	fmt.Println("Buah2 =>", buah2)
	fmt.Println("Copied elements =>", nn)
	fmt.Println("==========================================")

	//Slicing
	var sayur1 = []string{"wortel", "tomat", "gubis", "toge", "cabai"}

	var sayur2 = sayur1[1:4]
	fmt.Printf("%#v\n", sayur2)

	var sayur3 = sayur1[0:]
	fmt.Printf("%#v\n", sayur3)

	var sayur4 = sayur1[:3]
	fmt.Printf("%#v\n", sayur4)

	var sayur5 = sayur1[:]
	fmt.Printf("%#v\n", sayur5)
	fmt.Println("==========================================")

	//Slice (combining slicing and append)
	var fruitz = []string{"Jeruk", "Apel", "Mangga", "Salak", "Nanas"}
	fruitz = append(fruitz[:3], "Watermelon")
	fmt.Printf("%#v\n", fruitz)
	fmt.Println("==========================================")

	//Slice (backing array)
	var buahz = []string{"Jeruk", "Peach", "Durian", "Pisang"}

	fmt.Println("buahz cap:", cap(buahz))
	fmt.Println("buahz len:", len(buahz))

	fmt.Println(strings.Repeat("#", 20))

	var buahy = buahz[0:3]

	fmt.Println("buahy cap:", cap(buahy))
	fmt.Println("buah len:", len(buahy))

	fmt.Println(strings.Repeat("#", 20))

	var buahx = buahz[1:]

	fmt.Println("buahx cap:", cap(buahx))
	fmt.Println("buahx len:", len(buahx))
	fmt.Println("==========================================")

	//Slice (creating a new backing array)
	NCT := []string{"Mark", "Jaehyun", "Taeyeong", "Ten"}
	NCTU := []string{}

	NCTU = append(NCTU, NCT[0:2]...)

	NCT[0] = "Jaemin"
	fmt.Println("NCT: ", NCT)
	fmt.Println("NCTU: ", NCTU)
	fmt.Println("==========================================")
}
