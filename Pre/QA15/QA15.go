package main

import (
	"fmt"
)

func main() {
	array := []int{-1, 3, 3, 7, 10, 14, 14, 15, 21, 33, 35, 50, 51, 62, 67, 71, 75}
	target := 3
	result := firstGreaterThan(array, target)
	fmt.Println("first greater than:", result)

}

func firstGreaterThan(array []int, target int) int {
	result := -999
	low := 0
	high := len(array) - 1
	middle := 0
	for low <= high {
		middle = (low + high) / 2
		if middle < 1 {
			return result
		}
		fmt.Println("check middle:", middle, low, high)
		if array[middle] > target && array[middle-1] <= target {
			result = array[middle]
			fmt.Println(result)
			break
		} else if array[middle] < target {
			fmt.Println("up")
			low = middle + 1
		} else {
			fmt.Println("down")
			high = middle - 1
		}
	}
	return result
}
