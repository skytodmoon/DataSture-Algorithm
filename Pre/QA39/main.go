package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(0))
}

func numberOfSteps(num int) int {
	var i int
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1

		}
		i++
	}
	return i
}
