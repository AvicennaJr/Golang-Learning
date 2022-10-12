package main

import "fmt"

func main(){
	num := 9
	if num % 2 == 1 {
       fmt.Println("Number is odd")
    } else {
        fmt.Println("Number is even")
}

// non zero in go does not evaluate to true 
// using "if num % 2 {}" as  a condition doesn't 
// Short-circuiting is possible in Go using // and &&
//Ternary operators arent there in Go

}
