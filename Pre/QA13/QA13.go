package main

import "fmt"

func main() {
	fmt.Println(fobo(1))
	fmt.Println(fobo(2))
	fmt.Println(fobo(3))
	fmt.Println(fobo(4))
	fmt.Println(fobo(5))
	fmt.Println(fobo(6))
	fmt.Println(fobo(6))
	fmt.Println(fobo(7))
	fmt.Println(fobo(8))
	fmt.Println(fobo(9))
	fmt.Println("===========")

}

func fobo(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fobo(n-1) + fobo(n-2)
}
