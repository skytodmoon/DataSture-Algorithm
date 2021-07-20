package main

import (
	"fmt"
)

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}

func decompressRLElist(nums []int) []int {
	//初始化容量
	var total int
	for i := 0; i < len(nums)/2; i++ {
		total += nums[2*i]
	}
	var finalNums = make([]int, total, total)

	for i := 0; i < len(nums)/2; i++ {
		for j := 0; j < nums[2*i]; j++ {
			finalNums = append(finalNums, nums[2*i+1])
		}
	}
	return finalNums
}

//func decompressRLElist(nums []int) []int {
//	//初始化容量
//	var tempNumsStr string
//	k := 0
//	for i := 0; i < len(nums)/2; i++ {
//		for j := 0; j < nums[2*i]; j++ {
//			tempNumsStr += strconv.Itoa(nums[2*i+1]) + ","
//			k++
//		}
//	}
//	var finalNums = make([]int, k)
//	var tempNums2 = strings.Split(tempNumsStr, ",")
//	for index, s := range tempNums2 {
//		if s == "" {
//			break
//		}
//		finalNums[index], _ = strconv.Atoi(s)
//	}
//	return finalNums
//}
