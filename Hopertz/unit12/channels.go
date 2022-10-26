

package main 

import (
	"fmt"
	"time"
)

// goroutines to communicate with one another through
// pipes known as channels

// Create a channel using make ch := make(chan int)

// write to a channel ch <- 5

// Retrieve a value from the channel value := <- ch


//---send data into a channel---
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel...",time.Now())
   }
   
//---getting data from the channel---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
    fmt.Println("String retrieved from channel:", <-ch,time.Now())
   } 

func main(){
    ch := make(chan string)
	go sendData(ch)
    go getData(ch)
    fmt.Scanln()
	
}

