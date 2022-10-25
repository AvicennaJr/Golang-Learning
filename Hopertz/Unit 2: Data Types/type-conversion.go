package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//variable types - Go
	//Discovering the type of a variable
	start := time.Now() // need to import the time package

	fmt.Printf("%T\n", start)

	//Scanf() function scans text read from
	//standard input (the console), storing successive space-separated values into successive
	//arguments as determined by the format specif

	//The following is a program that reads user input and converts it to integer

	//var input string
	//fmt.Print("Please enter your age: ")
	//fmt.Scanf("%s", &input)
	//age, err := strconv.Atoi(input) // convert string to
	// int
	//if err != nil { // an error occurred
	// fmt.Println(err)
	// } else {
	// fmt.Println("Your age is:", age)
	//}

	b, err := strconv.ParseBool("t")
	fmt.Println(b)        // true
	fmt.Println(err)      // <nil>
	fmt.Printf("%T\n", b) // bool

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)        // 3.1415
	fmt.Println(err)      // <nil>
	fmt.Printf("%T\n", f) // float64

	//you can use int(), float32(), and float64() functions to convert to numeric types

	num1 := 5
	num2 := float32(num1)
	num3 := float64(num2)
	num4 := float32(num3)
	num5 := int(num4)
	fmt.Printf("%T\n", num1) // int
	fmt.Printf("%T\n", num2) // float32
	fmt.Printf("%T\n", num3) // float64
	fmt.Printf("%T\n", num4) // float32
	fmt.Printf("%T\n", num5) // int
}
