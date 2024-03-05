package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
  CREATE TABLE IF NOT EXISTS TODOS (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    TITLE TEXT,
    COMPLETED BOOLEAN
  )
`

type Todo struct {
	ID        int    `db:"ID"`
	Title     string `db:"TITLE"`
	Completed bool   `db:"COMPLETED"`
}

var db *sqlx.DB

func InitDb() {
	var err error
	db, err = sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	log.Println("Database initialized")
}

// Injects a todo into the database
func AddTodo(t Todo) {
	tx := db.MustBegin()
	log.Println("Adding todo:", t)
	tx.MustExec("INSERT INTO TODOS (TITLE, COMPLETED) VALUES (?, ?)", t.Title, t.Completed)
	tx.Commit()
}

// Returns a List of Todos from the database
func GetTodos() ([]Todo, error) {
	var todos []Todo
	err := db.Select(&todos, "SELECT * FROM TODOS")
	if err != nil {
		return nil, err
	}
	log.Println("DB Returned:", todos)
	return todos, err
}
