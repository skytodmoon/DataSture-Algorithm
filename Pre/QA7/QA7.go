package main

import (
	"fmt"
)

func main() {
	inputArray := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < len(inputArray); i++ {
		for j := 0; j < len(inputArray); j++ {
			fmt.Println(i, j)
			if i == j {
				continue
			}
			if inputArray[i] == inputArray[j] {
				fmt.Println("equal:", i, j, inputArray[i])
				inputArray = append(inputArray[:j], inputArray[j+1:]...)
			}
		}

	}
	fmt.Println(inputArray)
}
