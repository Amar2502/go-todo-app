package storage

import "github.com/Amar2502/go-todo-app/internal/types"

type Storage interface {
	InsertTodo(task string, start_time string) (lastId int64, err error)
	ReadTodo() ([]types.Todo, error)
	DeleteTodo(id int64) (msg string, err error)
}