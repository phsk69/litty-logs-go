package litty

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"
)

func newTestJSONHandler(buf *bytes.Buffer) *JSONHandler {
	return NewJSONHandler(WithWriter(buf), WithUTC(true))
}

func parseJSON(t *testing.T, raw string) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("JSON parse failed: %v — output is bricked: %q 💀", err, raw)
	}
	return m
}

// TestJSONHandler_ProducesValidJson — output must parse as JSON no cap
func TestJSONHandler_ProducesValidJson(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "we vibing")
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	parseJSON(t, strings.TrimSpace(buf.String()))
}

// TestJSONHandler_HasAllRequiredFields — timestamp, level, emoji, category, message
func TestJSONHandler_HasAllRequiredFields(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "we vibing")
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))

	required := []string{"timestamp", "level", "emoji", "category", "message"}
	for _, field := range required {
		if _, ok := m[field]; !ok {
			t.Errorf("missing required field %q in JSON — thats not bussin 💀", field)
		}
	}
}

// TestJSONHandler_CorrectLevelAndEmoji — each level gets the right label and emoji
func TestJSONHandler_CorrectLevelAndEmoji(t *testing.T) {
	tests := []struct {
		name  string
		level slog.Level
		label string
		emoji string
	}{
		{"trace gets eyes", slog.Level(-8), "trace", "👀"},
		{"debug gets magnifier", slog.LevelDebug, "debug", "🔍"},
		{"info gets fire", slog.LevelInfo, "info", "🔥"},
		{"warn gets angry", slog.LevelWarn, "warn", "😤"},
		{"error gets skull", slog.LevelError, "err", "💀"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			h := NewJSONHandler(WithWriter(&buf), WithUTC(true), WithLevel(slog.Level(-8)))

			r := makeRecord(tt.level, "test")
			_ = h.Handle(context.Background(), r)

			m := parseJSON(t, strings.TrimSpace(buf.String()))
			if m["level"] != tt.label {
				t.Errorf("level = %q, want %q 💀", m["level"], tt.label)
			}
			if m["emoji"] != tt.emoji {
				t.Errorf("emoji = %q, want %q 💀", m["emoji"], tt.emoji)
			}
		})
	}
}

// TestJSONHandler_ShortensCategoryByDefault — Server.HTTP becomes HTTP
func TestJSONHandler_ShortensCategoryByDefault(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)
	h2 := h.WithGroup("Server").(*JSONHandler).WithGroup("HTTP")

	r := makeRecord(slog.LevelInfo, "request")
	_ = h2.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["category"] != "HTTP" {
		t.Errorf("category = %q, want %q — shortening is bricked 💀", m["category"], "HTTP")
	}
}

// TestJSONHandler_FullCategoryWhenDisabled — no shortening means full group path
func TestJSONHandler_FullCategoryWhenDisabled(t *testing.T) {
	var buf bytes.Buffer
	h := NewJSONHandler(WithWriter(&buf), WithUTC(true), WithShortenCategories(false))
	h2 := h.WithGroup("Server").(*JSONHandler).WithGroup("HTTP")

	r := makeRecord(slog.LevelInfo, "request")
	_ = h2.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["category"] != "Server.HTTP" {
		t.Errorf("category = %q, want %q — full category is bricked 💀", m["category"], "Server.HTTP")
	}
}

// TestJSONHandler_DefaultCategory — no group means "app"
func TestJSONHandler_DefaultCategory(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "vibing")
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["category"] != "app" {
		t.Errorf("category = %q, want %q — default category is bricked 💀", m["category"], "app")
	}
}

// TestJSONHandler_PreResolvedAttrs — WithAttrs attrs show up in JSON
func TestJSONHandler_PreResolvedAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)
	h2 := h.WithAttrs([]slog.Attr{slog.String("service", "api")})

	r := makeRecord(slog.LevelInfo, "request")
	_ = h2.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["service"] != "api" {
		t.Errorf("service = %q, want %q — pre-resolved attrs missing 💀", m["service"], "api")
	}
}

// TestJSONHandler_InlineAttrs — record attrs show up in JSON
func TestJSONHandler_InlineAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "request")
	r.AddAttrs(slog.String("method", "GET"), slog.Int("status", 200))
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["method"] != "GET" {
		t.Errorf("method = %q, want %q 💀", m["method"], "GET")
	}
	// JSON numbers come back as float64 from json.Unmarshal
	if m["status"] != float64(200) {
		t.Errorf("status = %v, want %v — int attr is bricked 💀", m["status"], 200)
	}
}

// TestJSONHandler_GroupPrefixOnAttrs — WithGroup + WithAttrs gives prefixed keys
func TestJSONHandler_GroupPrefixOnAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)
	h2 := h.WithGroup("server")
	h3 := h2.WithAttrs([]slog.Attr{slog.String("port", "8080")})

	r := makeRecord(slog.LevelInfo, "listening")
	_ = h3.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["server.port"] != "8080" {
		t.Errorf("server.port = %q, want %q — group prefix is bricked 💀", m["server.port"], "8080")
	}
}

