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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var reqBody RequestBody
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	task = reqBody.Task
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task resived: %s", task)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Helo, %s", task)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)

	http.ListenAndServe("localhost:8080", nil)
}
