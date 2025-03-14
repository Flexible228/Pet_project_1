package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

var task string

func postHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body RequestBody
	err := decoder.Decode(body)
	if err != nil {
		panic(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello, ", task)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)
	http.ListenAndServe("localhost:8080", nil)
}
