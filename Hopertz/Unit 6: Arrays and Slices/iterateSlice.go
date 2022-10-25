package main

import (
 "fmt"
)
func main() {
	t := []int{1, 2, 3, 4, 5}
	for i, v := range t {
		fmt.Println(i, v)
	   }
}
