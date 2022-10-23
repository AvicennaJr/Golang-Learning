package main 

import (
	"fmt"
	"sort"
)

type kv struct {
	key string
	value int
   }
type kvPairs []kv
var heights map[string]int

func (p kvPairs) Len() int {
	// returns the length of the collection
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	// indicates the first value (height) must be smaller
	// than the second value
	return p[i].value < p[j].value
}


func (p kvPairs) Swap(i, j int) {
	// swaps the items in the collection
	p[i], p[j] = p[j], p[i]
}

func main(){
	
    heights := make(map[string]int)
    heights["Peter"] = 170
    heights["Joan"] = 168
    heights["Jan"] = 175
    p := make(kvPairs, len(heights))
    i := 0
    for k, v := range heights {
        p[i] = kv{k, v}
        i++

  }
  sort.Sort(p)
  fmt.Println(p)

   
  
}

