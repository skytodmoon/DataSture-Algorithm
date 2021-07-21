package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
	fmt.Println(checkIfPangram("az"))
	fmt.Println(checkIfPangram("onrcsnlxckptsxffbyswujpamfltvmdoxovggepknmtacrjkkorjgvgtgaiaudspnpxkwikevmjeephhiyvnoymjwjfopovscbefecnoytjxfwasabwohqujwowmakpyuuqvgfab"))
}

func checkIfPangram(sentence string) bool {
	var a = [26]int{}
	for _, char := range sentence {
		a[char-'a'] = 1
	}
	for i := 0; i < 26; i++ {
		if a[i] == 0 {
			return false
		}
	}
	return true
}
