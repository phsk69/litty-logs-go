package main

import (
	"log/slog"
	"os"

	litty "github.com/phsk69/litty-logs-go"
)

func main() {
	// one-liner JSON logger — bussin out the box 🔥
	logger := litty.NewJSONLogger(
		litty.WithWriter(os.Stdout),
		litty.WithLevel(slog.Level(-8)), // trace level so we see everything bestie
	)

	// all log levels with JSON output 🎯
	logger.Log(nil, slog.Level(-8), "lowkey peeking at everything")
	logger.Debug("investigating the vibes under the hood")
	logger.Info("everything is bussin and vibing", "method", "GET", "path", "/api/vibes")
	logger.Warn("something kinda sus but we not panicking", "retries", 3)
	logger.Error("something took a fat L", "err", "connection refused", "host", "db.internal")

	// WithGroup — categorized JSON logs 📦
	serviceLogger := logger.WithGroup("PaymentService")
	serviceLogger.Info("payment processed", "amount", 42.69, "currency", "USD")

	// WithAttrs — pre-resolved fields that show up on every log 💅
	requestLogger := logger.With("requestId", "abc-123", "userId", "user-69")
	requestLogger.Info("request slid in")
	requestLogger.Info("request processed", "duration", "420ms")

	// nested groups showing category shortening 📦
	nestedLogger := logger.WithGroup("Server").WithGroup("HTTP")
	nestedLogger.Info("handling request", "status", 200)
}
