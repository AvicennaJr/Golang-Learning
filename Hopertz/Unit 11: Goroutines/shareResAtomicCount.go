package main
import (
 "fmt"
 "math/rand"
 "sync/atomic"
 "time"
)
var balance int64
func credit() {
    for i := 0; i < 10; i++ {
    // adds 100 to balance atomically
    atomic.AddInt64(&balance, 100)
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    }
}
func debit() {
    for i := 0; i < 5; i++ {
    // deducts -100 from balance atomically
    atomic.AddInt64(&balance, -100)
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    }
}
func main() {
    balance = 200
    fmt.Println("Initial balance is", balance)
    go credit()
    go debit()
    fmt.Scanln()
    fmt.Println(balance)
}
