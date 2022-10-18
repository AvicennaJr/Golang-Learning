package main

import ("fmt")
func fib() func() int {
 f1 := 0
 f2 := 1
 return func() int {
 f1, f2 = f2, (f1 + f2)
  return f1
 }
}

func main() {
 gen := fib()
 fmt.Println(gen()) // 1
 fmt.Println(gen()) // 1
 for i := 0; i < 10; i++ {
    fmt.Println(gen())
 }
}
