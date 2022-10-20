package main

import (
 "fmt"
)
func main() {

	//For arrays, making a copy is very straightforward:
	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := nums1
	fmt.Println(nums1) // 1 2 3 4 5]
    fmt.Println(nums2) // 1 2 3 4 5]

	// If you now make some changes to nums2, the changes should only affect nums2 and not nums1
	nums2[0] = 0
    fmt.Println(nums1) // 1 2 3 4 5]
    fmt.Println(nums2) // 0 2 3 4 5]


    // making copy for slice

	// To create copy of slices, use the copy() function
	// syntax copy(destination, source)

	t := []int{1, 2, 3, 4, 5}

	v := make([]int, len(t))

	copy(v, t)
    fmt.Println(t) // [1 2 3 4 5]
    fmt.Println(v) // [1 2 3 4 5]

	//Note that the copy() function examines the length of both the destination and
    //source slices and copies the minimum of these two numbers of elements

	s := []int{1, 2, 3, 4, 5}
    u := make([]int, 2, 5)

	copy(u, s)
    fmt.Println(s) // [1 2 3 4 5]
    fmt.Println(u) // [1 2]

	// On the other hand, if v now has a length of 10, like this:
	w := make([]int, 10)
	copy(w, t)
    fmt.Println(t) // [1 2 3 4 5]
    fmt.Println(w) // [1 2 3 4 5 0 0 0 0 0]


}
