package main

import "fmt"

func main(){
	num := 5
    dayOfWeek := ""
    switch num {
       case 1:
         dayOfWeek = "Monday"
       case 2:
         dayOfWeek = "Tuesday"
       case 3:
          dayOfWeek = "Wednesday"
       case 4:
          dayOfWeek = "Thursday"
       case 5:
          dayOfWeek = "Friday"
       case 6:
           dayOfWeek = "Saturday"
       case 7:
            dayOfWeek = "Sunday"
       default:
           dayOfWeek = "--error--"
    }
  fmt.Println(dayOfWeek) // Friday
}
