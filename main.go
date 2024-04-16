// Package main
package main

import (
	"log"
	"net/http"

	"github.com/mvaldes14/api-go-otel/pkg/api"
	"github.com/mvaldes14/api-go-otel/pkg/app"
	"github.com/mvaldes14/api-go-otel/pkg/database"
)

func main() {
	// Initialize the database
	database.InitDb()

	mux := http.NewServeMux()

	// Register handlers.
	// Application Handlers (HTMX)
	mux.HandleFunc("GET /", app.IndexApp)

	// Application Handlers (HTMX)
	mux.HandleFunc("GET /twitch", app.TwitchApp)

	// Application Handlers (HTMX)
	mux.HandleFunc("GET /local", app.LocalApp)

	// API Handlers
	mux.HandleFunc("GET /api/health", api.IndexHandler)

	// API Handlers
	mux.HandleFunc("POST /api/todo", api.AddTodoHandler)

	// API Handlers
	mux.HandleFunc("GET /api/todo", api.GetTodoHandler)

	// API Handlers
	mux.HandleFunc("DELETE /api/todo/{id}", api.DeleteTodoHandler)

	log.Print("Listen Running on :3000")
	http.ListenAndServe(":3000", mux)
}
