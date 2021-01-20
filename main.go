package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"os"
	//"strings"
)

// for unmarshalling json
type Contents struct {
	Title   string
	Story   []string
	Options []map[string]string
}

func main() {
	// open file as byte data
	data, err := ioutil.ReadFile("gopher.json")
	// exit if file not opened
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(data)

	// gets keys as strings, values as interfaces which need to be unmarshalled further
	var obj map[string]Contents

	err = json.Unmarshal(data, &obj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj)
}
