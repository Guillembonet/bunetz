package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/server/handlers"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	static := handlers.NewStatic()
	blog := handlers.NewBlog()

	server, err := server.NewServer(":8080", static, blog)
	if err != nil {
		slog.Error("failed to create server", slog.Any("err", err))
		os.Exit(1)
	}

	go func() {
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	slog.Info("Server shutting down...")

	if err := server.Stop(); err != nil {
		slog.Error("Server failed to shutdown gracefully", slog.Any("err", err))
		os.Exit(1)
	}
}
