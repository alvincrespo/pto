package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":"+port, nil)
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
