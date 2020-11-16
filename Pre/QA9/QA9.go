package main

import (
	"fmt"
)

func main() {
	var m = "134524394587937548923746009237434567982374798679"
	var n = "1234565u76790230476762345673983409382070123975927"
	fmt.Println(maxMatchString(m, n))

}

func maxMatchString(m, n string) string {
	var matchString = ""
	var maxMatchStringStart = -1
	var length = 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(n); j++ {
			if m[i] == n[j] {
				for k := 0; k < len(m)-i && k < len(n)-j; k++ {
					if m[i+k] != n[j+k] {
						fmt.Println("max index:", maxMatchStringStart, length)
						break
					}
					fmt.Println("equal:", i, j)
					if k+1 > length {
						maxMatchStringStart = i
						length = k + 1
					}

				}
			}
		}
	}
	matchString = m[maxMatchStringStart : maxMatchStringStart+length]
	return matchString
}
