package main

import (
	"fmt"
)

func main() {
	var s = "lrloseumgh"
	var k = 6
	fmt.Println(reverseLeftWords(s, k))

}

func reverseLeftWords(s string, n int) string {

	var strArray = []byte(s)
	var target = string(append(strArray[n:], strArray[0:n]...))
	return target
}
