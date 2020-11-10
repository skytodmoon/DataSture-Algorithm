package main

import (
	"fmt"
	"time"
)

func main() {
	var i, j, k, count int = 0, 0, 0, 0
	start0 := time.Now().UnixNano()
	for i = 0; i <= 100/7; i++ {
		for j = 0; j <= 100/3; j++ {
			for k = 0; k <= 100/2; k++ {
				if i*7+j*3+k*2 == 100 {
					count++
					//fmt.Println(i, j, k)
				}
			}
		}
	}
	end0 := time.Now().UnixNano()
	fmt.Println("Time cost0 is :", (end0-start0)/1000)
	start := time.Now().UnixNano()
	for i = 0; i <= 100/7; i++ {
		for j = 0; j <= (100-i*7)/3; j++ {
			for k = 0; k <= (100-i*7-j*3)/2; k++ {
				if i*7+j*3+k*2 == 100 {
					count++
					//fmt.Println(i, j, k)
				}
			}
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("Time cost is :", (end-start)/1000)
	count = 0
	start2 := time.Now().UnixNano()
	for i = 0; i <= 100/7; i++ {
		for j = 0; j <= (100-i*7)/3; j++ {
			if 100-i*7-j*3 >= 0 && (100-i*7-j*3)%2 == 0 {
				count++
				//k = (100 - i*7 - j*3) / 2
				//fmt.Println(i, j, k)
			}
		}
	}
	end2 := time.Now().UnixNano()
	fmt.Println("Time cost2 is :", (end2-start2)/1000)
	//fmt.Println(count)

	count = 0
	start3 := time.Now().UnixNano()
	for i = 0; i <= 100/7; i++ {
		for j = 0; j <= 100/3; j++ {
			if 100-i*7-j*3 >= 0 && (100-i*7-j*3)%2 == 0 {
				count++
				//k = (100 - i*7 - j*3) / 2
				//fmt.Println(i, j, k)
			}
		}
	}
	end3 := time.Now().UnixNano()
	fmt.Println("Time cost3 is :", (end3-start3)/1000)
}
