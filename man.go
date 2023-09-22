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
func findLargestNumber(numbers []float64) float64 {
	if len(numbers) == 0 {
		return math.Inf(-1)
	}

	largest := numbers[0]

	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	return largest
}

func calculateSum(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	return sum
}


func main() {
	
	numbers := []float64{3.14, 2.718, 1.618, 42, 0.001}

	smallest := findSmallestNumber(numbers)
	fmt.Printf("The smallest number is: %f\n", smallest)
	
	largest := findLargestNumber(numbers)
	fmt.Printf("The largest number is: %f\n", largest)
	
	
	sum := calculateSum(numbers)
	fmt.Printf("The sum of the numbers is: %f\n", sum)


}

