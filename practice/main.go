package main

import (
	"fmt"
)

func main() {
	varDeclaration()
}

func varDeclaration() {
	var str = "Hello from Go!"
	fmt.Println(str)
	fmt.Printf("%T \n", str)

	var inte = 42
	fmt.Println(inte)
	fmt.Printf("%T \n", inte)

	strNew := "Hello from Go!"
	fmt.Println(strNew)
	fmt.Printf("%T \n", strNew)
}
