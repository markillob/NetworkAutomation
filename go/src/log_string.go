package main

import "os"
import "fmt"

func main() {
	var stringTest string
	stringTest="markilo"
	for i := 0; i < len(stringTest); i++ {
		for a := (i+1); a < len(stringTest); a++ {
			if stringTest[i] == stringTest[a] {
				fmt.Printf("String has duplicate values %c \n",stringTest[i])
				os.Exit(1)
			}
		}
	}
	fmt.Println("String has no duplicate values ",stringTest)
}
	