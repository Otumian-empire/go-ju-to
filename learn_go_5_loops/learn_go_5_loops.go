package main // main entry package

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Loops")

	// for loop
	for i := 0; i < 6; i++ {
		fmt.Println(i, "=", math.Pow(float64(i), 2.0))
	}

	fmt.Println("")
	j := 0
	for j < 10 {
		fmt.Println(j, "=", math.Pow(float64(j), 3.0))
		j++
	}

	// for each (in range)
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// we can do break and continue
}
