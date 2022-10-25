package main 

import (
	"fmt"
)

type point struct {
	x float32
	y float32
	z float32
   }


func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
   }

func main(){
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4) // &{7.8 9.1 2.3}
	fmt.Println(pt4.x) // 7.8

	// Making a Copy of a Struct
	 pt5 := pt4

	 fmt.Println(pt5)

	 // To verify this, let’s change a field in pt5 and then point out the values for both pt4
     // and pt5:
    pt5.x = 0
    fmt.Println(pt4) // &{0 9.1 2.3}
    fmt.Println(pt5) // &{0 9.1 2.3}

	//The output shows that modifying pt5 affects the values for pt4 as well, because
    // they’re both pointing to the same struct.

	//If you want to create an independent copy of pt4, you need to use the * character,
    //like this:
     pt6 := *pt4

	 pt6.z = 0
     fmt.Println(pt4) // &{0 9.1 2.3}
     fmt.Println(pt6)
    }

