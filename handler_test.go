package slogcolor_test

import (
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/SladkyCitron/slogcolor"
	"github.com/fatih/color"
)

func Example() {
	opts := slogcolor.DefaultOptions
	opts.Level = slog.LevelDebug

	opts.LevelTags = map[slog.Level]string{
		slog.LevelInfo: color.New(color.BgCyan, color.FgBlack).Sprint("Info"),
	}

	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))
	slog.Info("Initializing")
	slog.Debug("Init done", "duration", 500*time.Millisecond)
	slog.Warn("Slow request!", "method", "GET", "path", "/api/users", "duration", 750*time.Millisecond)
	slog.Error("DB connection lost!", "err", errors.New("connection reset"), "db", "horalky")
	// Output:

	opts.NoTime = true
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))
	slog.Info("Message without time")
}

func BenchmarkLog(b *testing.B) {
	b.StopTimer()
	l := slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions))

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		l.Info("benchmarking", "i", i)
		b.StopTimer()
	}
}
