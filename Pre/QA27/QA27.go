package main

import "fmt"

func main() {
	var testArray = []int{3, 6, -2, -5, 7, 3}
	fmt.Println(adjacentElementsProductGTY(testArray))
}

func adjacentElementsProductGTY(input_one []int) int {
	var result = -10000
	for i := 0; i < len(input_one)-1; i++ {
		target := input_one[i] * input_one[i+1]
		if result < target {
			result = target
		}
	}
	return result
}
