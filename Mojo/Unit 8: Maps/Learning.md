# Maps

Dictionaries in python

## Defining a Map

`map[keyType] valueType`

<pre>

package main

var heights map[string]int

func main() {
}
</pre>

Because the map type is a reference type, you need to first initialize it using the
make() function before you can use it:

`heights = make(map[string]int)`

Map values as such:

<pre>
heights["Joan"] = 168
heights["Jan"] = 175

// To retrieve

fmt.Println(heights["Joan"]) // 168
</pre>

## Basic operations

### Initializing a map with a map literal

Alternatively you can initialize a map like this:

<pre>
heights := map[string]int {
    "Joan": 170,
    "Jan": 175, // <--- Notice the coma
}
</pre>

### Checking the existence of a key

Using a key that doesn't a key will return a zero. A beter way is to implement
something like:

<pre>

if v, ok := heights["Jim"]; ok {
    fmt.Println(v)
} else {
    fmt.Println("Key doesn't exist")
}
</pre>

### Deleting a Key

Use `delete(map, key)`

### Getting the number of items in a map

`fmt.Println(len(heights))`

### Iterating over a map

<pre>
for k, v := range heights {
    fmt.Println(k, v)
}
</pre>

### Getting all keys in a map

Need to implement your own functions *like most things in go*

<pre>

var key []string

for k := range heights {
    keys = append(keys, k)
}

</pre>

### Setting the iteration order

<pre>
import (
    "fmt"
    "sort"
)

...

sort.Strings(keys)

for _ , k := range keys {
    fmt.Println(k, heights[k])
}

</pre>

### Sorting the items by values

Not straight forward

<pre>
package main

import (
    "fmt"
    "sort"
)

// Need to first define struct and a slice

type kv struct {
    key string
    val int
}

type kvPairs []kv

var heights map[string]int

// Need to implement interfaces - *well fuck me :(*


func (p kvPairs) Len() int {
    // returns the length of the collection
    return len(p)
}

func (p kvPairs) Less(i, j int) bool {
    // indicates the first value (height) must be smaller
    // than the second value
    return p[i].value < p[j].value
}

// Really wondering if go is worth rn


func (p kvPairs) Swap(i, j int) {
    // swaps the items in the collection
    p[i], p[j] = p[j], p[i]
}


func main() {
    heights := map(map[string]int)

    heights := make(map[string]int)
    heights["Peter"] = 170
    heights["Joan"] = 168
    heights["Jan"] = 175
    
    p := make(kvPairs, len(heights))

    i:= 0

    for k, v := range heights {
        p[i] = kv[k, v]
        i++
    }

## Struct and Maps

### Creating a map of structs

Suppose you have a struct of people and their dates of birth

<pre>

type dob struct {
    day int
    month int
    year int
}

type people struct {
    name string
    email string
    dob dob
}

// To store a collection of people struct, you can declare a map of int key and value
of type people:

var members map[int]people

func main() {
    members = make(map[int]people)
    members[1] = people{
        name: "Mary Smith",
        email: "marysmith@example.com",
        dob: dob{
            day: 17,
            month: 3,
            year: 1990,
        },
    }
    
    members[2] = people{
        name: "John Smith",
        email: "johnsmith@example.com",
        dob: dob{
            day: 9,
            month: 12,
            year: 1988,
        },
    }
    members[3] = people{
        name: "Janet Doe",
        email: "janetdoe@example.com",
        dob: dob{
            day: 1,
            month: 12,
            year: 1988,
        },
    }
    members[4] = people{
        name: "Adam Jones",
        email: "adamjones@example.com",
        dob: dob{
            day: 19,
            month: 8,
            year: 2001,
        },
    }

    // To print out

    for k, v := range members {
        fmt.Println(k, v.name, v.email, v.dob)
    }

}

</pre>

> To do: Custom sorting of map of structs
