package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apis map[int]string

// channel to store results
var c chan map[int]interface{}

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			var re = make(map[int]interface{})
			switch API {
			case 1:
				if result["error"] == false {
					re[API] = result["joke"]
				} else {
					fmt.Println(result["error"])
				}
				c <- re
				fmt.Println("Result for API 1 stored")
			case 2:
				if result["error"] == false {
					setup := result["setup"].(string)
					delivery := result["delivery"].(string)
					re[API] = fmt.Sprintf(setup, delivery)
				} else {
					fmt.Print(result["error"])
				}
				c <- re
				fmt.Println("Result for API 2 stored")
			}
		} else {
			log.Fatal(err)
		}
	}
}

func main() {
	c = make(chan map[int]interface{})

	apis = make(map[int]string)

	apis[1] = "https://v2.jokeapi.dev/joke/any?type=single"
	apis[2] = "https://v2.jokeapi.dev/joke/Any?type=twopart"

	go fetchData(1)
	go fetchData(2)

	for i := 0; i < 2; i++ {
		fmt.Println(fmt.Sprint(<-c))
	}

	fmt.Println("Done")
	fmt.Scanln()
}
