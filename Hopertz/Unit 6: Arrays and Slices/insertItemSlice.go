package main

import (
 "fmt"
 "errors"
)


	//Go doesn’t have built-in functions for inserting items into a slice. To do this, you
	//need to implement it yourself using the append() function. Let’s define a function
	//called insert()
    //func insert(orig []int, index int, value int)([]int, error) {}

    func insert(orig []int, index int, value int)([]int, error) {
    	if index < 0 {
    	return nil, errors.New(
    	"Index cannot be less than 0")
    	}
    	if index >= len(orig) {
    	return append(orig, value), nil
    	}
    	orig = append(orig[:index+1], orig[index:]...)
    	orig[index] = value
    	return orig, nil
    }

func main() {


    t := []int{1, 2, 3, 4, 5}
    t, err := insert(t, 2, 9)
    if err == nil {
     fmt.Println(t) // 1 2 9 3 4 5]
    } else {
     fmt.Println(err)
    }
    }

