package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()

	// Demo of switch construct without condition. It's equivalent to switch (true).
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon")
		// Breaks are placed automatically.
	default:
		fmt.Println("Good evening!")
	}

	fmt.Println("Time is: ", hour, "hours")
}
