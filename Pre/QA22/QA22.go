package main

import (
	"fmt"
)

func main() {
	var targetNumes = []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(targetNumes))
}

func numIdenticalPairs(nums []int) int {
	var total = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i < j && nums[i] == nums[j] {
				total++
				fmt.Println(i, j)
			}
		}
	}
	return total
}
