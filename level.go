package litty

import "log/slog"

// LevelInfo holds the emoji, label, and ANSI color for a log level.
// this is the litty level info that makes every log line bussin 🔥
type LevelInfo struct {
	Emoji string
	Label string
	Color string
}

// ANSI escape codes that hit different 🎨
const (
	Reset     = "\x1b[0m"
	Dim       = "\x1b[2m"
	Cyan      = "\x1b[36m"
	Blue      = "\x1b[34m"
	Green     = "\x1b[32m"
	Yellow    = "\x1b[33m"
	Red       = "\x1b[31m"
	BrightRed = "\x1b[91m"
)

// GetLevelInfo returns the emoji, label, and ANSI color for a slog level.
// Go's slog has Debug (-4), Info (0), Warn (4), Error (8).
// we also handle custom levels for trace vibes 👀
func GetLevelInfo(level slog.Level) LevelInfo {
	switch {
	case level < slog.LevelDebug:
		return LevelInfo{"👀", "trace", Cyan}
	case level < slog.LevelInfo:
		return LevelInfo{"🔍", "debug", Blue}
	case level < slog.LevelWarn:
		return LevelInfo{"🔥", "info", Green}
	case level < slog.LevelError:
		return LevelInfo{"😤", "warn", Yellow}
	default:
		return LevelInfo{"💀", "err", Red}
	}
}
