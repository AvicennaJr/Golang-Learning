package main

import (
 "fmt"
)

// An interface is an abstract type. It describes all the methods that a type can implement.

type DigitsCounter interface {
 CountOddEven() (int, int)
}
type DigitString string


// DigitString implements DigitsCounter                                    
func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0
	for _, digit := range ds {
	   if digit%2 == 0 {
      	evens++
	   } else {
	       odds++
       }
	}
	return odds, evens
   }
func main() {
	s := DigitString("123456789")
    fmt.Println(s.CountOddEven()) // 5 4


	// Because DigitsCounter is a type in itself, you can also create a variable of type
    // DigitsCounter and then assign the variable s to it:
    var d DigitsCounter
    d = s
    fmt.Println(d.CountOddEven()) // 5 4
}
