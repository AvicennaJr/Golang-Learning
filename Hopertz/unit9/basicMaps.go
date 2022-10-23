package main 

import (
	"fmt"
	"sort"
)


// A Go map type has the following syntax:
// map[keyType] valueType

var heights map[string]int

func main(){

	//Because the map type is a reference type, you need to first initialize it using the
    //make() function before you can use it:
	heights = make(map[string]int)

	heights["Peter"] = 170
    heights["Joan"] = 168    
	heights["Jan"] = 175

	fmt.Println(heights["Peter"]) // 170
    fmt.Println(heights["Joan"]) // 168
    fmt.Println(heights["Jan"]) // 175

	// Initializing a map with a map literal
	newheights := map[string]int{
		"Peter": 170,
		"Joan": 168,
		"Jan": 175, // <-- note the comma here
	   }
    fmt.Println(newheights)

	// Checking the existence of a key
	fmt.Println(heights["Jim"]) // 0

	// Deleting a key
	// syntax delete(map, key)
	if _, ok := heights["Joan"]; ok {
		delete(heights, "Joan")
	   } else {
		fmt.Println("Key does not exist")
	}

	// Getting the number of items in a map
	fmt.Println(len(heights))

	// Iterating over a map
	for k, v := range heights {
		fmt.Println(k, v)
	   }


	// Getting all the keys in a map
	var keys []string
    for k := range heights {
     keys = append(keys, k)
    }
    fmt.Println(keys) // [Jan Peter ]

	// Setting the iteration order in a map
	sort.Strings(keys)
    fmt.Println(keys) // [Jan Peter]
    
	// After the keys are sorted, you can then use the for-range loop to iterate over the
    //keys and print out the value of each key:
	for _, k := range keys {
		fmt.Println(k, heights[k])
	   }



 
}

