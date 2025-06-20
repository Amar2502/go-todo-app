package sqlite

import (
	"database/sql"

	"github.com/Amar2502/go-todo-app/internal/config"
	"github.com/Amar2502/go-todo-app/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		start_time DATETIME NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) InsertTodo(task string, start_time string) (int64, error) {

	stmt, err := s.Db.Prepare(`INSERT INTO todos(task, start_time) VALUES (?, ?)`)

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(task, start_time)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil

}

func (s *Sqlite) ReadTodo() ([]types.Todo, error) {

	stmt, err := s.Db.Prepare(`SELECT * FROM todos`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []types.Todo

	for rows.Next() {

		var todo types.Todo

		err := rows.Scan(&todo.ID, &todo.Task, &todo.StartTime)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)

	}

	return todos, nil
}


func (s *Sqlite) DeleteTodo(id int64) (msg string, err error) {

	stmt, err := s.Db.Prepare(`DELETE FROM todos WHERE id = ?`)

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return "", err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	if rows == 0 {
		return "no todo found", nil
	}

	return "todo deleted successfully", nil
}
