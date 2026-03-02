package litty

import (
	"io"
	"log/slog"
	"os"
)

// Options controls the vibe of the litty handler bestie.
// all defaults are bussin out the box no cap 🔥
type Options struct {
	// Level is the minimum log level that gets through.
	// anything below this gets yeeted into the void 🗑️
	// default: slog.LevelInfo
	Level slog.Leveler

	// UseColors enables ANSI color codes in output.
	// disable if your terminal cant handle the drip 🎨
	// default: true
	UseColors bool

	// ShortenCategories yeets the namespace bloat from logger names.
	// "github.com/user/pkg.Service" becomes just "Service" fr fr
	// default: true
	ShortenCategories bool

	// TimestampFirst puts timestamp before level label in output.
	// false = [emoji level] [timestamp] [category] message (default, RFC 5424 vibes)
	// true = [timestamp] [emoji level] [category] message (observability style)
	// default: false
	TimestampFirst bool

	// UseUtcTimestamp uses UTC instead of local time.
	// true for that international rizz 🌍
	// default: true
	UseUtcTimestamp bool

	// Writer is where the output goes bestie.
	// default: os.Stderr
	Writer io.Writer
}

// DefaultOptions returns options that are bussin out the box no cap 🔥
func DefaultOptions() *Options {
	return &Options{
		Level:             slog.LevelInfo,
		UseColors:         true,
		ShortenCategories: true,
		TimestampFirst:    false,
		UseUtcTimestamp:   true,
		Writer:            os.Stderr,
	}
}

// Option is a functional option for configuring the litty handler.
// use With* funcs so you only override what you want and the bussin defaults stay intact 💅
type Option func(*Options)

// WithLevel sets the minimum log level — anything below gets yeeted 🗑️
func WithLevel(level slog.Leveler) Option {
	return func(o *Options) { o.Level = level }
}

// WithColors enables or disables ANSI color codes 🎨
func WithColors(v bool) Option {
	return func(o *Options) { o.UseColors = v }
}

// WithShortenCategories enables or disables namespace bloat yeeting
func WithShortenCategories(v bool) Option {
	return func(o *Options) { o.ShortenCategories = v }
}

// WithTimestampFirst puts timestamp before level label (observability style) 📊
func WithTimestampFirst(v bool) Option {
	return func(o *Options) { o.TimestampFirst = v }
}

// WithUTC enables or disables UTC timestamps for that international rizz 🌍
func WithUTC(v bool) Option {
	return func(o *Options) { o.UseUtcTimestamp = v }
}

// WithWriter sets the output destination bestie
func WithWriter(w io.Writer) Option {
	return func(o *Options) { o.Writer = w }
}
