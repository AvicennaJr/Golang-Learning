package main
import (
 "fmt"
 "math/rand"
 "time"
)
var balance int
func credit() {
     for i := 0; i < 10; i++ {
     balance += 100
     time.Sleep(time.Duration(rand.Intn(100)) *
     time.Millisecond)
     fmt.Println("After crediting, balance is",balance)
 }
}

func debit() {
	for i := 0; i < 5; i++ {
	balance -= 100
	time.Sleep(time.Duration(rand.Intn(100)) *
	time.Millisecond)
	fmt.Println("After debiting, balance is", balance)
	}
   }

func main() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
}
