package main

import (
	"fmt"
	"net/http"

	"github.com/mvaldes14/api-go-otel/pkg/api"
	"github.com/mvaldes14/api-go-otel/pkg/app"
)

func main() {
	mux := http.NewServeMux()

	// Application Handlers (HTMX)
	mux.Handle("GET /", app.IndexApp())

	// API Handlers
	mux.HandleFunc("POST /api/tasks", api.ApiIndexHandler)

	// API Handlers
	mux.HandleFunc("GET /api/tasks", api.ApiIndexHandler)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", mux)

}
