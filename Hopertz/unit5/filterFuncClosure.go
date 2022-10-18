package main 

import ("fmt")

func filter(arr []int, cond func(int) bool) []int {
   result := []int{}
   for _, v := range arr {
   if cond(v) {
      result = append(result, v)
   }
   }
   return result
}

func main() {
  a := []int{1, 2, 3, 4, 5}
  evens := filter(a,
  func(val int) bool {
     return val%2 == 0
   })
   
 fmt.Println(evens)
}
