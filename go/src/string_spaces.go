package main

import "fmt"

//import "reflect"

//space is 32
func main() {
	//var mainString[7] string
	mainString := make([]string, 15)
	mainString[0] = "a"
	mainString[1] = "r"
	mainString[2] = "r"
	mainString[3] = "e"
	mainString[4] = " "
	mainString[5] = "o"
	mainString[6] = "r"
	mainString[7] = "o"
	mainString[8] = " "
	mainString[9] = "s"
	mainString[10] = "i"
	mainString[11] = " "
	mainString[12] = " "
	mainString[13] = " "
	mainString[14] = " "
	fmt.Println(mainString)
	totalLength := len(mainString)
	numberSpaces := 0
	//Get number of Spaces
	for i := 0; i < totalLength; i++ {
		if mainString[i] == " " {
			numberSpaces++
		}
		//fmt.Println(mainString[i])
		//fmt.Printf("%c", mainString[i])
	}
	numberSpaces = numberSpaces / 3
	//fmt.Println(numberSpaces)
	endOfArray := len(mainString) - (numberSpaces * 2)
	fmt.Println(endOfArray)
	// move array
	//endOfArray2 := endOfArray - 1
	for whiteSpace := 0; whiteSpace < endOfArray; whiteSpace++ {
		count := whiteSpace
		if mainString[whiteSpace] == " " {
			for z := (len(mainString) - 3); z > (whiteSpace); z-- {
				fmt.Println(z)
				fmt.Println(mainString)
				//fmt.Println(endOfArray2)
				mainString[z] = mainString[count]
				count++
			}
			//mainString[whiteSpace] = "%"
			//mainString[whiteSpace+1] = "2"
			//mainString[whiteSpace+2] = "0"
			fmt.Println(mainString)
		}
	}
	fmt.Println(mainString)
}

/*


	fmt.Println(numberSpaces)
	fmt.Println(endOfArray)
	// move array
	for whiteSpace := 0; whiteSpace < endOfArray; whiteSpace++ {
		if mainString[whiteSpace] == 32 {
			fmt.Println(mainString[whiteSpace])
			mainString[whiteSpace] = 30
			fmt.Println(mainString[whiteSpace])
			//var temp1 byte
			//var temp2 byte
			//temp1 = mainString[whiteSpace+1]
			//temp2 = mainString[whiteSpace+2]
			//
			//temp3:= string(mainString[:whiteSpace+2])
			//fmt.Println(reflect.TypeOf(temp3))
			//fmt.Println(temp1)
			//fmt.Println("\n",temp2)
			//fmt.Println(temp3)
			//mainString[endOfArray] = temp1
			//mainString[endOfArray] = temp1
		}
	}
}
*/
