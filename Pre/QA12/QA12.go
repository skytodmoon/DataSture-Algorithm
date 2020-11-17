package main

import (
	"fmt"
)

func main() {
	var x = "x"
	var y = "y"
	var z = "z"
	fmt.Println("move total times", hanio(2, x, y, z))
	fmt.Println("=============")
	fmt.Println("move total times", hanio(3, x, y, z))
	fmt.Println("=============")
	fmt.Println("move total times", hanio(4, x, y, z))
	fmt.Println("=============")
	fmt.Println("move total times", hanio(10, x, y, z))
	//2 N times -1
}

func hanio(n int, x, y, z string) int {
	if n < 1 {
		fmt.Println("n can not < 1")
		return 0
	}
	if n == 1 {
		fmt.Println("move ", x, "->", z)
		return 1
	} else {
		var m = 0
		m += hanio(n-1, x, z, y)
		fmt.Println("move ", x, "->", z)
		m++
		m += hanio(n-1, y, x, z)
		return m
	}
}
