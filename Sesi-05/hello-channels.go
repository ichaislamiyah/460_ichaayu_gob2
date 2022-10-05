package main

import "fmt"

func main() {
	//membuat channel
	c := make(chan string)
	//set parameter
	go introduce("Airell", c)
	go introduce("Mark", c)
	go introduce("Mallo", c)

	msg1 := <-c
	fmt.Println(msg1)
	msg2 := <-c
	fmt.Println(msg2)
	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hei, My name is %s", student)
	c <- result
}
