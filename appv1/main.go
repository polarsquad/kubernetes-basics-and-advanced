package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "This page intentionally left blank. Version 1."
		jsonResp, err := json.Marshal(resp)
		HandleError(err)
		w.Write(jsonResp)
		log.Printf("blank called\n")

	})

	fmt.Printf("Server running (port=8080)\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
