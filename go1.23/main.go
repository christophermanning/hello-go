package main

import (
	"fmt"
	"slices"
)

func main() {
	// the new `slices.Sort` function in 1.23 can sort any `Ordered` type: https://pkg.go.dev/cmp#Ordered

	// strings
	hello_world := []string{"world", "hello"}
	slices.Sort(hello_world)
	fmt.Println(hello_world)

	// ints
	numbers := []int{3, 1, 2}
	slices.Sort(numbers)
	fmt.Println(numbers)

	// floats
	floats := []float64{3.45, 1.23, 2.34}
	slices.Sort(floats)
	fmt.Println(floats)

	// in previous versions, a type specific sort function was needed
	// e.g. `sort.Strings()`, `sort.Ints()`, ...
}
