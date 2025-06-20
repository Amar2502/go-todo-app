package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Amar2502/go-todo-app/internal/config"
)

func main() {

	//load config
	cfg := config.MustLoad()


	//database setup


	//router setup
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server has started")
	})
	

	//server setup
	server := http.Server {
		Addr: cfg.HttpServer.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address", cfg.HttpServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Cannot Start Server: ", err.Error())
		}
	}()

	<- done

	slog.Info("Shutting down server")

	//gracefully shutting down server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}