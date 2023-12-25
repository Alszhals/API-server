package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, this is a JSON response!"}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
