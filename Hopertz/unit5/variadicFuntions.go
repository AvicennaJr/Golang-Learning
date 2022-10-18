package main

import ("fmt")
// A variadic function takes in a variable number of arguments

// You can also have a fixed parameter together with a variadic parameter:
// variadic parameter must always be the last parameter in the function
// eg func addNums(total int, nums ...int) int {}

func addNums(nums ... int) int {
   total := 0
   for _, n := range nums {
      total += n
    }
 return total
}


func main(){
	fmt.Println(addNums(1,2,3,4,5)) 
    fmt.Println(addNums(1,2,3))

	

}
