// Package litty transforms your boring Go logs into bussin gen alpha energy
// with emojis, ANSI colors, and abbreviated categories no cap 🔥
//
// litty implements the [log/slog.Handler] interface so it works with Go's standard
// structured logging. one line setup and youre eating bestie:
//
//	logger := litty.NewLogger()
//	logger.Info("we vibing fr fr", "key", "value")
//	// output: [🔥 info] [2026-03-02T21:45:00.420Z] [app] we vibing fr fr key=value
//
// override only what you want — defaults stay bussin:
//
//	logger := litty.NewLogger(litty.WithColors(false), litty.WithTimestampFirst(true))
package litty
