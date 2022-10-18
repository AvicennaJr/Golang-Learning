package main

import ("fmt")

func addNum(num1 int, num2 int) (sum int) {
 sum = num1 + num2
 return // you can still use "return sum"
}


func main(){
	s := addNum(4,3);
	fmt.Println(s) 

	

}
