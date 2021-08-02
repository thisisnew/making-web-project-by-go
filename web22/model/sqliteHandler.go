package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos WHERE sessionId = ?", sessionId)
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
func (s *sqliteHandler) AddTodo(name string, sessionId string) *Todo {
	statement, err := s.db.Prepare("INSERT INTO todos(sessionId, name, completed, createdAt) VALUES (?, ?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(sessionId, name, false)
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

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	query := "CREATE TABLE IF NOT EXISTS TODOS(id INTEGER PRIMARY KEY AUTOINCREMENT, sessionId STRING, name TEXT, completed BOOLEAN, createdAt DATETIME);"
	query += "CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON todos (sessionId ASC);"
	statement, _ := database.Prepare(query)
	statement.Exec()
	m := &sqliteHandler{db: database}
	return m
}
