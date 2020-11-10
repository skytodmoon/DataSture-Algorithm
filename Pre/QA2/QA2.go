package main

import "fmt"

func main() {
	var a = []int{1, 3, 4, 1, 3, 2, 3, 5, 10, 11, 4, 3}
	var timeTmp = 0
	var timeMax = 0
	var valMax = -1

	var i, j int = 0, 0
	for i = 0; i < len(a); i++ {
		timeTmp = 0
		for j = 0; j < len(a); j++ {
			if a[i] == a[j] {
				timeTmp++
			}
		}
		if timeTmp > timeMax {
			timeMax = timeTmp
			valMax = a[i]
		}
	}

	fmt.Println(valMax, timeMax)

}
