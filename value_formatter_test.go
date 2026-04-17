package slogcolor_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/SladkyCitron/slogcolor"
	"github.com/fatih/color"
)

func TestHandlerUsesCustomValueFormatter(t *testing.T) {
	var buf bytes.Buffer

	h := slogcolor.NewHandler(&buf, &slogcolor.Options{
		Level:       slog.LevelDebug,
		NoColor:     true,
		NoTime:      true,
		SrcFileMode: slogcolor.Nop,
		MsgPrefix:   "",
		MsgColor:    color.New(),
		ValueFormatter: func(v slog.Value) string {
			return "custom(" + v.String() + ")"
		},
	})

	logger := slog.New(h)
	logger.Info("hello", "duration", 1500*time.Millisecond, "count", 7)

	out := buf.String()
	if !strings.Contains(out, "duration=custom(1.5s)") {
		t.Fatalf("expected custom formatter output for duration, got: %q", out)
	}
	if !strings.Contains(out, "count=custom(7)") {
		t.Fatalf("expected custom formatter output for int, got: %q", out)
	}
}

func TestHandlerUsesDefaultValueFormattingWhenFormatterNil(t *testing.T) {
	var buf bytes.Buffer

	h := slogcolor.NewHandler(&buf, &slogcolor.Options{
		Level:       slog.LevelDebug,
		NoColor:     true,
		NoTime:      true,
		SrcFileMode: slogcolor.Nop,
		MsgPrefix:   "",
		MsgColor:    color.New(),
	})

	logger := slog.New(h)
	logger.Info("hello", "duration", 1500*time.Millisecond)

	out := buf.String()
	if !strings.Contains(out, "duration=1.5s") {
		t.Fatalf("expected default value string formatting, got: %q", out)
	}
}
