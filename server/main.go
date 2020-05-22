package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"encoding/json"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to the Help API Server"))
	})

	r.Get("/requests", func(w http.ResponseWriter, r *http.Request){

		json.NewEncoder(w)

	})

	http.ListenAndServe(":3000", r)
}
