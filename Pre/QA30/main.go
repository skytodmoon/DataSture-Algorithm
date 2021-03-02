package main

func main() {
	var encoded = []int{1, 2, 3}
	var first = 1

	var resultArr = decode(encoded, first)
	for _, v := range resultArr {
		println(v)

	}
}

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first
	for i := 0; i < len(encoded); i++ {
		arr[i+1] = encoded[i] ^ arr[i]
	}
	return arr
}
