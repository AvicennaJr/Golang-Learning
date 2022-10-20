package main

import (
 "fmt"
)
func main() {
   
	var c [3]string
    c[0] = "iOS"
    c[1] = "Android"
    c[2] = "Windows"

	fmt.Println(c[0:2]) // [iOS Android]
    fmt.Println(c[:2]) // [iOS Android]
	fmt.Println(c[1:]) // [Android Windows]
	fmt.Println(c[:]) // [iOS Android Windows]

	//when slicing  length and capacity of the resultant slice may change
	t := []int{1, 2, 3, 4, 5}
    fmt.Println(len(t)) // 5
    fmt.Println(cap(t)) // 5

	t = t[2:4]

	// length and capacity updated
	fmt.Println(t) // [3 4]
    fmt.Println(len(t)) // 2
    fmt.Println(cap(t)) // 3
  
	// the capacity of the resultant slice will depend on the start index of
    // the first element of the slice. If you slice from 1:3, then your capacity will change

	//Whether you’re slicing an array or a slice, the result is always a slice.
	s := []int{1, 2, 3, 4, 5}
	s = s[1:3]
    fmt.Println(s) // [2 3]
    fmt.Println(len(s)) // 2
    fmt.Println(cap(s)) // 4


}
