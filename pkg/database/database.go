// Package database contains the database logic
package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// Todo struct must contain the attributes and they should match the schema
type Todo struct {
	ID        int    `db:"ID"`
	Title     string `db:"TITLE"`
	Completed bool   `db:"COMPLETED"`
}

// TodoList is a list of Todos
type TodoList []Todo

// Database instance
var db *sqlx.DB

// InitDb initializes the DB and creates the schema
func InitDb() {
	var err error
	db, err = sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(`CREATE TABLE IF NOT EXISTS TODOS (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    TITLE TEXT,
    COMPLETED BOOLEAN
  )`)
	log.Println("Database initialized")
}

// AddTodo injects a todo into the database
func AddTodo(t Todo) {
	tx := db.MustBegin()
	defer tx.Commit()
	tx.MustExec("INSERT INTO TODOS (TITLE, COMPLETED) VALUES (?, ?)", t.Title, t.Completed)
}

// GetTodos returns a List of Todos from the database
func GetTodos() (TodoList, error) {
	var todos TodoList
	if db == nil {
		fmt.Println("DB not initialized")
		return nil, nil
	}
	err := db.Select(&todos, "SELECT * FROM TODOS")
	if err != nil {
		return todos, err
	}
	return todos, nil
}
