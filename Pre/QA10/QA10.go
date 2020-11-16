package main

import "fmt"

func main() {
	var sourceStr = "today is a fine day"
	var targetStr = reverseStr(sourceStr)
	fmt.Println(targetStr)
}

func reverseStr(str string) string {
	var resultStr = ""
	var strArray = []string{}
	var index = 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			strArray = append(strArray, str[index:i])
			index = i + 1
			continue
		}
		if i == len(str)-1 {
			strArray = append(strArray, str[index:len(str)])
		}

	}

	for j := len(strArray) - 1; j >= 0; j-- {
		if j > 0 {
			resultStr = resultStr + strArray[j] + " "
		} else {
			resultStr = resultStr + strArray[j]
		}
	}
	return resultStr
}
