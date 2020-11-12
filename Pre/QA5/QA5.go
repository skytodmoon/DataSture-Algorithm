package main

import "fmt"

func main() {
	//demo is 20 people pop by every 5 person
	ring(20, 5, 10)

}

func ring(n, m, k int) {
	linkedList := CreateList()
	var M Method
	M = linkedList
	for i := 1; i <= n; i++ {
		M.Push(i)
	}
	for i := 1; i < k; i++ {
		element := M.Pop()
		M.Push(element.Data)
	}
	i := 1
	for !M.IsNull() {
		element := M.Pop()
		fmt.Println("start check:", element.Data)
		if i < m {
			M.Push(element.Data)
			i++
		} else {
			i = 1
			fmt.Println(element.Data)
		}
	}
}
