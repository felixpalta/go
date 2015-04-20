package main

import (
	"fmt"
)

func main() {

	// Arrays and slices are initialized using 'constructor' syntax.
	arr := [...]int{3, 4, 5, 6}
	// For arrays size either is a compile time constant [4]int{...}, or calculated by a list of initializers, like: [...]int{...}

	fmt.Printf("arr is %T\n", arr)
	for _, val := range arr {
		fmt.Println(val)
	}

	// Out of range access causes panic: runtime error: index out of range.
	/*
		i := 4
		arr[i] = 0
		i := -1
		arr[i] = 0
	*/

	// Array of length 0 is a valid construct.
	arr2 := [0]int{}
	for _, val := range arr2 {
		fmt.Println(val) // No elements, so nothing is printed.
	}
	fmt.Printf("arr is %T, length: %v\n", arr2, len(arr2))

	// Slice can be created from an array.
	sl := arr[:2]
	print_slice(sl, "sl")

	// Cannot convert array to a slice.
	//print_slice(arr, "arr")
	// Instead need to explicitly create a slice from array.
	print_slice(arr[:], "wut")

	a := []string{"one", "two"}
	b := []string{"three", "four"}

	// b... means append all elements of slice b to a. b... is equivalent to using list of literals.
	a = append(a, b...)
	fmt.Println(a)

	zomg := make([]int, 3, 5)
	fmt.Printf("%v len: %v, cap: %v\n", fmt.Sprint(zomg), len(zomg), cap(zomg))
	zomg = zomg[:2]
	zomg[1] = 10
	fmt.Printf("%v len: %v, cap: %v\n", fmt.Sprint(zomg), len(zomg), cap(zomg))

}

func print_slice(sl []int, name string) {
	fmt.Printf("%s is %T, length: %v\n", name, sl, len(sl))
}
