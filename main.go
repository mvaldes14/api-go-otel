package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/mvaldes14/api-go-otel/views"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  index, err := views.IndexView().Render(ctx context.Background(), w)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", templ.Handler(indexHandler))

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", mux)

}
