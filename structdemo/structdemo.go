package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	fmt.Println(Point{1, 2}) // Creating nameless object of struct Point.

	// Creating and using nameless object of nameless struct declared in place.
	fmt.Println(
		struct {
			a, b int
			c    float64
		}{3, 4, 0.1})

	// Creating named object of struct Point.
	p := Point{66, 77}
	fmt.Println(p)

	// Struct can only be initialized using 'constructor' syntax, also called 'Struct literal'.
	// In this case, there is no need to specifiy type (Point) after variable name.
	var pp = Point{88, 99}
	fmt.Println(pp)

	// Struct does not have to be forward declared, if it is defined in the global scope.
	r := Rect{100, 200}
	fmt.Println(r.width, r.height)

	// 'circle' is not a struct tag, it is a typedef. There is no struct tags in Go.
	type circle struct {
		x      int
		y      int
		radius float32
	}

	// Local struct has to be defined before it is used.
	mycircle := circle{3, 4, 5.5}
	fmt.Println(mycircle)

	// Creating named object of unnamed struct.
	position := struct {
		alpha, beta, gamma float32
	}{1.1, 2.2, 3.3}

	fmt.Println(position)

	smth()

	// Struct fields are default initialized to zero values.
	var ppp Point
	fmt.Println(ppp)
	// Same as previous.
	ppp2 := Point{}
	fmt.Println(ppp2)

	// Using designated initilizers for struct fields.
	ppp3 := Point{x: 3, y: 4}
	fmt.Println(ppp3)

	// Initializers are in wrond order, but the struct is initialized correctly.
	ppp4 := Point{y: 10, x: 12}
	fmt.Println(ppp4)

	// Uninitialized fields default to zero values.
	ppp5 := Point{y: 66}
	fmt.Println(ppp5)

	// Struct fields are accessible through pointer to struct.
	ptr := &ppp5
	ptr.x = 64
	fmt.Println(*ptr)

	// Pointers are default initialized to nil.
	var nptr *Point
	fmt.Println(nptr) // prints <nil>

	// Nameless object is not 'temporary' as in C++. Pointer to it is still valid apparently.
	temp_ptr := &Point{33, 44}
	fmt.Println(*temp_ptr)

	// new_point() returns pointer to a locally created object and it remains valid.
	fmt.Println(*new_point(66, 12))

	fmt.Println(*new_int(2))

}

func smth() {
	p := Rect{7, 8}
	fmt.Println("Rect:", p.height, p.width)

	// Struct, defined in another scope is not avaiable here.
	//mycircle := circle{1, 2, 4}
}

func new_point(xx, yy int) *Point {
	return &Point{xx, yy}
}

// Address of a local variable (of any type) can be safely returned from a function.
func new_int(x int) *int {
	ret := x + 1
	return &ret
}

type Rect struct {
	width  int
	height int
}
