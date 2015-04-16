package main

import (
	"fmt"
)

func fact(n uint) uint {
	res := uint(1)

	for i := uint(2); i <= n; i++ {
		res = res * i
	}
	return res
}

func main() {

	for i := uint(0); i < 10; i++ {

		fmt.Printf("%v! = %v\n", i, fact(i))
	}
}
