package main

import (
	"log/slog"

	"github.com/phsk69/litty-logs-go"
)

func main() {
	// === default options — bussin out the box 🔥 ===
	logger := litty.NewLogger()

	logger.Debug("for when you lowkey wanna see everything 🔍")
	logger.Info("everything is bussin and vibing fr fr 🔥")
	logger.Warn("something kinda sus but we not panicking yet 😤")
	logger.Error("something took a fat L and we gotta deal with it 💀")

	// === with structured attributes 💅 ===
	logger.Info("request slid in", "method", "GET", "path", "/api/vibes", "status", 200)

	// === with a category via WithGroup 📦 ===
	serviceLogger := logger.WithGroup("MyService")
	serviceLogger.Info("service is cooking bestie")

	// === timestamp-first mode for the observability besties 📊 ===
	tsLogger := litty.NewLogger(litty.WithTimestampFirst(true))
	tsLogger.WithGroup("ObservabilityVibe").Info("timestamp comes first for the sort key besties")

	// === no colors mode for piping to files 📄 ===
	plainLogger := litty.NewLogger(litty.WithColors(false))
	plainLogger.Info("no colors — for when the terminal cant handle the drip")

	// === stack options like a combo meal 🍔 ===
	customLogger := litty.NewLogger(
		litty.WithColors(false),
		litty.WithTimestampFirst(true),
		litty.WithLevel(slog.LevelDebug),
	)
	customLogger.Debug("debug mode AND timestamp first AND no colors bestie")

	// === set as default slog logger 🌍 ===
	slog.SetDefault(logger)
	slog.Info("this uses the default litty logger now bestie")
}
