package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 115)
	rotate([]int{-1, -100, 3, 99}, 2)
}

//func rotate(nums []int, k int) {
//	var temp int
//	for i := 0; i < k%len(nums); i++ {
//		for j := len(nums) - 1; j >= 0; j-- {
//			if j == len(nums)-1 {
//				temp = nums[len(nums)-1]
//			}
//			if j == 0 {
//				nums[j] = temp
//			} else {
//				nums[j] = nums[j-1]
//			}
//		}
//	}
//	fmt.Println(nums)
//}

func rotate(nums []int, k int) {
	var tempNums = make([]int, len(nums))
	var length = len(nums)
	for i := 0; i < length; i++ {
		tempNums[i] = nums[i]
	}
	for i := 0; i < length; i++ {
		nums[(i+k)%len(nums)] = tempNums[i]
	}
	fmt.Println(nums)
}
