package main

import (
	"fmt"
)

func main() {
	inputArray := []int{2, 1, 4, 5, 3}

	min, max, minIndex, maxIndex, avg, total := 99, 0, 0, 0, 0, 0

	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] < min {
			min = inputArray[i]
			minIndex = i
		}
	}
	inputArray = append(inputArray[:minIndex], inputArray[minIndex+1:]...)
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] > max {
			max = inputArray[i]
			maxIndex = i
		}
	}
	inputArray = append(inputArray[:maxIndex], inputArray[maxIndex+1:]...)

	fmt.Println(inputArray)

	for i := 0; i < len(inputArray); i++ {
		total += inputArray[i]
	}

	avg = total / len(inputArray)
	fmt.Println("avg:", avg)

}
