package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	var target = nums
	var lastValue = 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			target[i] = nums[0]
		} else {
			target[i] = lastValue + nums[i]

		}
		lastValue = target[i]
	}
	return target
}



func main() {
	var nums = []int{1, 2, 3, 4}
	//var nums = []int{3, 1, 2, 10, 1}
	var target = runningSum(nums)
	fmt.Println(target)

}
