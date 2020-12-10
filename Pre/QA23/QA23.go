package main

import "fmt"

func main() {
	// var array = []int{2, 5, 1, 3, 4, 7}
	// var n = 3
	var array = []int{1, 2, 3, 4, 4, 3, 2, 1}
	var n = 4
	fmt.Println(shuffle(array, n))
}

func shuffle(nums []int, n int) []int {
	var targetArray = make([]int, 2*n)
	for i := 0; i < n; i++ {
		targetArray[2*i] = nums[:n][i]
		targetArray[2*i+1] = nums[n:][i]
		//targetArray = append(targetArray, nums[:n][i], nums[n:][i])

	}
	return targetArray
}
