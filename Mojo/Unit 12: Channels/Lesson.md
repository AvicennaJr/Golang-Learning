# Communicating between Goroutines Using Channels

## Introduction

<pre>
ch := make(chan int)

ch <- 5

value := <- ch

</pre>

## How they work

<pre>
package main
import (
    "fmt"
    "time"
)
//---send data into a channel---
func sendData(ch chan string) {
    fmt.Println("Sending a string into channel...")
    time.Sleep(2 * time.Second)
    ch <- "Hello"
}
//---getting data from the channel---
func getData(ch chan string) {
    fmt.Println("String retrieved from channel:", <-ch)
}
func main() {
    ch := make(chan string)
    go sendData(ch)
    go getData(ch)
    fmt.Scanln()
}

</pre>

## How they are used

A bit complex topic but code should be explanatory

<pre>

// Create a function that excepts a slice of intergers and a channel 
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
    sum += v
    }
    c <- sum // assigns values to the channel

}



func main() {
    s := []int{}
    sliceSize := 10 // assign valus to the slice

    for i := 0; i < sliceSize; i++ {
        s = append(s, rand.Intn(100))
    }

    // create a new channel
    c := make(chan int)

    // divide the slice into two to make concurrent summation

    partSize := 2
    parts := sliceSize / partSize
    i := 0
    for i < parts {
        go sum(s[i*partSize:(i+1)*partSize], c) // take a portion of the initial slice
        i += 1
    }

    i = 0
    total := 0
    for i < parts {
        partialSum := <-c // read from channel
        fmt.Println("Partial Sum: ", partialSum)
        total += partialSum
        i += 1
    }
    fmt.Println("Total: ", total)
    fmt.Scanln()
}

</pre>

## Iterating through channels

<pre>

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
</pre>

## Asynchronously Waiting on Channels

<pre>


package main
import (
    "fmt"
    "time"
)

func fib(n int, c chan int) {
    a, b := 1, 1
    for i := 0; i < n; i++ {
        c <- a
        a, b = b, a+b
        time.Sleep(2 * time.Second)
    }
    close(c)
}
func counter(n int, c chan int) {
    for i := 0; i < n; i++ {
        c <- i
        time.Sleep(1 * time.Second)
    }
    close(c)
}

func main() {
    c1 := make(chan int)
    c2 := make(chan int)
    go fib(10, c1)
    go counter(10, c2)
    c1Closed := false
    c2Closed := false
    for {
        select {
             case n, ok := <-c1:
                if !ok {
                    // channel closed and drained
                    c1Closed = true
                    if c1Closed && c2Closed {
                        return
                    }
                } else {
                    fmt.Println("fib()", n)
                }
            case m, ok := <c2:
                if !ok {
                    // channel closed and drained
                    c2Closed = true
                    if c1Closed && c2Closed {
                        return
                    }
                } else {
                    fmt.Println("counter()", m)
                }
            }
        }
}

</pre>

You can also run the for loop as a goroutine;

<pre>

go func() {
    for {
        select{
            case n, ok := <-c1:
            .......
            case n, ok := <-c2:
            .....
            }
        }
    }()

fmt.Print("Continue to do something else...")
fmt.Scanln() // needed here to prevent the program from existing before all
// the channel values are read
</pre>

## Using Buffered Channels

A buffered channel allows multiple values to be stored in the channel. Your 
code will only block when you try to send a value to a channel that is full or 
when you try to read from an empty channel.

<pre>


func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
    fmt.Println("Done and can continue to do other work")
}
func main() {
    s := []int{}
    sliceSize := 10

    for i := 0; i < sliceSize; i++ {
        s = append(s, rand.Intn(100))
    }
    c := make(chan int, 5) // buffered channel of length 5
    partSize := 2
    parts := sliceSize / partSize
    i := 0
    for i < parts {
        go sum(s[i*partSize:(i+1)*partSize], c)
        i += 1
    }

    i = 0
    total := 0
    time.Sleep(1 * time.Second) // simulate retrieving at a later time
    for i < parts {
        partialSum := <-c // read from channel
        fmt.Println("Partial Sum: ", partialSum)
        total += partialSum
        i += 1
    }
    fmt.Println("Total: ", total)
    fmt.Scanln()
}
</pre>

> It will take a little more time to fully understand this
