package main

import "fmt"
import "strconv"

func main() {
	str := "s3d1c1s5b"
	//fmt.Println(compress("aaannnddd"))
	fmt.Println(decompress(str))
}

func decompress(input string) string {
	if len(input) == 0 {
		return "String Empty, try again "
	}
	tempLetterInitial := input[0]
	output := ""
	for i := 1; i < len(input); i++ {
		if isNumber(input[i]) {
			tempStringLetter, _ := strconv.Atoi(string(input[i]))
			numTempLetter := int(tempStringLetter)
			for count2 := 0; count2 < numTempLetter; count2++ {
				//fmt.Println(numTempLetter)
				output = output + string(tempLetterInitial)
			}
		} else {
			tempLetterInitial = input[i]
		}
	}
	output = output + string(input[(len(input)-1)])
	return output
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

func isNumber(input byte) bool {
	if input > 47 && input < 58 {
		return true
	}
	return false
}
