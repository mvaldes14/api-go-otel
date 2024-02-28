package app

import (
	"net/http"

	"github.com/a-h/templ"
	views "github.com/mvaldes14/api-go-otel/views/layouts"
)

func IndexApp() http.Handler {
	return templ.Handler(views.IndexView())
}
