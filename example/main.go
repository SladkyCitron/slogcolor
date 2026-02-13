package main

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/SladkyCitron/slogcolor"
)

func main() {
	opts := slogcolor.DefaultOptions
	opts.Level = slog.LevelDebug
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))

	slog.Info("Initializing")
	slog.Debug("Init done", "duration", 500*time.Millisecond)
	slog.Warn("Slow request!", "method", "GET", "path", "/api/users", "duration", 750*time.Millisecond)
	slog.Error("DB connection lost!", "err", errors.New("connection reset"), "db", "horalky")
}
