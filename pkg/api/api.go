// Package api provides the HTTP handlers for the API endpoints
package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mvaldes14/api-go-otel/pkg/database"
	"github.com/mvaldes14/api-go-otel/views/components"
)

// IndexHandler Returns a 200 OK for health checks
func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy")
}

// AddTodoHandler Creates a new Todo object and adds it to the database
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	todoTitle := r.FormValue("name")
	t := database.Todo{Title: todoTitle, Completed: false}
	database.AddTodo(t)
}

// GetTodoHandler Returns a list of todos from the database
func GetTodoHandler(w http.ResponseWriter, _ *http.Request) {
	todo, err := database.GetTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error loading todos from database:%v", err)
		return
	}
	components.TodoList(todo).Render(context.Background(), w)
}

// DeleteTodoHandler removes a todo from the database
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return
	}
	database.DeleteTodos(id)
	fmt.Fprintf(w, "Todo Deleted")
}
