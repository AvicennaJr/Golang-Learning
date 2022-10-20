package main

import (
 "fmt"
 "errors"
)

// Here is the delete() function to delete an item from a slice:
  func delete(orig []int, index int) ([]int, error) {
     if index < 0 || index >= len(orig) {
     return nil, errors.New("Index out of range")
     }
     orig = append(orig[:index], orig[index+1:]...)
     return orig, nil
}
func main() {
	//You can now test the delete() function:
	t := []int{1, 2, 3, 4, 5}
	t, err := delete(t, 2)
	if err == nil {
	 fmt.Println(t) // [1 2 4 5]
	} else {
	 fmt.Println(err)
	}
}
