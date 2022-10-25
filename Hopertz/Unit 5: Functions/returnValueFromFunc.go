package main

import ("fmt")

func addNum(num1, num2 int) int { // returns int
 return num1 + num2
}

func countOddEven(s string) (int,int) {
   odds, evens := 0,0
   for _, c := range s {
   if int(c) % 2 == 0 {
      evens++
    } else {
      odds++
    }
   }
   return odds,evens
  }

func main(){
	s := addNum(4,3);
	fmt.Println(s) 

	o, e := countOddEven("12345")
    fmt.Println(o,e) // 3 2

	_, f := countOddEven("12345")
	fmt.Println(f) 
}
