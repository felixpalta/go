package main

import (
	"fmt"
	"math"
)

// Implements error.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// Need to convert error type to basic type before using it in fmt-like functions.
	// Because they look for String() or Error() methods for printing the type.
	// Without this conversion, call to fmt.Sprint(e) would call endless recursion.
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
