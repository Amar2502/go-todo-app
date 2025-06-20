package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Amar2502/go-todo-app/internal/storage"
	"github.com/Amar2502/go-todo-app/internal/types"
	"github.com/Amar2502/go-todo-app/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func AddTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var todo types.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)

		slog.Info("Decoded the request body", "todo", todo)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		slog.Info("Creating a new todo")

		if err := validator.New().Struct(todo); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		lastId, err := storage.InsertTodo(todo.Task, todo.StartTime)

		slog.Info("User created sucessfully", slog.String("LastId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"lastId": lastId})

	}
}
