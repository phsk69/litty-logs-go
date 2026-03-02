package litty

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
)

// TimestampFormat is the ISO 8601 layout with milliseconds for that international rizz 🌍
const TimestampFormat = "2006-01-02T15:04:05.000Z07:00"

// FormatLogLine formats a complete log line with all the litty goodness.
// this is the shared brain of litty-logs — every output path eats from here no cap 🧠
func FormatLogLine(level slog.Level, t time.Time, category string, msg string, opts *Options) string {
	info := GetLevelInfo(level)

	displayCategory := category
	if opts.ShortenCategories {
		displayCategory = ShortenCategory(category)
	}

	if opts.UseUtcTimestamp {
		t = t.UTC()
	}
	timestamp := t.Format(TimestampFormat)

	var sb strings.Builder

	// build the two bracket segments then order based on config
	var levelBracket, timestampBracket string

	if opts.UseColors {
		levelBracket = fmt.Sprintf("%s[%s %s]%s ", info.Color, info.Emoji, info.Label, Reset)
		timestampBracket = fmt.Sprintf("%s[%s] [%s]%s ", Dim, timestamp, displayCategory, Reset)
	} else {
		levelBracket = fmt.Sprintf("[%s %s] ", info.Emoji, info.Label)
		timestampBracket = fmt.Sprintf("[%s] [%s] ", timestamp, displayCategory)
	}

	if opts.TimestampFirst {
		sb.WriteString(timestampBracket)
		sb.WriteString(levelBracket)
	} else {
		sb.WriteString(levelBracket)
		sb.WriteString(timestampBracket)
	}

	// sanitize newlines to prevent log injection —
	// \n in a message would create fake log entries and thats NOT it bestie 🔒
	msg = strings.ReplaceAll(msg, "\r\n", " ")
	msg = strings.ReplaceAll(msg, "\n", " ")
	msg = strings.ReplaceAll(msg, "\r", " ")
	sb.WriteString(msg)

	return sb.String()
}

// formatAttr formats a single slog.Attr as key=value for inline display 💅
func formatAttr(a slog.Attr, useColors bool) string {
	if a.Equal(slog.Attr{}) {
		return ""
	}

	v := a.Value.Resolve()

	if useColors {
		return fmt.Sprintf(" %s%s%s=%s", Dim, a.Key, Reset, v.String())
	}
	return fmt.Sprintf(" %s=%s", a.Key, v.String())
}
