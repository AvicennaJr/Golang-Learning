package main

import "fmt"

func main(){
	grade := "B"
    switch grade {
       case "A":
        fallthrough
       case "B":
        fallthrough
	   case "C":
        fallthrough
       case "D":
        fmt.Println("Passed")
       case "F":
       fmt.Println("Failed")
       default:
          fmt.Println("Absent")
}
}
