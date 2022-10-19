# Interfaces

An interface defines the behaviour of an object, specifying the methods that it
needs to implement.

## Defining an Interface and Using it

<pre>

package main
import "fmt"

type Shape interface {
    Area() float64
}
type Circle struct {
    radius float64
}

type Square struct {
    length float64
}
// Circle implements Shape
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}
// Square implements Shape
func (s Square) Area() float64 {
    return s.length * s.length
}
func calculateArea(listOfShapes []Shape) {
for _ , shape := range listOfShapes {
    fmt.Println("Area: ", shape.Area())
    }
}


func main() {
    c1 := Circle{radius: 5}
    s1 := Square{length: 6}
    shapes := []Shape{c1, s1}
    calculateArea(shapes)
}

</pre>

## Stringer Interface

You can customize how fmt.Println() will print an object by customizing the
String() interface used by Println()

<pre>


func (c Circle) String() string {
    return fmt.Sprintf(
         "Area is %v Circumference is %v",
            c.Area(), c.Circumference())
    }


func main() {
    c1 := Circle{radius: 5, name: "c1"}
    fmt.Println(c1)
    }
    // Output will be
    // Area is 78.53981633974483 Circumference is 31.41592653589793

</pre>

## Empty Interface

Very useful if you don't know the data type

<pre>
func doSomething(i interface()) {
    fmt.Println(i)
}
</pre>

## Finding out if a value implements a specific interface

<pre>
// Asign a variable to a type

c1 := Circle{radius: 5, name: "c1"}

// Then assign it to an empty interface

var v interface() = c1

// Then use *type assertion* gain access to an interface's concrete value.
// *Type assertion allows you to test if a value stored in an interface variable is of a
particular type*

v, ok := v.(Shape) // test if v has the shape interface

// if it does ok==true. Otherwise false.

if ok {
    fmt.Prinln(v.Area)
}

</pre>

> Thats all about interfaces
