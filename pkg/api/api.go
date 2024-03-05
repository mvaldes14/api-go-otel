package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mvaldes14/api-go-otel/pkg/database"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy")
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	t := database.Todo{}
	json.Unmarshal(payload, &t)

	database.AddTodo(t)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo, err := database.GetTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error loading todos from database:%v", err)
		return
	}
	fmt.Fprintf(w, "Todos: %v", todo)
}
