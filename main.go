package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// for unmarshalling json
type Contents struct {
	Title   string
	Story   []string
	Options []map[string]string
}

// struct for JSON of each arc including arc title, to be implemented in http handle
type Webpage struct {
	Arc     string
	Details Contents
}

// parses JSON story
func parseJSON() map[string]Contents {
	// open file as byte data
	data, err := ioutil.ReadFile("gopher.json")
	// exit if file not opened
	if err != nil {
		log.Fatal(err)
	}

	// gets keys as strings, values as type Contents
	var obj map[string]Contents

	err = json.Unmarshal(data, &obj)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func (data Webpage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	obj := parseJSON()
	introContent := obj["intro"]

	// http handler
	mux := http.NewServeMux()

	// handle 404 requests
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "No chapter found. Stop messing with the URL!")
	})

	introHandler := Webpage{"intro", introContent}
	mux.Handle("/cyoa", introHandler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
