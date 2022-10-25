package main

import (
	"fmt"
)

func main() {
	//Interpolating strings

	queue := 5
	name := "John"
	// you can’t simply directly concatenate string and integer values
	//s := name + ", your queue number is:" + queue

	// This is okay but we can do better
	//s := name + ", your queue number is:" + strconv.Itoa(queue)

	// Better solution is to use the Sprintf() function from the fmt package:
	s := fmt.Sprintf("%s, your queue number is %d", name, queue)
	// The Sprintf() function formats a string based on the formatting verbs (such as %d and %s)
	fmt.Println(s)
}
