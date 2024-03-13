// Package main
package main

import (
	"fmt"
	"net/http"

	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/mvaldes14/api-go-otel/pkg/api"
	"github.com/mvaldes14/api-go-otel/pkg/app"
	"github.com/mvaldes14/api-go-otel/pkg/database"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	// Application Handlers (HTMX)
	handleFunc("GET /", app.IndexApp)

	// Application Handlers (HTMX)
	handleFunc("GET /twitch", app.TwitchApp)

	// Application Handlers (HTMX)
	handleFunc("GET /local", app.LocalApp)

	// API Handlers
	handleFunc("GET /api/health", api.IndexHandler)

	// API Handlers
	handleFunc("POST /api/todo", api.AddTodoHandler)

	// API Handlers
	handleFunc("GET /api/todo", api.GetTodoHandler)

	// API Handlers
	handleFunc("DELETE /api/todo/{id}", api.DeleteTodoHandler)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}

func main() {
	// Initialize the database
	database.InitDb()

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":3000",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		fmt.Println("Server is running on port 3000")
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}
