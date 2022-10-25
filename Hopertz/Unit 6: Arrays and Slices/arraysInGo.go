package main

import (
 "fmt"
)
func main() {
  var nums [5]int // an array of int (5 elements)
  fmt.Println(nums) // [0 0 0 0 0]
  //Go are zero-based
  fmt.Println(nums[0]) // first element
  fmt.Println(nums[1]) // second element

  // initialize  an array with some initial values
  anotherNums := [5]int{1, 2, 3, 4, 5}
  fmt.Println(anotherNums)

  // you can omit the length of the size by using the ... notation:
  omitlennums := [...]int{1, 2, 3, 4, 5}
  fmt.Println(len( omitlennums)) // 5
}
