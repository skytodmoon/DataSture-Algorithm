package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b [5]int

	var i int
	//first method
	fmt.Println(len(a))

	for i = 0; i < 5; i++ {
		b[len(a)-i-1] = a[i]
	}
	fmt.Println(b)
	//second method

	var tmp = 0
	for i = 0; i < len(a)/2; i++ {
		tmp = a[i]
		a[i] = a[len(a)-i-1]
		a[len(a)-i-1] = tmp
	}
	fmt.Println(a)
}
