package main
//Declaring Always-Changing Variables

import "fmt"

func main() {

    var num1 = 5     // type inferred
    var num2 int // explicitly typed
    fmt.Println(num1)
    fmt.Println(num2)
    var num3 float32 // floating point variable
    var name string // string variable
    var raining bool // boolean variable
    fmt.Println(num3) // 0
    fmt.Println(name) // "" (empty string)
    fmt.Println(raining) // false
    var rates float32 = 4.5 // declared as float32 and // then initialized
    fmt.Println(rates)
    
   // Using the short variable declaration operator
   // firstName := "Wei-Meng"
   //:= only works in functions and the lower case 't' is so that it is only visible to the package (unexported).
   // firstName, lastName, age := "Wei-Meng", "Lee", 25
   // var firstName, lastName string = "Wei-Meng", "Lee"
   
   // var firstName, lastName string, age int = "Wei-Meng","Lee", 25  not allowed
   
   // var (
   //    firstName string="Wei-Meng"
   //    lastName string="Lee"
   //    age int=25
   //)
   //  }
   
   //The := operator can only be used for declaring and initializing variables inside functions.

