package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

var JSON = `{
	"name" : "Gopher",
	"title" : "programmer"
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(c)

	var m map[string]interface{}
	err = json.Unmarshal([]byte(JSON), &m)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(m)

	data, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}
