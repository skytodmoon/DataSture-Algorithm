package main

import (
	"fmt"
)

func main() {
	var targetStr = "ab"
	var testArray = []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(targetStr, testArray))
	var targetStr1 = "abc"
	var testArray1 = []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(targetStr1, testArray1))
	var targetStr2 = "cad"
	var testArray2 = []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings(targetStr2, testArray2))
}

func countConsistentStrings(allowed string, words []string) int {
	var result = 0
	var tempHash = make(map[rune]int)

	for _, item := range allowed {
		tempHash[item] = 0
	}
	for _, item := range words {
		var check = true
		for _, item2 := range item {
			if _, ok := tempHash[item2]; !ok {
				check = false
				break
			}
		}
		if check {
			result++
		}
	}
	return result
}
