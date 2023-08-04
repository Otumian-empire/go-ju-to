package main // main entry package

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Arrays and slices")

	// Array
	// Arrays are numbered, fixed-length sequences of
	// elements of the same data type.
	// Arrays differ from slices in that their size
	// cannot be changed after being created.
	// Arrays are passed as value (a copy) but a slice.
	// as a reference
	var evenNumbers [5]int // array declaration

	for i := 0; i < 5; i++ {
		evenNumbers[i] = int(math.Pow(float64(i), 2))
	}

	fmt.Println(evenNumbers)

	// we can also do
	oddNumbers := [5]float32{ /* we can pass comma separated values here */ }
	for i := 0; i < 5; i++ {
		oddNumbers[i] = float32(evenNumbers[i] + 1)
	}

	fmt.Println(oddNumbers)

	// slice
	// A slice in Go is a pointer to a section of an array.
	// They are like arrays but can be dynamically sized.
	// A slice has a length and a capacity.
	// length = number of elements in the slice, len(slice_name).
	// capacity = number of elements in the underlying array from
	// the slice's start index to the end of the array, cap(slice_name)

	evenNumberSlice := evenNumbers[0:4]
	fmt.Println(evenNumberSlice)

	fmt.Println("New Slice:", evenNumberSlice, "Length:", len(evenNumberSlice), "Capacity:", cap(evenNumberSlice))

	// create an empty slice with make()
	someSlice := make([]int, 0, 5)
	fmt.Println(someSlice)
	fmt.Println("Size", len(someSlice))
	fmt.Println("Capacity", cap(someSlice))

	// update slice
	someSlice = append(someSlice, 1, 2, 3, -1)
	fmt.Println(someSlice)
	fmt.Println("Size", len(someSlice))
	fmt.Println("Capacity", cap(someSlice))
}
