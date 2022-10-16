package main

import "fmt"

func main() {
	var name string

	fmt.Print("Hello! What is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Println("Nice to meet you", name)
}
