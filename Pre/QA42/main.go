package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var left, right int
	for right = 1; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		} else {
			left++
			nums[left] = nums[right]
		}
	}
	nums = append(nums[:left])
	left++
	return left
}
