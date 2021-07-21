package main

import (
	"fmt"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
	nums2 := []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent(nums2))
	nums3 := []int{7, 7, 7, 7}
	fmt.Println(smallerNumbersThanCurrent(nums3))
}

///参考冒泡排序的原理，每个位置只比较一次，哪个位置大，哪个位置++；采用双循环，时间复杂度略差，但是比较好理解。
func smallerNumbersThanCurrent(nums []int) []int {
	resultNums := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				resultNums[i] += 1
			} else if nums[i] < nums[j] {
				resultNums[j] += 1
			}
		}
	}
	return resultNums
}
