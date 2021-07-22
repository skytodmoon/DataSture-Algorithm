package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
}

func createTargetArray(nums []int, index []int) []int {
	for i := 0; i < len(nums); i++ {
		if index[i] == i {
			continue
		} else if index[i] < i {
			for j := i; j > index[i]; j-- {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		} else {
			for j := i; j < index[i]; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
