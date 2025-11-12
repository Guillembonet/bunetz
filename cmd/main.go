package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"

	"github.com/guillembonet/bunetz/external/telegram"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/server/handlers"
)

var (
	FlagToken  = flag.String("token", "", "Telegram bot token")
	FlagChatID = flag.Int64("chat-id", 0, "Telegram chat id")
)

func main() {
	flag.Parse()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	if *FlagToken == "" || *FlagChatID == 0 {
		slog.Error("token and chat-id are required")
		os.Exit(1)
	}

	telegramClient, err := telegram.NewClient(telegram.Config{
		Token:  *FlagToken,
		ChatID: *FlagChatID,
	})
	if err != nil {
		slog.Error("failed to create telegram client", slog.String("error", err.Error()))
		os.Exit(1)
	}

	static := handlers.NewStatic()
	blog := handlers.NewBlog()
	contact := handlers.NewContact(telegramClient)

	server, err := server.NewServer(":8080", static, blog, contact)
	if err != nil {
		slog.Error("failed to create server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	exitCode := atomic.Int32{}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", slog.String("error", err.Error()))
			exitCode.Store(1)
			cancel()
		}
	}()

	<-ctx.Done()
	slog.Info("Shutting down gracefully...")

	if err := server.Stop(); err != nil {
		slog.Error("Server failed to shutdown gracefully", slog.String("error", err.Error()))
		os.Exit(1)
	}

	<-stopped

	if code := exitCode.Load(); code != 0 {
		os.Exit(int(code))
	}
}
