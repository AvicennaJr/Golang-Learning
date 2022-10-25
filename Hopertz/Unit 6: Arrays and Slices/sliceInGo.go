package main

import (
 "fmt"
)
func main() {
   s := make([]int, 5)
   fmt.Println(s)
   //slice is more flexible — the size of a slice can change as you append items to or remove items from it.

  // To find out the length and capacity of the slice s, you use the len() and cap() functions, respectively:
  fmt.Println(len(s)) // 5
  fmt.Println(cap(s)) // 5

  // you can create a slice with a specific length and capacity

  r := make([]int, 2, 5)
  fmt.Println(len(r)) // 2
  fmt.Println(cap(r)) // 5

  //Creating and initializing a slice

  t := []int{1, 2, 3, 4, 5}
  fmt.Println(len(t)) // 5
  fmt.Println(cap(t)) // 5

  // Appending to a slice

  t = append(t, 6, 7, 8)
  fmt.Println(t)// [1 2 3 4 5 6 7

  fmt.Println(len(t)) // 8
  fmt.Println(cap(t)) // 10  now it has  a bigger capacity

  //You can now append two more items to t without causing the underlying array to change
  t = append(t, 9, 10)
  fmt.Println(len(t)) // 10
  fmt.Println(cap(t)) // 10

 // Let’s now assign the slice t to another variable named u:
 u := t
 fmt.Println(u) // [1 2 3 4 5 6 7 8 9 10]
 fmt.Println(t) // [1 2 3 4 5 6 7 8 9 10]

 // both u and t are pointing to the same underlying array

 // to prove see below
 u[9] = 100
 fmt.Println(u) // [1 2 3 4 5 6 7 8 9 100]
 fmt.Println(t) // [1 2 3 4 5 6 7 8 9 100]

 // Now, let’s add a new item to the slice t and print out the values of both u and t:

 t = append(t, 11)
 fmt.Println(u) // [1 2 3 4 5 6 7 8 9 100]
 fmt.Println(t) // [1 2 3 4 5 6 7 8 9 100 11]

 // This is because when you added the new item to t, it exceeded its capacity, so a new underlying 
 //array was created to accommodate it

 // t and u are now pointing to two different underlying arrays. You
 // can verify this by printing the length and capacity of u and t:

 fmt.Println(len(u)) // 10
 fmt.Println(cap(u)) // 10
 fmt.Println(len(t)) // 11
 fmt.Println(cap(t)) // 20


}
