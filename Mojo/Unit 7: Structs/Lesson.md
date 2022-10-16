## Defining Structs

<pre>

package main

import "fmt"

type point struct {
    x float32
    y float32
    z float32
}

func main() {
    var pt1 point // create a variable pt1 of type point

    pt1.x = 3.1
    pt1.y = 5.7
    pt1.z = 4.2

    pt2 := point(5.6,7.3,3.4)
}

fmt.Println(pt1.x)
fmt.Println(pt1.y)
fmt.Println(pt1.z)

</pre>

## Creating a Go Struct

<pre>

func newPoint(x, y, z float32) `*point` {
    p := point(x:x, y:y, z:z)
    return &p
}

pt3 = newPoint(7.5,8.3,2.1)
fmt.Println(pt3) // &(7.5 8.3 2.1) the & indicates a pointer

</pre>

## Copying a Struct

Using y := x will make both x and y point to the same struct. Meaning if x
changes then y also changes.

To create and independent copy use `y := *x` thus changes of one won't affect the other.


## Defining Methods in Structs

<pre>
package main

import (
    "fmt"
    "math"
)

type point struct {
    x float32
    y float32
    z float32
}

func (p point) length() float64 {
    return math.Sqrt(
        (math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2) +
        math.Pow(float64(p.z), 2)))
}

func newPoint(x, y, z float32) `*point` {
    p := point(x:x, y:y, z:z)
    return &p
}

func main() {
    pt4 := newPoint(7.8, 4.4, 3.2)
    fmt.Println(pt4.length())
}
</pre>

## Comparing Structs

You can compare two structs to see if they’re equal, provided all the fields inside
the struct are comparable. 

<pre>

pt1 := point{x: 5.6, y: 3.8, z: 6.9}
pt2 := point{x: 5.6, y: 3.8, z: 6.9}
pt3 := point{x: 6.5, y: 3.8, z: 6.9}
fmt.Println(pt1 == pt2) // true
fmt.Println(pt2 == pt3) // false
</pre>

> And done on structs
