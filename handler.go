package litty

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
)

// Handler is the slog.Handler that makes ALL logs bussin with emojis,
// ANSI colors, and gen alpha energy fr fr 🔥
type Handler struct {
	opts     *Options
	group    string      // current group prefix (used as category)
	attrs    []slog.Attr // pre-resolved attributes
	mu       *sync.Mutex // protects Writer — shared across clones
}

// NewHandler creates a new litty handler with functional options.
// starts from bussin defaults and only overrides what you pass no cap 🔥
//
//	h := litty.NewHandler() // all defaults
//	h := litty.NewHandler(litty.WithColors(false)) // no colors, everything else stays bussin
func NewHandler(opts ...Option) *Handler {
	o := DefaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	return &Handler{
		opts: o,
		mu:   &sync.Mutex{},
	}
}

// NewHandlerWithOptions creates a new litty handler from a full Options struct.
// for power users who call DefaultOptions() and tweak from there 💅
func NewHandlerWithOptions(opts *Options) *Handler {
	if opts == nil {
		opts = DefaultOptions()
	}
	if opts.Writer == nil {
		opts.Writer = DefaultOptions().Writer
	}
	if opts.Level == nil {
		opts.Level = slog.LevelInfo
	}
	return &Handler{
		opts: opts,
		mu:   &sync.Mutex{},
	}
}

// NewLogger creates a new slog.Logger with the litty handler.
// convenience function so besties dont have to import slog separately 💅
//
//	logger := litty.NewLogger() // all defaults
//	logger := litty.NewLogger(litty.WithColors(false)) // no colors bestie
func NewLogger(opts ...Option) *slog.Logger {
	return slog.New(NewHandler(opts...))
}

// Enabled reports whether the handler handles records at the given level.
// if its below the minimum level it gets yeeted no cap 🗑️
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// Handle formats and writes the log record — this is where the magic happens bestie 🔥
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	// format the core log line
	line := FormatLogLine(r.Level, r.Time, h.category(), r.Message, h.opts)

	// append pre-resolved attrs and record attrs
	var attrStr strings.Builder
	for _, a := range h.attrs {
		attrStr.WriteString(formatAttr(a, h.opts.UseColors))
	}
	r.Attrs(func(a slog.Attr) bool {
		attrStr.WriteString(formatAttr(a, h.opts.UseColors))
		return true
	})

	if attrStr.Len() > 0 {
		line += attrStr.String()
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := fmt.Fprintln(h.opts.Writer, line)
	return err
}

// WithAttrs returns a new Handler with the given attributes pre-resolved.
// these attrs get attached to every log line from this handler bestie 💅
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(h.attrs), len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)

	for _, a := range attrs {
		if h.group != "" {
			a.Key = h.group + "." + a.Key
		}
		newAttrs = append(newAttrs, a)
	}

	return &Handler{
		opts:  h.opts,
		group: h.group,
		attrs: newAttrs,
		mu:    h.mu,
	}
}

// WithGroup returns a new Handler with the given group name.
// the group becomes the category in log output — namespace your vibes bestie 📦
func (h *Handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	newGroup := name
	if h.group != "" {
		newGroup = h.group + "." + name
	}
	return &Handler{
		opts:  h.opts,
		group: newGroup,
		attrs: h.attrs,
		mu:    h.mu,
	}
}

// category returns the group name used as category in log output.
// if no group is set, defaults to "app" bestie
func (h *Handler) category() string {
	if h.group != "" {
		return h.group
	}
	return "app"
}
