package litty

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"
)

func newTestHandler(buf *bytes.Buffer) *Handler {
	return NewHandler(WithWriter(buf), WithColors(false), WithUTC(true))
}

func makeRecord(level slog.Level, msg string) slog.Record {
	return slog.NewRecord(time.Date(2026, 3, 2, 21, 45, 0, 420000000, time.UTC), level, msg, 0)
}

// TestHandler_Enabled_RespectsMinimumLevel — levels below minimum get yeeted 🗑️
func TestHandler_Enabled_RespectsMinimumLevel(t *testing.T) {
	h := NewHandler(WithLevel(slog.LevelWarn))

	if h.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Debug should not be enabled when minimum is Warn — yeet it 🗑️")
	}
	if h.Enabled(context.Background(), slog.LevelInfo) {
		t.Error("Info should not be enabled when minimum is Warn — yeet it 🗑️")
	}
	if !h.Enabled(context.Background(), slog.LevelWarn) {
		t.Error("Warn should be enabled when minimum is Warn — let it through bestie")
	}
	if !h.Enabled(context.Background(), slog.LevelError) {
		t.Error("Error should be enabled when minimum is Warn — errors always get through no cap")
	}
}

// TestHandler_Handle_WritesToWriter — output goes to the configured writer bestie
func TestHandler_Handle_WritesToWriter(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	r := makeRecord(slog.LevelInfo, "we vibing")
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v — thats not bussin 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "we vibing") {
		t.Errorf("expected message in output %q — wheres the vibes 💀", output)
	}
	if !strings.Contains(output, "🔥") {
		t.Errorf("expected fire emoji in output %q — wheres the energy 🔥", output)
	}
}

// TestHandler_Handle_WithAttrs — pre-resolved attrs show up in output 💅
func TestHandler_Handle_WithAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	h2 := h.WithAttrs([]slog.Attr{slog.String("service", "api")})

	r := makeRecord(slog.LevelInfo, "request slid in")
	if err := h2.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "service=api") {
		t.Errorf("expected pre-resolved attr in output %q — attrs missing bestie 💀", output)
	}
}

// TestHandler_Handle_WithInlineAttrs — attrs from the log call show up too
func TestHandler_Handle_WithInlineAttrs(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	r := makeRecord(slog.LevelInfo, "request slid in")
	r.AddAttrs(slog.String("method", "GET"), slog.Int("status", 200))

	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "method=GET") {
		t.Errorf("expected method attr in output %q 💀", output)
	}
	if !strings.Contains(output, "status=200") {
		t.Errorf("expected status attr in output %q 💀", output)
	}
}

// TestHandler_WithGroup_SetsCategory — group name becomes the category in output 📦
func TestHandler_WithGroup_SetsCategory(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	h2 := h.WithGroup("MyService")

	r := makeRecord(slog.LevelInfo, "cooking")
	if err := h2.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "[MyService]") {
		t.Errorf("expected group as category in output %q — wheres the category 📦", output)
	}
}

// TestHandler_WithGroup_NestedGroups — nested groups join with dots
func TestHandler_WithGroup_NestedGroups(t *testing.T) {
	var buf bytes.Buffer
	h := NewHandler(WithWriter(&buf), WithColors(false), WithUTC(true), WithShortenCategories(false))

	h2 := h.WithGroup("Server").(*Handler)
	h3 := h2.WithGroup("HTTP")

	r := makeRecord(slog.LevelInfo, "request slid in")
	if err := h3.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "[Server.HTTP]") {
		t.Errorf("expected nested group in output %q — groups should join with dots bestie", output)
	}
}

// TestHandler_WithGroup_EmptyName — empty group name returns same handler
func TestHandler_WithGroup_EmptyName(t *testing.T) {
	h := NewHandler()
	h2 := h.WithGroup("")

	if h != h2 {
		t.Error("empty group name should return the same handler — dont clone for nothing bestie")
	}
}

// TestHandler_DefaultCategory — no group means category is "app"
func TestHandler_DefaultCategory(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	r := makeRecord(slog.LevelInfo, "vibing")
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "[app]") {
		t.Errorf("expected default category [app] in output %q", output)
	}
}

// TestHandler_ConcurrentWrites — mutex prevents interleaved output 🔒
func TestHandler_ConcurrentWrites(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

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
}

// TestHandler_NoOptions — no options should use defaults and not panic
func TestHandler_NoOptions(t *testing.T) {
	var buf bytes.Buffer
	h := NewHandler(WithWriter(&buf))

	r := makeRecord(slog.LevelInfo, "no opts vibing")
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle with no opts failed: %v — should use defaults bestie 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "no opts vibing") {
		t.Error("expected output with no options — defaults should kick in")
	}
	// UTC should be on by default — timestamp should end with Z not an offset
	if !strings.Contains(output, "Z]") {
		t.Errorf("expected UTC timestamp (Z suffix) in %q — UTC should be default bestie 🌍", output)
	}
}

// TestNewLogger_ReturnsWorkingLogger — convenience constructor works no cap
func TestNewLogger_ReturnsWorkingLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(WithWriter(&buf), WithColors(false))

	logger.Info("litty logger works")

	output := buf.String()
	if !strings.Contains(output, "litty logger works") {
		t.Errorf("NewLogger output %q missing message — convenience constructor is bricked 💀", output)
	}
	// UTC should be on by default even tho we only set colors and writer
	if !strings.Contains(output, "Z]") {
		t.Errorf("expected UTC timestamp in %q — WithColors shouldnt kill UTC bestie 🌍", output)
	}
}

// TestNewHandlerWithOptions_WorksForPowerUsers — struct API still works for the real ones 💅
func TestNewHandlerWithOptions_WorksForPowerUsers(t *testing.T) {
	var buf bytes.Buffer
	opts := DefaultOptions()
	opts.Writer = &buf
	opts.UseColors = false
	h := NewHandlerWithOptions(opts)

	r := makeRecord(slog.LevelInfo, "power user vibes")
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	if !strings.Contains(buf.String(), "power user vibes") {
		t.Errorf("NewHandlerWithOptions is bricked: %q 💀", buf.String())
	}
}

// TestHandler_WithAttrs_GroupPrefix — attrs get group prefix when group is set 📦
func TestHandler_WithAttrs_GroupPrefix(t *testing.T) {
	var buf bytes.Buffer
	h := newTestHandler(&buf)

	h2 := h.WithGroup("server")
	h3 := h2.WithAttrs([]slog.Attr{slog.String("port", "8080")})

	r := makeRecord(slog.LevelInfo, "listening")
	if err := h3.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle failed: %v 💀", err)
	}

	output := buf.String()
	if !strings.Contains(output, "server.port=8080") {
		t.Errorf("expected group-prefixed attr in output %q — group prefix missing bestie", output)
	}
}
