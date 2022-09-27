package main

import (
	"fmt"
	"sync"
)

// Waitgroup (Implementation)
func main() {
	buah := []string{"Apel", "Peach", "Semangka", "Pisang", "Jambu"}

	var wg sync.WaitGroup

	for index, buah := range buah {
		wg.Add(1)
		go printBuah(index, buah, &wg)
	}

	wg.Wait()
}

func printBuah(index int, buah string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, buah => %s\n", index, buah)
	wg.Done()
}
