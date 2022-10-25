
package main

import "fmt"
   var OS [3]string
   

func main(){
	OS[0] = "iOS"
    OS[1] = "Android"
    OS[2] = "Windows"

	// The range keyword returns index and values of the OS array
	for i, v := range OS {
        fmt.Println(i, v)
    }

	// if you dont care about the index
	for _, v := range OS {
       fmt.Println(v)
    }

	// if you only care about the index
	for i, _ := range OS {
    fmt.Println(i)
    }

	// Or
    for i := range OS {
    fmt.Println(i)
   }

   
  
}
