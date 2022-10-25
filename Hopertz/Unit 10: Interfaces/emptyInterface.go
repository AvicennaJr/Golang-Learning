package main 

import (
	"fmt"
)


// An empty interface is useful when you want to handle data of unknown type
func doSomething(i interface{}) {
	fmt.Println(i)
}
func main(){
	doSomething("Hi!") // string
	doSomething(3.14) // float
	doSomething([]int{3, 4}) // array
}

