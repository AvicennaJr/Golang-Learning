package main

import "fmt"



func main(){
   Outerloop:
       for i := 1; i <= 5; i++ {
            for j := 1; j <= 5; j++ {
                if i == 3 {
                    continue Outerloop
                }
                fmt.Printf("%d * %d = %d\n", i, j, i*j)
            }
            fmt.Println("-----------")
       }
}
