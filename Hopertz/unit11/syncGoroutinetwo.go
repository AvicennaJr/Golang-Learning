package main
import (
 "fmt"
 "math/rand"
 "sync"
 "sync/atomic"
 "time"
)
var balance int64
var mutex = &sync.Mutex{}
func credit(wg *sync.WaitGroup) {
 // notify the WaitGroup when we are done
 defer wg.Done()
 for i := 0; i < 10; i++ {
 mutex.Lock()
 // adds 100 to balance atomically
 atomic.AddInt64(&balance, 100)
 fmt.Println("After crediting, balance is", balance)
 mutex.Unlock()
 time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
 }
}
func debit(wg *sync.WaitGroup) {
 // notify the WaitGroup when we are done
 defer wg.Done()
 for i := 0; i < 5; i++ {
	mutex.Lock()
	// deducts -100 from balance atomically
	atomic.AddInt64(&balance, -100)
	fmt.Println("After debiting, balance is", balance)
	mutex.Unlock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
   }
func main() {
   var wg sync.WaitGroup
   balance = 200
   fmt.Println("Initial balance is", balance)
   wg.Add(1)
   go credit(&wg)
   wg.Add(1)
   go debit(&wg)
   wg.Wait()
   fmt.Println("Final balance is", balance)
}
