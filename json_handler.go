package litty

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"sync"
)

// JSONHandler is the slog.Handler that outputs structured JSON with litty energy 🔥
// one compact JSON object per line — machines can parse it AND it still looks bussin no cap
type JSONHandler struct {
	opts  *Options
	group string      // current group prefix (used as category)
	attrs []slog.Attr // pre-resolved attributes
	mu    *sync.Mutex // protects Writer — shared across clones
}

// NewJSONHandler creates a new litty JSON handler with functional options.
// same API as NewHandler but outputs JSON instead of text brackets bestie 💅
//
//	h := litty.NewJSONHandler() // all defaults
//	h := litty.NewJSONHandler(litty.WithLevel(slog.LevelDebug)) // debug mode JSON
func NewJSONHandler(opts ...Option) *JSONHandler {
	o := DefaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	return &JSONHandler{
		opts: o,
		mu:   &sync.Mutex{},
	}
}

// NewJSONHandlerWithOptions creates a new litty JSON handler from a full Options struct.
// for power users who like struct control bestie 💅
func NewJSONHandlerWithOptions(opts *Options) *JSONHandler {
	if opts == nil {
		opts = DefaultOptions()
	}
	if opts.Writer == nil {
		opts.Writer = DefaultOptions().Writer
	}
	if opts.Level == nil {
		opts.Level = slog.LevelInfo
	}
	return &JSONHandler{
		opts: opts,
		mu:   &sync.Mutex{},
	}
}

// NewJSONLogger creates a new slog.Logger with the litty JSON handler.
// convenience so besties dont have to import slog separately 💅
//
//	logger := litty.NewJSONLogger() // JSON bussin out the box
//	logger := litty.NewJSONLogger(litty.WithLevel(slog.LevelDebug))
func NewJSONLogger(opts ...Option) *slog.Logger {
	return slog.New(NewJSONHandler(opts...))
}

// Enabled reports whether the handler handles records at the given level.
// if its below the minimum level it gets yeeted no cap 🗑️
func (h *JSONHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// Handle formats and writes the log record as a single JSON object — structured vibes bestie 🔥
func (h *JSONHandler) Handle(_ context.Context, r slog.Record) error {
	info := GetLevelInfo(r.Level)

	cat := h.category()
	if h.opts.ShortenCategories {
		cat = ShortenCategory(cat)
	}

	t := r.Time
	if h.opts.UseUtcTimestamp {
		t = t.UTC()
	}

	// sanitize newlines in message — same as text handler 🔒
	msg := strings.ReplaceAll(r.Message, "\r\n", " ")
	msg = strings.ReplaceAll(msg, "\n", " ")
	msg = strings.ReplaceAll(msg, "\r", " ")

	// build ordered JSON manually for consistent field ordering 🔥
	// encoding/json.Marshal uses map which randomizes order — not bussin
	var buf bytes.Buffer
	buf.WriteByte('{')

	writeString := func(key, val string) {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		keyBytes, _ := json.Marshal(key)
		valBytes, _ := json.Marshal(val)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		buf.Write(valBytes)
	}

	writeString("timestamp", t.Format(TimestampFormat))
	writeString("level", info.Label)
	writeString("emoji", info.Emoji)
	writeString("category", cat)
	writeString("message", msg)

	// write attrs as flat key:value pairs at root level
	writeAttr := func(a slog.Attr) {
		if a.Equal(slog.Attr{}) {
			return
		}
		v := a.Value.Resolve()
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		keyBytes, _ := json.Marshal(a.Key)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		// use the appropriate JSON type for the value bestie 💅
		switch v.Kind() {
		case slog.KindInt64:
			valBytes, _ := json.Marshal(v.Int64())
			buf.Write(valBytes)
		case slog.KindUint64:
			valBytes, _ := json.Marshal(v.Uint64())
			buf.Write(valBytes)
		case slog.KindFloat64:
			valBytes, _ := json.Marshal(v.Float64())
			buf.Write(valBytes)
		case slog.KindBool:
			valBytes, _ := json.Marshal(v.Bool())
			buf.Write(valBytes)
		case slog.KindDuration:
			valBytes, _ := json.Marshal(v.Duration().String())
			buf.Write(valBytes)
		case slog.KindTime:
			valBytes, _ := json.Marshal(v.Time().Format(TimestampFormat))
			buf.Write(valBytes)
		default:
			valBytes, _ := json.Marshal(v.String())
			buf.Write(valBytes)
		}
	}

	// pre-resolved attrs first, then inline record attrs
	for _, a := range h.attrs {
		writeAttr(a)
	}
	r.Attrs(func(a slog.Attr) bool {
		writeAttr(a)
		return true
	})

	buf.WriteByte('}')
	buf.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.opts.Writer.Write(buf.Bytes())
	return err
}

// WithAttrs returns a new JSONHandler with the given attributes pre-resolved.
// same clone pattern as text handler bestie 💅
func (h *JSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(h.attrs), len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)

	for _, a := range attrs {
		if h.group != "" {
			a.Key = h.group + "." + a.Key
		}
		newAttrs = append(newAttrs, a)
	}

	return &JSONHandler{
		opts:  h.opts,
		group: h.group,
		attrs: newAttrs,
		mu:    h.mu,
	}
}

// WithGroup returns a new JSONHandler with the given group name.
// the group becomes the category in JSON output — namespace your vibes bestie 📦
func (h *JSONHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	newGroup := name
	if h.group != "" {
		newGroup = h.group + "." + name
	}
	return &JSONHandler{
		opts:  h.opts,
		group: newGroup,
		attrs: h.attrs,
		mu:    h.mu,
	}
}

// category returns the group name used as category in JSON output.
// if no group is set, defaults to "app" bestie
func (h *JSONHandler) category() string {
	if h.group != "" {
		return h.group
	}
	return "app"
}
