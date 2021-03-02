package main

func main() {
	var guess = []int{1, 2, 3}
	var answer = []int{3, 2, 1}
	println(game(guess, answer))
}

func game(guess []int, answer []int) int {
	var result = 0
	//for i := 0; i < len(guess); i++ {
	//	if guess[i] == answer[i] {
	//		result++
	//	}
	//}

	for k, v := range guess {
		if v == answer[k] {
			result++
		}
	}
	return result
}
