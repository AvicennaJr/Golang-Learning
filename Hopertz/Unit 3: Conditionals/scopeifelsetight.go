package main 
import "fmt"

func doSomething() (int, bool) {
 //...
 // just an example of some return values
 return 5, true
}

func main(){
	v, err := doSomething()
   if err {
       fmt.Println("I am error")
   } else {
       fmt.Println(v)
  }
}
