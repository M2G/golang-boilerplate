package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func BarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func BarByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	id := r.PathValue("id")
	fmt.Fprintf(w, "Item ID: %s", id)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
