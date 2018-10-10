package main

import "os"
import "fmt"

func main() {
	var stringTest = "markilob"
	var contador int
	contador = 1
	for i := 0; i <= len(stringTest); i++ {
		for a := i; a <= len(stringTest); a++ {
			if stringTest[a] == stringTest[i] {
				contador += 1
				if contador > 1 {
					fmt.Println(stringTest[a])
					os.Exit(1)
				}
			} else {
				contador == 0
			}
		}
	}
	fmt.Println("String has all unique values")
}
