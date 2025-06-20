package sqlite

import (
	"database/sql"

	"github.com/Amar2502/go-todo-app/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err!=nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		start_time DATETIME NOT NULL
	)`)
	if err!=nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
} 