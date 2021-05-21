package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

type Profile struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling API...")

	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
