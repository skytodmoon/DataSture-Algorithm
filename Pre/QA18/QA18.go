package main

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
}
