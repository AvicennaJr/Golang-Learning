package main

import "fmt"

func main(){
  //for (init; condition; post) {}
  //++x and --x are not allowed in Go 
  // Parentheses surrounding the for loop are  not allowed
  fmt.Println("Normal for loop in Go")
  for i:=1;i<=5; i++ {
	fmt.Println(i)
  }

  // init & post are optional see the following 
  //  we are printing the fibonacci series
  fmt.Println("The fibonacci series")
  max := 100
  a, b := 0, 1
  for ;b <= max; {
     fmt.Println(a)
     a, b = b, a+b
}

// init and post can be ommmited
  fmt.Println("The fibonacci series with no init or post")
  newmax := 100
  c, d := 0, 1
  for d <= newmax {
     fmt.Println(c)
     c, d = d, c+d
}
   
   fmt.Println("Reduce the number of lines for the Fibonacci sequence")
   anotherMax := 100
   for e, f := 0, 1; f <= anotherMax ; e, f = f, e+f {
     fmt.Println(e)
   }
}
