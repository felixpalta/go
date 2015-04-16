package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Print("Saturday is ")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("today!")
	case today + 1:
		fmt.Println("tomorrow!")
	case today + 2:
		fmt.Println("the day after tomorrow!")
	default:
		fmt.Printf("in %v days", int(time.Saturday-today))
	}

}
