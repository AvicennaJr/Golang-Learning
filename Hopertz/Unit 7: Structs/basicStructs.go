package main 

import (
	"fmt"
)

type point struct {
	x float32
	y float32
	z float32
   }

func main(){

	var pt1 point
    pt1.x = 3.1
    pt1.y = 5.7
    pt1.z = 4.2

	fmt.Println(pt1.x,pt1.y,pt1.z)

	// another ways to create and initialize a struct

	pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	fmt.Println(pt2)

	pt3 := point{5.6, 3.8, 6.9}
	fmt.Println(pt3)

	// split the struct literal over multiple lines,

	pt4 := point{
		x: 5.6,
		y: 3.8,
		z: 6.9, // <-- comma here
	   }
	fmt.Println(pt4)
    //You can also leave out a specific field when initializing the struct:

	pt5 := point{x: 5.6, y: 3.8} //Fields that are omitted during initialization will be zero-based:
	fmt.Println(pt5) // {5.6 3.8 0}
}

