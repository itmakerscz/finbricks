package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func ChatBot(w http.ResponseWriter, r *http.Request) {
	//
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	w.Header().Set("Content-Type", "application/json")
	//
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(b)
}

func main() {

	http.HandleFunc("/transcribe", Transcribe)
	http.HandleFunc("/chatbot", ChatBot)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Transcribe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Transcribe %s!", r.URL.Path[1:])
}
