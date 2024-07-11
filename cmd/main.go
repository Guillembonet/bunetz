package main

import (
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
		slog.Error("failed to create telegram client", slog.Any("err", err))
		os.Exit(1)
	}

	static := handlers.NewStatic()
	blog := handlers.NewBlog()
	contact := handlers.NewContact(telegramClient)

	server, err := server.NewServer(":8080", static, blog, contact)
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
