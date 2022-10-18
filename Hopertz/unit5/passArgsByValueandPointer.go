package main

import ("fmt")

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
  x := 5
  y := 6
  swap(&x, &y)
  fmt.Println(x, y) // 6 5
}
