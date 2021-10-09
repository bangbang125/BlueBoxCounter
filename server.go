package main

import (
	"log"
	"net/http"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func main() {
	// register hander
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
