package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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

	// fmt.Println(obj)

	// http handler
	mux := http.NewServeMux()
	testHandler := demo{"working from main"}
	mux.Handle("/test", testHandler)

	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type demo struct {
	Example string
}

func (d demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	err = parsedTemplate.Execute(w, d)
	if err != nil {
		log.Fatal(err)
	}
}
