package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}
func (s *sqliteHandler) AddTodo(name string) *Todo {
	statement, err := s.db.Prepare("INSERT INTO todos(name, completed, createdAt) VALUES (?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(name, false)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	return &todo
}
func (s *sqliteHandler) RemoveTodo(id int) bool {
	statement, err := s.db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := result.RowsAffected()
	return cnt > 0
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	statement, err := s.db.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(complete, id)
	if err != nil {
		panic(err)
	}
	cnt, _ := result.RowsAffected()
	return cnt > 0
}
func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	query := "CREATE TABLE IF NOT EXISTS TODOS(id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT,completed BOOLEAN,createdAt DATETIME)"
	statement, _ := database.Prepare(query)
	statement.Exec()
	m := &sqliteHandler{db: database}
	return m
}
