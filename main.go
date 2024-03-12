// Package main
package main

import (
	"fmt"
	"net/http"

	"github.com/mvaldes14/api-go-otel/pkg/api"
	"github.com/mvaldes14/api-go-otel/pkg/app"
	"github.com/mvaldes14/api-go-otel/pkg/database"
)

func main() {

	// Initialize the database
	database.InitDb()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Application Handlers (HTMX)
	mux.HandleFunc("GET /", app.IndexApp)

	// Application Handlers (HTMX)
	mux.HandleFunc("GET /twitch", app.TwitchApp)

	// // Application Handlers (HTMX)
	mux.HandleFunc("GET /local", app.LocalApp)

	// API Handlers
	mux.HandleFunc("GET /api/health", api.IndexHandler)

	// API Handlers
	mux.HandleFunc("POST /api/todo", api.AddTodoHandler)

	// API Handlers
	mux.HandleFunc("GET /api/todo", api.GetTodoHandler)

	// API Handlers
	mux.HandleFunc("DELETE /api/todo/{id}", api.DeleteTodoHandler)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", mux)
}
