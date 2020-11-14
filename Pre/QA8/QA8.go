package main

import "fmt"

func stringCheck(m, n string) {

	var index = -1
	if len(m) < len(n) {
		fmt.Println("Not long enough!")
		return
	}
	for i := 0; i < (len(m) - len(n) + 1); i++ {
		if m[i] == n[0] {
			check := true
			j := 0
			for j = 0; j < len(n); j++ {
				if m[i+j] != n[j] {
					check = false
					break
				}

			}
			if check && j < (len(n)-1) {
				fmt.Println("Not long enough!")
				check = false
			}

			if check {
				index = i
				fmt.Println("Find it at:", index)
			}

		}
	}
}

func main() {
	m := "goodgooglegooglegoogkegooglegooggoog"
	n := "google"
	stringCheck(m, n)

}
