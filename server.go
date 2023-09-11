package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// post request
func postHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Process the data
	// You can access data["key"] to get values from the JSON
	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Name = %s", data["name"])
}

func main() {
	http.HandleFunc("/hello", helloHandler) //Get request handler
	http.HandleFunc("/post", postHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
