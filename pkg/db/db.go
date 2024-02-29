package db

import (
	"log"

	"github.com/mvaldes14/api-go-otel/pkg/db"
)

var schema = `
  CREATE TABLE TODOS (
    ID SERIAL PRIMARY KEY,
    TITLE TEXT,
    COMPLETED BOOLEAN
  )
`

type Todo struct {
	id        int
	title     string
	completed bool
}

func init() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
}

func AddTodo(t Todo) {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO TODOS (TITLE, COMPLETED) VALUES (?, ?)", t.title, t.completed)
	tx.Commit()
}

func GetTodos() ([]Todo, error) {
	todoList := []Todo{}
  err := db.Get($todoList, "SELECT * FROM TODOS")
  if err != nil {
    return nil, err
  }
}
