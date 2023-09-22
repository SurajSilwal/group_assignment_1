package main

import (
	"fmt"
	"math"
)

func calculateSum(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
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

func calculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := calculateSum(numbers)
	average := sum / float64(len(numbers))
	return average
}

func countEvenNumbers(numbers []float64) int {
	count := 0

	for _, num := range numbers {
		if math.Mod(num, 2) == 0 {
			count++
		}
	}

	return count
}

func findMedian(numbers []float64) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}

	sort.Float64s(numbers)
	middle := len(numbers) / 2

	if len(numbers)%2 == 0 {
		// Average of two middle numbers if the slice has an even length
		return (numbers[middle-1] + numbers[middle]) / 2.0
	} else {
		// Middle number if the slice has an odd length
		return numbers[middle]
	}
}



func main() {

	numbers := []float64{3.14, 2.718, 1.618, 42, 0.001}

	smallest := findSmallestNumber(numbers)
	fmt.Printf("The smallest number is: %f\n", smallest)

	sum := calculateSum(numbers)
	fmt.Printf("The sum of the numbers is: %f\n", sum)

	average := calculateAverage(numbers)
	fmt.Printf("The average of the numbers is: %f\n", average)

	evenCount := countEvenNumbers(numbers)
	fmt.Printf("Number of even numbers: %d\n", evenCount)
	
	median := findMedian(numbers)
	fmt.Printf("The median of the numbers is: %f\n", median)

}
