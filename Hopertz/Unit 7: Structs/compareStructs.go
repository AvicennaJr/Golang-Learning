package main 

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type point struct {
	x float32
	y float32
	z float32
   }

 type points struct {
	x float32
	y float32
	z float32
	name []string
   }

func main(){

	pt1 := point{x: 5.6, y: 3.8, z: 6.9}
	pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	pt3 := point{x: 6.5, y: 3.8, z: 6.9}
	fmt.Println(pt1 == pt2) // true
	fmt.Println(pt2 == pt3) // false
	pt4 := points{x: 5.6, y: 3.8, z: 6.9,
		name: []string{"pt1"}}

	pt5 := points{x: 5.6, y: 3.8, z: 6.9,
		name: []string{"pt2"}}

    // invalid operation: pt4 == pt5 (struct containing
    // []string cannot be compared)
    //fmt.Println(pt4 == pt5)
	
	//You can’t directly compare structs that contain fields that aren’t comparable, but
    //you can use the cmp package (https://pkg.go.dev/github.com/google/go-cmp/
	//cmp) to do that
	
	fmt.Println(cmp.Equal(pt4, pt5,cmp.AllowUnexported(points{}))) // false
    fmt.Println(cmp.Equal(pt1, pt2,cmp.AllowUnexported(point{}))) // true

}

