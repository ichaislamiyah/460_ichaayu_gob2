package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Mark", "Ariell", "Millo"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("students", student)
			result := fmt.Sprintf("Hei, my name is %s", students)
			c <- result
		}(v)
	}
	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)

}

func print(c chan string) {
	fmt.Println(<-c)
}
