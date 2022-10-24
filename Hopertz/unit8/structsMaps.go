package main
import (
	"fmt"
	"sort"
)


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

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	   }

	// get all the keys in members
    var keys []int
    for k := range members {
     keys = append(keys, k)
    }
    // sort the keys in ascending order
    sort.Ints(keys)
    // copy the value in members to the slice
    var sliceMembers []people
    for _, k := range keys {
     sliceMembers = append(sliceMembers, members[k])
    }

	fmt.Println(sliceMembers)

	//To sort a slice, you can use the SliceStable() function from the sort package.
    // The SliceStable() function has the following signature:

	//func SliceStable(slice interface{},
    //      less func(i, j int) bool)

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		// compare year
		if sliceMembers[i].dob.year != sliceMembers[j].dob.year {
		    return sliceMembers[i].dob.year < sliceMembers[j].dob.year
		}
		// compare month
		if sliceMembers[i].dob.month !=sliceMembers[j].dob.month {
		   return sliceMembers[i].dob.month < sliceMembers[j].dob.month
		}
		// compare day
		return sliceMembers[i].dob.day <
		sliceMembers[j].dob.day
	   })
	   for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	   }

	   //What if you just want to sort the members based on their year? In this case, you
       //simply need to perform the comparison for the year in the SliceStable()
       //function:

	   sort.SliceStable(sliceMembers, func(i, j int) bool {
           return sliceMembers[i].dob.year < sliceMembers[j].dob.year})

		for _, v := range sliceMembers {
			fmt.Println(v.name, v.email, v.dob)
		   }
	   // The SliceStable() function has a cousin function that is very similar: Slice().
       // Both these two functions allow you to sort a slice, but the Slice() function
       // doesn’t guarantee the sort result to be stable.

	   // If the order of the original elements is important, you should always use the
       // SliceStable() function for sorting slices.
}
