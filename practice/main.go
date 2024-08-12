package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//varDeclaration()
	//inputFromUser()
	//convertNumber()
	//mathOperations()
	//dateAndTime()
	keyAndValue()
}

func keyAndValue() {
	k := make(map[string]int)
	k["test"] = 42
	fmt.Println(k)
}

func dateAndTime() {
	n := time.Now()
	fmt.Println("Time Now is :", n)

	t := time.Date(2004, time.January, 12, 9, 30, 0, 0, time.Local)
	fmt.Println("We Got Married at", t)
	fmt.Println(t.Format(time.ANSIC))
}

func mathOperations() {
	intSum()
	intfloat()
}

func intfloat() {
	//f1, f2, f3 := 1.1, 1.23, 1.234
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum : ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float Sum Rounded : ", floatSum)

	radius := 15.5
	circumference := radius * 2 * math.Pi
	fmt.Printf("circumferenc %.2f", circumference)
}

func intSum() {
	i1, i2, i3 := 1, 11, 111
	intSum := i1 + i2 + i3
	fmt.Println("Interger Sum : ", intSum)
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
