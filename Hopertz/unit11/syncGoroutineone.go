package main
import (
 "fmt"
 "math/rand"
 "sync"
 "sync/atomic"
 "time"
)
var balance int64

func credit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 10; i++ {
	// adds 100 to balance atomically
	atomic.AddInt64(&balance, 100)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
   }
func debit(wg *sync.WaitGroup) {
    // notify the WaitGroup when we are done
    defer wg.Done()
    for i := 0; i < 5; i++ {
    // deducts -100 from balance atomically
    atomic.AddInt64(&balance, -100)
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    }
}
func main() {
    var wg sync.WaitGroup
    balance = 200
    fmt.Println("Initial balance is", balance)
    wg.Add(1) // add 1 to the WaitGroup counter
    go credit(&wg)
    wg.Add(1) // add 1 to the WaitGroup counter
    go debit(&wg)
    wg.Wait() // blocks until WaitGroup counter is 0
    fmt.Println("Final balance is", balance)
}
