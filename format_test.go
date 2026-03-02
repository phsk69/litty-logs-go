package litty

import (
	"log/slog"
	"strings"
	"testing"
	"time"
)

var testTime = time.Date(2026, 3, 2, 21, 45, 0, 420000000, time.UTC)

// TestFormatLogLine_EmitsCorrectEmoji — each level gets its proper emoji no cap 🔥
func TestFormatLogLine_EmitsCorrectEmoji(t *testing.T) {
	tests := []struct {
		level slog.Level
		emoji string
	}{
		{slog.Level(-8), "👀"},
		{slog.LevelDebug, "🔍"},
		{slog.LevelInfo, "🔥"},
		{slog.LevelWarn, "😤"},
		{slog.LevelError, "💀"},
	}

	opts := DefaultOptions()
	opts.UseColors = false

	for _, tt := range tests {
		line := FormatLogLine(tt.level, testTime, "test", "msg", opts)
		if !strings.Contains(line, tt.emoji) {
			t.Errorf("level %d: expected emoji %s in %q — where the vibes at 💀", tt.level, tt.emoji, line)
		}
	}
}

// TestFormatLogLine_EmitsCorrectLevelLabel — labels gotta match fr fr
func TestFormatLogLine_EmitsCorrectLevelLabel(t *testing.T) {
	tests := []struct {
		level slog.Level
		label string
	}{
		{slog.Level(-8), "trace"},
		{slog.LevelDebug, "debug"},
		{slog.LevelInfo, "info"},
		{slog.LevelWarn, "warn"},
		{slog.LevelError, "err"},
	}

	opts := DefaultOptions()
	opts.UseColors = false

	for _, tt := range tests {
		line := FormatLogLine(tt.level, testTime, "test", "msg", opts)
		if !strings.Contains(line, tt.label) {
			t.Errorf("level %d: expected label %q in %q — labels are off bestie 😤", tt.level, tt.label, line)
		}
	}
}

// TestFormatLogLine_IncludesIso8601Timestamp — gotta have that ISO 8601 for international rizz 🌍
func TestFormatLogLine_IncludesIso8601Timestamp(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "msg", opts)

	// should contain a timestamp in ISO 8601 format with milliseconds
	if !strings.Contains(line, "T") || !strings.Contains(line, "Z") {
		t.Errorf("expected ISO 8601 timestamp in %q — where the international rizz at 🌍", line)
	}
}

// TestFormatLogLine_IncludesAnsiColorCodes — colors gotta be there when enabled 🎨
func TestFormatLogLine_IncludesAnsiColorCodes(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = true

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "msg", opts)

	if !strings.Contains(line, "\x1b[") {
		t.Errorf("expected ANSI codes in %q — wheres the drip 🎨", line)
	}
}

// TestFormatLogLine_OmitsAnsiCodes_WhenColorsDisabled — no colors mode should be clean
func TestFormatLogLine_OmitsAnsiCodes_WhenColorsDisabled(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "msg", opts)

	if strings.Contains(line, "\x1b[") {
		t.Errorf("expected no ANSI codes in %q — colors should be off bestie", line)
	}
}

// TestFormatLogLine_DefaultOrder_LevelBeforeTimestamp — RFC 5424 style by default
func TestFormatLogLine_DefaultOrder_LevelBeforeTimestamp(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false
	opts.UseUtcTimestamp = true

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "msg", opts)

	levelIdx := strings.Index(line, "[🔥 info]")
	timeIdx := strings.Index(line, "[20")
	if levelIdx < 0 || timeIdx < 0 || levelIdx >= timeIdx {
		t.Errorf("expected level before timestamp in %q — RFC 5424 vibes only bestie", line)
	}
}

// TestFormatLogLine_TimestampFirst — observability style puts timestamp first
func TestFormatLogLine_TimestampFirst(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false
	opts.TimestampFirst = true
	opts.UseUtcTimestamp = true

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "msg", opts)

	levelIdx := strings.Index(line, "[🔥 info]")
	timeIdx := strings.Index(line, "[20")
	if levelIdx < 0 || timeIdx < 0 || timeIdx >= levelIdx {
		t.Errorf("expected timestamp before level in %q — observability besties need this", line)
	}
}

// TestFormatLogLine_SanitizesNewlines — log injection prevention is a must bestie 🔒
func TestFormatLogLine_SanitizesNewlines(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "line1\nline2", opts)

	if strings.Contains(line, "\n") {
		t.Errorf("newlines should be sanitized in %q — log injection is NOT it 🔒", line)
	}
	if !strings.Contains(line, "line1 line2") {
		t.Errorf("newlines should become spaces in %q", line)
	}
}

// TestFormatLogLine_SanitizesCrLf — windows style newlines get yeeted too
func TestFormatLogLine_SanitizesCrLf(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "line1\r\nline2", opts)

	if strings.Contains(line, "\r") || strings.Contains(line, "\n") {
		t.Errorf("CRLF should be sanitized in %q — log injection is NOT it 🔒", line)
	}
	if !strings.Contains(line, "line1 line2") {
		t.Errorf("CRLF should become a single space in %q", line)
	}
}

// TestFormatLogLine_CategoryShortening — namespace bloat gets yeeted when enabled
func TestFormatLogLine_CategoryShortening(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false
	opts.ShortenCategories = true

	line := FormatLogLine(slog.LevelInfo, testTime, "github.com/user/pkg.MyService", "msg", opts)

	if !strings.Contains(line, "[MyService]") {
		t.Errorf("expected shortened category in %q — namespace bloat survived 💀", line)
	}
}

// TestFormatLogLine_CategoryNoShortening — full category when shortening disabled
func TestFormatLogLine_CategoryNoShortening(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false
	opts.ShortenCategories = false

	line := FormatLogLine(slog.LevelInfo, testTime, "github.com/user/pkg.MyService", "msg", opts)

	if !strings.Contains(line, "[github.com/user/pkg.MyService]") {
		t.Errorf("expected full category in %q — shortening should be off", line)
	}
}

// TestFormatLogLine_IncludesMessage — the actual message better be there fr fr
func TestFormatLogLine_IncludesMessage(t *testing.T) {
	opts := DefaultOptions()
	opts.UseColors = false

	line := FormatLogLine(slog.LevelInfo, testTime, "test", "we vibing bestie", opts)

	if !strings.Contains(line, "we vibing bestie") {
		t.Errorf("expected message in %q — wheres the message at 💀", line)
	}
}
