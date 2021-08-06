package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2}))
}

//分析：只有低买高卖才会获得利润，买卖过程其实并不重要，只要保证上次卖出，这次就可以考虑买入
//买卖看作一个整体，买卖的发生条件是i<i+1 i+1>i+2
func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			if i+2 >= len(prices) {
				profit += prices[i+1] - prices[i]
				break
			}
			if prices[i+1] > prices[i+2] {
				profit += prices[i+1] - prices[i]
			} else {
				profit += prices[i+2] - prices[i]
			}
			i++
		}
	}
	return profit
}
