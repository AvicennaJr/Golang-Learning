## Arrays

Numbered sequence of items of the same type with a specific length

*Note that the size of an array can't change once deined*

### Declaring an Array

`var array_name [size_of_array]datatype`

Example:

<pre>
package main

import "fmt"

func main() {
    var num[5]int // an array of 5 int elements
    fmt.Println(nums) // [0,0,0,0,0]
}
</pre>

### Initializing an Array

<pre>
nums := [5]int[1,2,3,4,5] 
fmt.Println(nums)
</pre>

This is also allowed:

`nums := [...]int[1,2,3]`

### Multidimensional Arrays

<pre>
var table [5][6]int

for row := 0; row<5; row++ {
    for column := 0; column<6; column++ {
        table[row][column] = row * column
    }
}

fmt.Println(table)
</pre>

## Slices

Slices are arrays that can grow and shrink in size

### Creating an Empty slice

`s := make([]int, 5)`

Slices have a header which have details of the ponter that points to the array,
the length of the array and a capacity that shows the maximum allowed numbers
in an array

<pre>
r := make([]int, 2,5) // array has a length of 2 and max capacity of 5
fmt.Println(len(r)) //2
fmt.Println(cap(r)) //5
</pre>

### Creating and Initializing a Slice

<pre>
t := []int{1,2,3,4,5}

fmt.Println(len(t)) // 5
fmt.Println(cap(t)) // 5

</pre>

### Appending to a Slice

<pre>

t = append(t, 6,7,8)

fmt.Println(t) // 1,2,3,4,5,6,7,8

fmt.Println(len(t)) //8
fmt.Println(cap(t)) //10

</pre>

For efficiency reasons, the capacity will be double of the initial array.
Shouldn't be something to worry about though.

Consider this:

<pre>
u := t // both u and t will be pointing to the same array

// if you change the value of an element in one array:

u[7] = 100

// it will change for both u and t.

// However if you append values to one array that exceed the capacity:

t = append(t, 9, 10, 11)

// The capacity of t will change from 10 to 20 and it will point to a new array
// Thus u and t will be different

</pre>

## Slicing and Ranging

### Extracting parts of an array or slice

Basically like in python

`c[0:2] or c[:2] // first and second element`

`c[3:] //from 3 to the end`
`c[:] // the whole thing`

*Note that the result of slicing an array or slice is a slice*

### Iterating through a slice

<pre>
for i, v := range t {
    fmt.Println(i, v)
}
</pre>

### Making copies of an array or slice

For arrays (unlike slices) simply:
<pre>
arr2 := arr1
</pre>

Changes in one won't affect the other. For slices use this:

<pre>
t := []int[1,2,3,4,5]

v := make([]int, len(t)) // length must be the same as in t

// use copy function to copy 

copy(v, t)
</pre>

Now changes in one won't affect the other. 

*Note that the copy() function examines the length of both the destination and
source slices and copies the minimum of these two numbers of elements*

### Inserting items into a slice

You need to create your own `insert` function

<pre>
func insert(orig []int, index int, value int) // accepts arr and two ints
    ([]int, error) { // returns an int array
    if index < 0 { // No index less than zero, return error
        return nil, errors.New(
            "Index cannot be less than 0")
    }
    if index >= len(orig) { // index is beyond the len of arr thn append
        return append(orig, value), nil
    }

    // if not the two above cases then insert in the array

    orig = append(orig[:index], orig[index:]...) // the ... is because append-
// is variadic
    orig[index] = value
    return orig, nil
}
</pre>


### Removing an item from a slice

Create a custom function *fucking golang :(*

<pre>

func delete(orig []int, index int) ([]int, error) {
    if index < 0 || index >= len(orig) {
        return nil, errors.New("Index out of range")
    }

    orig = append(orig[:index], orig[index+1:]...)
    return orig, nil
}
</pre>

 > And we're done on arrays and slices
