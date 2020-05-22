package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"encoding/json"
	"github.com/jpeeling13/help-me-api/platform/requestFeed"
)

func main() {
	tds := requestFeed.New()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to the Help Me API Server"))
	})

	r.Get("/requests", func(w http.ResponseWriter, r *http.Request){
		requests := tds.GetAll()
		json.NewEncoder(w).Encode(requests)
	})

	r.Post("/requests", func(w http.ResponseWriter, r *http.Request){
		newRequest := map[string]string{}
		json.NewDecoder(r.Body).Decode(&newRequest)
		tds.Add(requestFeed.Request{
			newRequest["Description"],
			newRequest["RequesterName"],
		})
		w.Write([]byte("Thanks for the new request!"))
	})

	fmt.Println("Starting help-me-api server")
	http.ListenAndServe(":3000", r)
}
