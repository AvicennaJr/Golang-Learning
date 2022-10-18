package main

import ("fmt";"time")

func displayDate(){
	fmt.Println(time.Now().Date())
}

func main(){
	displayDate();
}
