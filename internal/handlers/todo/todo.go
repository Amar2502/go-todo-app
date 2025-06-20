package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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

func ReadTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		todos, err := storage.ReadTodo()
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if todos == nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(errors.New("no todos found")))
			return
		}

		response.WriteJson(w, http.StatusOK, todos)

	}
}

func DeleteTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting student by id", slog.String("id", id))

		intId, err := 	strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		msg, err := storage.DeleteTodo(intId)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": msg})

	}
}