package main

import (
	"fmt"
)

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
	fmt.Println(xorOperation(1, 7))
	fmt.Println(xorOperation(10, 5))
}

func xorOperation(n int, start int) int {
	var result = 0
	for i := 0; i < n; i++ {
		result = (start + 2*i) ^ result
	}
	return result
}

func xorOperation2(n int, start int) int {
	if n == 1 {
		return start
	}
	return xorOperation(n-1, start) ^ (start + 2*(n-1))
}
