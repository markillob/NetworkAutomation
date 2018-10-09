package main 

import "fmt"
import "os"

func main() {
	var uniqueString string 
	fmt.Print("Enter string: ")
	fmt.Scanf("%s",&uniqueString)
	fmt.Println("String to check:", uniqueString)
	//checkStringLenght
	var stringLenght int
    uniqueMap := make(map[byte]int)
	// var uniqueMap map[byte]int
    stringLenght = len(uniqueString)
	for i := 0; i<stringLenght; i++{
		value, ok :=uniqueMap[uniqueString[i]]
		if ok {
			fmt.Printf("String has duplicate values %c\n",uniqueString[value])
			os.Exit(1)
		}
		uniqueMap[uniqueString[i]]= 1
	}
	fmt.Printf("String has no duplicate values")
	//fmt.Printf("%v", uniqueMap)
	//for i := 1; i<=10; i++{
	//	fmt.Println(i)
	//}
	//fmt.Printf("%c", x[1])

}