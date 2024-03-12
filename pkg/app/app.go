// Package app provides the HTTP handlers for the application endpoints
package app

import (
	"context"
	"net/http"

	"github.com/mvaldes14/api-go-otel/pkg/database"
	views "github.com/mvaldes14/api-go-otel/views/layouts"
)

// IndexApp returns the main todo index page
func IndexApp(w http.ResponseWriter, _ *http.Request) {
	todos, err := database.GetTodos()
	if err != nil {
		return
	}
	todoLayout := views.TodoLayout(todos)
	views.BaseLayout(todoLayout).Render(context.Background(), w)
}

// TwitchApp returns the twitch app
func TwitchApp(w http.ResponseWriter, _ *http.Request) {
	twitchLayout := views.TwitchLayout()
	views.BaseLayout(twitchLayout).Render(context.Background(), w)
}

// LocalApp returns the local app
func LocalApp(w http.ResponseWriter, _ *http.Request) {
	localLayout := views.LocalLayout()
	views.BaseLayout(localLayout).Render(context.Background(), w)
}
