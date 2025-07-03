package middleware

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// TODO: do a better log, maybe use slog?
func InitLogger() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	slog.SetDefault(Logger)
}
