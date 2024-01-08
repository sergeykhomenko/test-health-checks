package app

import (
	"log/slog"
	"os"
)

func InitLogger(debugMode bool) {
	minLogLevel := slog.LevelInfo
	if debugMode {
		minLogLevel = slog.LevelDebug
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: minLogLevel,
	})

	slog.SetDefault(slog.New(logHandler))
}
