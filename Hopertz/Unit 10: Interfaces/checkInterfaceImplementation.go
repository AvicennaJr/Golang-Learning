package main

import (
	"fmt"
	 "math"
)

type Shape interface {
 Area() float64
}
type Circle struct {
 radius float64
}

   
// Circle implements Shape
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}


func (c Circle) String() string {
	return fmt.Sprintf("Area is %v Circumference is %v",c.Area(), c.Circumference())
   
}


func main() {
	c1 := Circle{radius: 5}
	var v interface{} = c1
	s, ok := v.(Shape)



	if ok {
		fmt.Println(s.Area())
	}

	
}
