package main

import "strings"

func main() {
	var address = "1.1.1.1"
	println(defangIPaddr(address))

}

// func defangIPaddr(address string) string {
// 	return strings.ReplaceAll(address, ".", "[.]")
// }

func defangIPaddr(address string) string {
	var s strings.Builder
	for i := range address {
		if address[i] == '.' {
			s.WriteString("[.]")
		} else {
			s.WriteByte(address[i])
		}
	}
	return s.String()
}
