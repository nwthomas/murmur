package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

// Logger wraps the structured logger
type Logger struct {
	*slog.Logger
}

// New creates a new logger instance
func New(level string, debug bool) *Logger {
	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn", "warning":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	var output io.Writer = os.Stdout
	if debug {
		// In debug mode, use a more verbose format
		opts := &slog.HandlerOptions{
			Level: logLevel,
			AddSource: true,
		}
		handler := slog.NewTextHandler(output, opts)
		return &Logger{slog.New(handler)}
	}

	// Production mode with JSON output
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(output, opts)
	return &Logger{slog.New(handler)}
}

// WithField adds a field to the logger context
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Logger.With(key, value)}
}

// WithFields adds multiple fields to the logger context
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	args := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	return &Logger{l.Logger.With(args...)}
}