// TestJSONHandler_NoAnsiCodes — JSON must NEVER contain ANSI escape codes
func TestJSONHandler_NoAnsiCodes(t *testing.T) {
	var buf bytes.Buffer
	// explicitly enable colors — they should still not appear in JSON
	h := NewJSONHandler(WithWriter(&buf), WithColors(true), WithUTC(true))

	r := makeRecord(slog.LevelInfo, "no colors in JSON bestie")
	_ = h.Handle(context.Background(), r)

	output := buf.String()
	if strings.Contains(output, "\x1b") {
		t.Errorf("JSON output contains ANSI escape codes — thats not bussin: %q 💀", output)
	}
}

// TestJSONHandler_LiteralEmojis — raw output has literal emoji chars, not escaped surrogate pairs
func TestJSONHandler_LiteralEmojis(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "vibing")
	_ = h.Handle(context.Background(), r)

	output := buf.String()
	if !strings.Contains(output, "🔥") {
		t.Errorf("JSON output should have literal 🔥 emoji but got: %q 💀", output)
	}
	// make sure its not escaped as surrogate pairs
	if strings.Contains(output, "\\uD83D") || strings.Contains(output, "\\uDD25") {
		t.Errorf("JSON output has escaped surrogate pairs instead of literal emojis: %q 💀", output)
	}
}

// TestJSONHandler_SanitizesNewlines — newlines in message get replaced with spaces
func TestJSONHandler_SanitizesNewlines(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "line1\nline2\r\nline3")
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	msg := m["message"].(string)
	if strings.Contains(msg, "\n") || strings.Contains(msg, "\r") {
		t.Errorf("message should have sanitized newlines but got: %q 💀", msg)
	}
}

// TestJSONHandler_ConcurrentWrites — mutex prevents interleaved JSON 🔒
func TestJSONHandler_ConcurrentWrites(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := makeRecord(slog.LevelInfo, "concurrent vibes")
			_ = h.Handle(context.Background(), r)
		}()
	}
	wg.Wait()

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 100 {
		t.Errorf("expected 100 lines, got %d — mutex might be slacking 🔒", len(lines))
	}

	// every line must be valid JSON
	for i, line := range lines {
		var m map[string]any
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			t.Errorf("line %d is not valid JSON: %v — interleaving detected 💀", i, err)
		}
	}
}

// TestJSONHandler_UTCTimestamp — timestamp ends with Z when UTC enabled
func TestJSONHandler_UTCTimestamp(t *testing.T) {
	var buf bytes.Buffer
	h := NewJSONHandler(WithWriter(&buf), WithUTC(true))

	r := makeRecord(slog.LevelInfo, "utc vibes")
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	ts := m["timestamp"].(string)
	if !strings.HasSuffix(ts, "Z") {
		t.Errorf("UTC timestamp should end with Z but got: %q 💀", ts)
	}
}

// TestJSONHandler_LocalTimestamp — timestamp has offset when UTC disabled
func TestJSONHandler_LocalTimestamp(t *testing.T) {
	var buf bytes.Buffer
	h := NewJSONHandler(WithWriter(&buf), WithUTC(false))

	r := slog.NewRecord(time.Date(2026, 3, 2, 21, 45, 0, 420000000, time.FixedZone("EST", -5*3600)), slog.LevelInfo, "local vibes", 0)
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	ts := m["timestamp"].(string)
	if strings.HasSuffix(ts, "Z") {
		t.Errorf("local timestamp should NOT end with Z but got: %q 💀", ts)
	}
}

// TestNewJSONLogger — convenience constructor works no cap
func TestNewJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJSONLogger(WithWriter(&buf), WithUTC(true))

	logger.Info("litty JSON works")

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["message"] != "litty JSON works" {
		t.Errorf("message = %q, want %q — convenience constructor is bricked 💀", m["message"], "litty JSON works")
	}
}

// TestJSONHandler_BoolAttr — boolean attrs come through as JSON booleans
func TestJSONHandler_BoolAttr(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "test")
	r.AddAttrs(slog.Bool("active", true))
	_ = h.Handle(context.Background(), r)

	m := parseJSON(t, strings.TrimSpace(buf.String()))
	if m["active"] != true {
		t.Errorf("active = %v, want true — bool attr is bricked 💀", m["active"])
	}
}

// TestJSONHandler_EmptyGroupIgnored — empty group name returns same handler
func TestJSONHandler_EmptyGroupIgnored(t *testing.T) {
	h := NewJSONHandler()
	h2 := h.WithGroup("")
	if h != h2 {
		t.Error("empty group name should return same handler — dont clone for nothing bestie")
	}
}

// TestJSONHandler_FieldOrdering — timestamp comes first, then level, emoji, category, message
func TestJSONHandler_FieldOrdering(t *testing.T) {
	var buf bytes.Buffer
	h := newTestJSONHandler(&buf)

	r := makeRecord(slog.LevelInfo, "ordering test")
	_ = h.Handle(context.Background(), r)

	output := strings.TrimSpace(buf.String())
	// verify ordering by checking that timestamp appears before level which appears before message
	tsIdx := strings.Index(output, `"timestamp"`)
	lvlIdx := strings.Index(output, `"level"`)
	emojiIdx := strings.Index(output, `"emoji"`)
	catIdx := strings.Index(output, `"category"`)
	msgIdx := strings.Index(output, `"message"`)

	if tsIdx >= lvlIdx || lvlIdx >= emojiIdx || emojiIdx >= catIdx || catIdx >= msgIdx {
		t.Errorf("field ordering is bricked — expected timestamp < level < emoji < category < message in: %q 💀", output)
	}
}
