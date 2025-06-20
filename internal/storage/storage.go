package storage

type Storage interface {
	InsertTodo(task string, start_time string) (lastId int64, err error)
}