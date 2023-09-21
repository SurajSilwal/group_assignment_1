package main

import (
	"fmt"
	"math"
)

func findSmallestNumber(numbers []float64) float64 {
	if len(numbers) == 0 {
		return math.Inf(1)
	}

	smallest := numbers[0]

	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}

	return smallest
}



func main() {
	
	numbers := []float64{3.14, 2.718, 1.618, 42, 0.001}

	smallest := findSmallestNumber(numbers)
	fmt.Printf("The smallest number is: %f\n", smallest)

}
