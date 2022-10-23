package main
import (
 "encoding/json"
 "fmt"
)

 // Sometimes the keys in your JSON string can’t be mapped directly to members of
 // your struct in Go eg keys in this JSON string  can not have spaces in them.
 // consider the following
 //{
	//"base currency":"EUR",
	//"destination currency":"USD"
   //}

type Rates struct {
 Base string `json:"base currency"`
 Symbol string `json:"destination currency"`
}
func main() {
 jsonString :=`{
	"base currency":"EUR",
	"destination currency":"USD"
	}`
	var rates Rates
	json.Unmarshal([]byte(jsonString), &rates)
	fmt.Println(rates.Base) // EUR
	fmt.Println(rates.Symbol) // USD
}
