package app

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/mvaldes14/api-go-otel/pkg/database"
	views "github.com/mvaldes14/api-go-otel/views/layouts"
)

func IndexApp() http.Handler {
	todos, err := database.GetTodos()
	if err != nil {
		return nil
	}
	return templ.Handler(views.BaseLayout(todos))
}
