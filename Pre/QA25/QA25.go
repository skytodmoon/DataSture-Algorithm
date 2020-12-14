package main

import "fmt"

func main() {
	var s = "aAAbbbb"
	var j = "aA"
	fmt.Println(numJewelsInStones(j, s))
	var s1 = "ZZ"
	var j1 = "z"
	fmt.Println(numJewelsInStones(j1, s1))
	var s2 = "ZZIhHhgkKkGgjJjk"
	var j2 = "zACdF"
	fmt.Println(numJewelsInStones(j2, s2))

}

//hash表数据量较少时，性能优势不明显
func numJewelsInStones(J string, S string) (sum int) {
	jewelsSet := map[byte]bool{}
	for i := range J {
		jewelsSet[J[i]] = true
	}
	for i := range S {
		if jewelsSet[S[i]] {
			sum++
		}
	}
	// for _, s := range S {
	// 	for _, j := range J {
	// 		if s == j {
	// 			sum++
	// 			break
	// 		}
	// 	}
	// }
	return
}
