package main

import "fmt"
import "strconv"

func main() {
	str := "1ola"
	//fmt.Println(compress("aaannnddd"))
	fmt.Println(isNumber(str[0]))
}

//s4b2s1

func compress(input string) string {
	if len(input) == 0 {
		return input
	}
	count := 1
	tempLetter := input[0]
	output := ""
	for i := 1; i < len(input); i++ {
		if input[i] == tempLetter {
			count++
		} else {
			output = output + string(tempLetter) + strconv.Itoa(count)
			tempLetter = input[i]
			count = 1
		}
	}
	output = output + string(tempLetter) + strconv.Itoa(count)
	if len(input) > len(output) {
		return output
	}
	return input
}

func decompress(input string) string {
	if len(input) == 0 {
		return
	}
	tempLetter := input[0]
	for i := 1; i < len(input); i++ {

	}
}

func isNumber(input byte) bool {
	if input > 47 && input < 58 {
		return true
	}
	return false
}
