package main

import "fmt"

func main(){
	grade := ""
    switch grade {
       case "A", "B", "C", "D":
         fmt.Println("Passed")
       case "F":
         fmt.Println("Failed")
       default:
         fmt.Println("Undefined")
}
}
