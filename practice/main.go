package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	// keyAndValue()
	//pointer()
	//slicesPracice()
	mapPractice()

}

func mapPractice() {
	states := make(map[string]string)
	states["WB"] = "West Bengal"
	states["HYD"] = "Hyderabad"
	states["BAN"] = "Bangalore"
	fmt.Println(states)

	westBengal := states["WB"]
	fmt.Println(westBengal)

	delete(states, "HYD")
	fmt.Println(states)

	states["DEL"] = "Delhi"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

func slicesPracice() {
	var color = []string{"Red", "Orange", "Blue"}
	fmt.Println(color)
	color = append(color, "Purple")
	fmt.Println(color)

	color = append(color[0:len(color)])
	fmt.Println(color)

	color = append(color[1 : len(color)-1])
	fmt.Println(color)

	numbers := make(
		[]int,
		5,
		5)
	numbers[0] = 73
	numbers[1] = 10
	numbers[2] = 49
	numbers[3] = 48
	numbers[4] = 75
	fmt.Println(numbers)
	numbers = append(numbers, 81)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}

func pointer() {
	a := 45
	aValue := &a
	fmt.Println("Memory Address", aValue)
	fmt.Println("Actual Value", *aValue)

	//changing Value through pointer
	*aValue = *aValue / 21
	fmt.Println("Pointer ", *aValue)
	fmt.Println("Varible ", a)
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
