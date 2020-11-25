package main

<<<<<<< HEAD
func main() {
	var str1 = "anagramc"
	var str2 = "nagcaram"
	var check = isAnagram(str1, str2)
	println("check is anagram:", check)
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var c1, c2 [26]int
	for _, ch := range a {
		c1[ch-'a']++
	}
	for _, ch := range b {
		c2[ch-'a']++
	}
	return c1 == c2
=======
import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 17, 19, 22}
	target := 2
	result := checkExist(array, target)
	fmt.Println("if exist in the array", result)
}

func checkExist(array []int, target int) bool {
	var result = false
	low := 0
	high := len(array) - 1
	middle := 0
	for low <= high {
		middle = (high + low) / 2
		fmt.Println("check middle:", middle)
		if array[middle] == target {
			result = true
			break
		} else if array[middle] < target {
			fmt.Println("up")
			low = middle + 1
		} else {
			fmt.Println("down")
			high = middle - 1
		}
	}
	return result
>>>>>>> b294ca31202dc73888b65b0651de654db5a46cbc
}
