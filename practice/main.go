package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//varDeclaration()
	inputFromUser()
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

func inputFromUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You Entered :", input)
}
