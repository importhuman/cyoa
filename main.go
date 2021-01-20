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
// After completion of project: didn't really need this, could've done with Contents
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

	// to get arc names for handlers
	// for key := range obj {
	// 	fmt.Println(key)
	// }
	return obj
}

// for http.Handlers
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
	// assign result of parseJSON to obj
	obj := parseJSON()
	// introContent := obj["intro"]

	// http handler/multiplexer
	mux := http.NewServeMux()

	// handle 404 requests
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "No chapter found. Stop messing with the URL!")
	})

	// handlers for different pages
	intro := Webpage{"intro", obj["intro"]}
	home := Webpage{"home", obj["home"]}
	newyork := Webpage{"new-york", obj["new-york"]}
	debate := Webpage{"debate", obj["debate"]}
	seankelly := Webpage{"sean-kelly", obj["sean-kelly"]}
	markbates := Webpage{"mark-bates", obj["mark-bates"]}
	denver := Webpage{"denver", obj["denver"]}

	mux.Handle("/cyoa/", intro)
	mux.Handle("/cyoa/home", home)
	mux.Handle("/cyoa/new-york", newyork)
	mux.Handle("/cyoa/debate", debate)
	mux.Handle("/cyoa/sean-kelly", seankelly)
	mux.Handle("/cyoa/mark-bates", markbates)
	mux.Handle("/cyoa/denver", denver)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
