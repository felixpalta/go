package main

import (
	"fmt"
)

func main() {
	// Deferred function is executed after the surrounding function returns.
	defer fmt.Println("world!")

	fmt.Println("Hellou")

	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done.")
}
