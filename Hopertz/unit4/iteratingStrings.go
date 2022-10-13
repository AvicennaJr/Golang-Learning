package main

import "fmt"

func main(){
   for pos, char := range "Hello, world!" {
       fmt.Println(pos, char)
	   // char value is Unicode code for each character in the string. 
    }
   // To get actual character itself use Printf with format specifier
   for pos, char := range "Hello, world!" {
      fmt.Printf("%d %c\n", pos, char)
	}
}
