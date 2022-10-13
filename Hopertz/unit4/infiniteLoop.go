package main

import ("fmt";"strings")

// Infine loops in Go is written by  for {}
func main(){
	// The following code snippet repeatedly waits for the user to input a string, 
	//until the user enters the string QUIT:

	for {
	  // Also Go uses continue and break same as with other 
	  //programming langauges such pyhton
      fmt.Println("Enter QUIT to exit")
      var input string
      fmt.Print("Please enter a string:")
      fmt.Scanln(&input)
      if strings.ToUpper(input) == "QUIT" {
          break
      }
 }
	

}
