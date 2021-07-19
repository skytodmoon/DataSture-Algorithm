package main

import (
	"fmt"
	"sort"
)

func main() {
	temp1 := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(temp1))
	temp2 := []int{4, 2, 5, 9, 7, 4, 8}
	fmt.Println(maxProductDifference(temp2))
	sort.Ints()
}

///转化位最大值、最小值问题
func maxProductDifference(nums []int) int {
	length := len(nums)
	quickSort(nums, 0, len(nums)-1)
	return nums[length-1]*nums[length-2] - nums[0]*nums[1]

}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
