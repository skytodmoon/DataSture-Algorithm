package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
	fmt.Println(subtractProductAndSum(705))
}

func subtractProductAndSum(n int) int {
	var total int
	var cross = 1
	for n != 0 {
		temp := n % 10
		total += temp
		cross *= temp
		n /= 10
	}
	return cross - total
}
