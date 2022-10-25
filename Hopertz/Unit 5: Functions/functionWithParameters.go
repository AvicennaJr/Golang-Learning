package main

import ("fmt";"time")

func displayDate(format string){
	fmt.Println(time.Now().Format(format))
}

func main(){
	displayDate("Mon 2006-01-02 15:04:56");
	//displayDate("15:04:05, 2006-Jan02 Mon"
	//displayDate("15:04:05, 2006-Jan-02 Monday")
	//displayDate("15:04:05, 2006-January-02 MST Mon")
	//displayDate("3:4:05 pm, 2006-1-02 MST Mon")
}
