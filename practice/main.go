package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//varDeclaration()
	//inputFromUser()
	convertNumber()
}

func convertNumber() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat)
	}
}

func inputFromUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You Entered :", input)
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
