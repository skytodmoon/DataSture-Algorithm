package main

import (
	"fmt"
	"time"
)

func main() {
	var a = []int{1, 3, 4, 1, 3, 2, 3, 5, 10, 11, 4, 3}
	var timeTmp = 0
	var timeMax = 0
	var valMax = -1

	var i, j int = 0, 0
	//first method
	start0 := time.Now().UnixNano()
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
	end0 := time.Now().UnixNano()
	fmt.Println("Time cost0 is :", (end0-start0)/1000)

	//second method
	start1 := time.Now().UnixNano()
	checkMap := make(map[int]int)
	for i = 0; i < len(a); i++ {
		if _, ok := checkMap[a[i]]; ok {
			checkMap[a[i]]++
		} else {
			checkMap[a[i]] = 1
		}
	}

	var targetIndex = 0
	var targetValue = -1

	for value := range checkMap {
		if checkMap[value] > targetValue {
			targetIndex = value
			targetValue = checkMap[value]
		}
	}

	fmt.Println(targetIndex, targetValue)
	end1 := time.Now().UnixNano()
	fmt.Println("Time cost1 is :", (end1-start1)/1000)

}
