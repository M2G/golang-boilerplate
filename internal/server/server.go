package server

import (
  "golang-boilerplate/internal/handlers"
  "net/http"
)

func RegisterRoutes() {
  http.HandleFunc("/bar", handlers.BarHandler)
  http.HandleFunc("/bar/{id}", handlers.BarByIDHandler)
  http.HandleFunc("/health", handlers.HealthHandler)
}
