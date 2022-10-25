package main 

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}
   
func (p Person) String() string {
	return fmt.Sprintf("%v %v (%d years old)",
	p.FirstName, p.LastName, p.Age)
}

func main(){
	me := Person{"Lugano", "Abel", 24}
    fmt.Println(me)
 
}

