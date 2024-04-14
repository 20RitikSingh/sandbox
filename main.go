package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle the POST request here
	// You can access the request body using r.Body

	// Parse the request body
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Run the script and wait for it to complete
	script()

	jsonData, err := json.Marshal(map[string]interface{}{"message": "Script executed successfully"})
	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/run", runHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
