package main
import (

 "encoding/json"
 "fmt"

)

type People struct {
	Firstname string
	Lastname string
   }

func main() {

	var person People
    jsonString := `{"firstname":"Wei-Meng","lastname":"Lee"}`
	err := json.Unmarshal([]byte(jsonString), &person)
	// If the unmarshalling (decoding) works, the json.Unmarshal() function returns a nil, 
    if err == nil {
      fmt.Println(person.Firstname)
      fmt.Println(person.Lastname)
    } else {
       fmt.Println(err)
    }

	var persons []People
	jsonStrings :=
	`[
	  {
          "firstname":"Wei-Meng",
	      "lastname":"Lee"
	  },
	  {
	     "firstname":"Mickey",
	     "lastname":"Mouse"
	  }
	 ]`

	json.Unmarshal([]byte(jsonStrings), &persons)
	for _, person := range persons {
	    fmt.Println(person.Firstname)
	    fmt.Println(person.Lastname)
	}


}
