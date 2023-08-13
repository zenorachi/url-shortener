package logger

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func init() {
	slog.SetDefault(logger)
}

func InfoService(name string, args ...any) {
	slog.Info("service:"+name, args...)
}

func ErrorService(name string, args ...any) {
	slog.Error("service:"+name, args...)
}

func InfoHandler(name string, args ...any) {
	slog.Info("handler:"+name, args...)
}

func ErrorHandler(name string, args ...any) {
	slog.Error("handler:"+name, args...)
}
