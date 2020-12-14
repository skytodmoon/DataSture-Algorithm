package main

import (
	"fmt"
)

func main() {
	var candies = []int{2, 3, 5, 1, 3}
	var extraCandies = 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
	var candies1 = []int{4, 2, 1, 1, 2}
	var extraCandies1 = 1
	fmt.Println(kidsWithCandies(candies1, extraCandies1))
	var candies2 = []int{12, 1, 12}
	var extraCandies2 = 10
	fmt.Println(kidsWithCandies(candies2, extraCandies2))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var length = len(candies)
	var result = make([]bool, length)
	var candiesResult = candies
	var max = candies[0]

	for i := 0; i < length; i++ {
		result[i] = true
		if candies[i] > max {
			max = candies[i]
		}
		candiesResult[i] += extraCandies
	}
	for j := 0; j < length; j++ {
		if candiesResult[j] < max {
			result[j] = false
		}
	}
	return result

}
