package main

import (
	"encoding/json"
	"fmt"
)

type Bio struct {
	Height int `json:"Height"`
	Weight int `json:"Weight"`
}

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Bio  Bio    `json:"Bio"`
}

func main() {

	// JSON to struct
	jsonRawbytes := []byte(`{"Name": "Alps", "Age": 23, "Bio": {"Height": 167, "Weight": 62}}`)

	var person Person

	err := json.Unmarshal(jsonRawbytes, &person)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", person)

	// Struct to JSON

	data, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
