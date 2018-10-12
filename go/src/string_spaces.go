package main

import "fmt"

//space is 32
func main() {
	var mainString string
	mainString = "Ho la  "
	totalLength := len(mainString)
	numberSpaces := 0
	//Get number of Spaces
	for i := 0; i < totalLength; i++ {
		if mainString[i] == 32 {
			numberSpaces++
		}
		//fmt.Println(mainString[i])
		//fmt.Printf("%c", mainString[i])
	}
	numberSpaces = numberSpaces / 3
	endOfArray := len(mainString) - (numberSpaces * 2)
	fmt.Println(numberSpaces)
	fmt.Println(endOfArray)
	// move array
	for whiteSpace := 0; whiteSpace < endOfArray; whiteSpace++ {
		if mainString[whiteSpace] == 32 {
			var temp1 byte
			var temp2 byte
			temp1 = mainString[whiteSpace+1]
			temp2 = mainString[whiteSpace+2]
			fmt.Println(temp1)
			fmt.Println(temp2)
			mainString[endOfArray] = temp1
			mainString[endOfArray] = temp1
		}
	}
	fmt.Println(mainString)
}
