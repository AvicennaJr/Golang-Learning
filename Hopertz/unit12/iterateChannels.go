package main
import (
 "fmt"
 "time"
)
func fib(n int, c chan int) {
   a, b := 1, 1
   for i := 0; i < n; i++ {
   c <- a // blocked until value is received from channel
   a, b = b, a + b
   time.Sleep(1 * time.Second)
 }
   close(c) // close the channel
}
func main() {
   c := make(chan int)
   go fib(10, c)
   for i := range c { // read from channel until channel is closed
   fmt.Println(i)
 }
}
