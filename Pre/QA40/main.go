package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))
	fmt.Println(maxDepth(")("))
	fmt.Println(maxDepth("(()"))
	fmt.Println(maxDepth(""))
	fmt.Println(maxDepth("()()"))
	fmt.Println(maxDepth("()(()())"))
}

func maxDepth(s string) int {
	var depth int
	var index int
	for _, v := range s {
		if v == '(' {
			index++
			if index > depth {
				depth = index
			}
		}
		if v == ')' && index > 0 {
			index--
		}

	}
	if index != 0 {
		return 0
	}
	return depth
}
