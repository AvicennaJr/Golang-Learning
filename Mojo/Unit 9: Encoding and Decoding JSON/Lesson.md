# Encoding and Decoding Data Using JSON

## Decoding JSON

### Decoding JSON to Struct

<pre>
// Lets say you have this JSON 

{
    "firstname": "Mojo",
    "lastname": "Jojo"
}

// Go code

package main

import {
    "encoding/json"
    "fmt"
}

type People struct {
    Firstname string
    Lastname  string
} // fields in struct must match the keys in json 

func main() {
    var person People

    jsonString := `{"firstname": "Mojo",
                    "lastname": "Jojo"}`

    err := json.Unmarshal([]byte(jsonString), &person) //pass to address of person
    if err == nil {
        fmt.Println(person.Firstname)
        fmt.Println(person.Lastname)
    } else {
        fmt.Println(err)
    }
}
</pre>

### Decoding JSON to Arrays

Usually JSON is like this:

<pre>
[
    {
        "firstname": "Mojo",
        "lastname": "Jojo"
    },
    {
        "firstname": "Hoper",
        "lastname": "TZ"
    }
]
</pre>

Then use this:

<pre>
func main() {
    var persons []People // Array that has type People
    json.Unmarshal([]bye(jsonString), &persons)

    for _, person := range persons {
        fmt.Println(person.Firstname)
        fmt.Println(person.Lastname)
    }
}
</pre>

### Decoding Embedded Objects

<pre>


package main
    import (
        "encoding/json"
        "fmt"
)
type People struct {
    Firstname string
    Lastname string
    Details struct {
        Height int
        Weight float32
    }
}
func main() {
    var persons []People
    jsonString :=
        `[
            {
                "firstname":"Wei-Meng",
                "lastname":"Lee",
                "details": {
                    "height":175,
                    "weight":70.0
                    }
            },
            {

                "firstname":"Mickey",
                "lastname":"Mouse",
                "details": {
                    "height":105,
                    "weight":85.5
                }
            }
        ]`
    json.Unmarshal([]byte(jsonString), &persons)
    
    for _ , person := range persons {
        fmt.Println(person.Firstname)
        fmt.Println(person.Lastname)
        fmt.Println(person.Details.Height)
        fmt.Println(person.Details.Weight)
    }
}

</pre>

### Mapping Custom Attribute Names

<pre>
{
    "base currency":"EUR",
    "destination currency":"USD"
}

// In go write

type Rates struct {
    Base string `json:"base currency"`
    Symbol string `json:"destination currency"`
}

</pre>


## Encoding JSON

### Encoding JSON to structs

<pre>

package main
import (
	"encoding/json"
	"fmt"
	"time"
)
type Name struct {
	FirstName string
	LastName string
}
type Address struct {
	Line1 string
	Line2 string
	Line3 string
}
type Customer struct {
	Name Name
	Email string
	Address Address
	DOB time.Time
}

func main() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	john := Customer{
	Name: Name{ FirstName: "John",
		LastName: "Smith",},
	Email: "johnsmith@example.com",
	Address: Address{
		Line1: "The White House",
		Line2: "1600 Pennsylvania Avenue NW",
		Line3: "Washington, DC 20500",
		},
		DOB: dob,
	}
	johnJSON, err := json.Marshal(john)
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}
</pre>

### Encoding Interfaces to JSON

<pre>

package main
import (
	"encoding/json"
	"fmt"
	"time"
)


type Customer map[string]interface{}
type Name map[string]interface{}
type Address map[string]interface{}
func main() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	john := Customer{
		"Name": Name{
			"FirstName": "John",
			"LastName": "Smith",
			},
		"Email": "johnsmith@example.com",
		"Address": Address{
			"Line1": "The White House",
			"Line2": "1600 Pennsylvania Avenue"Line3": "Washington, DC 20500",
			},
		"DOB": dob,
		}
	johnJSON, err := json.MarshalIndent(john, "", "
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}
</pre>

> That's all about JSON 
