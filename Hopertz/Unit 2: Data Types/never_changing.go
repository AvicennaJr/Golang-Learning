package main

import "fmt"

const publisher = "Wiley"

var num1 = 5  //if a variable is declared outside a function and not used anywhere in the program,the compiler wouldn’t complain:
func main() {
   var num2 = 13 //For variables declared inside function bodies, you need to use them, or else you’ll get an error message.
   _ = num2  // The complier is now happy
   fmt.Println(publisher)
  
}
