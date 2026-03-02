package litty

import (
	"log/slog"
	"os"
	"testing"
)

// TestDefaultOptions_HasBussinDefaults — make sure the defaults are bussin out the box no cap 🔥
func TestDefaultOptions_HasBussinDefaults(t *testing.T) {
	opts := DefaultOptions()

	if opts.Level.Level() != slog.LevelInfo {
		t.Errorf("Level should be Info but got %v — thats not it bestie 💀", opts.Level)
	}
	if !opts.UseColors {
		t.Error("UseColors should be true — we need the drip by default 🎨")
	}
	if !opts.ShortenCategories {
		t.Error("ShortenCategories should be true — yeet that namespace bloat fr fr")
	}
	if opts.TimestampFirst {
		t.Error("TimestampFirst should be false — RFC 5424 style is the default vibe")
	}
	if !opts.UseUtcTimestamp {
		t.Error("UseUtcTimestamp should be true — international rizz by default 🌍")
	}
	if opts.Writer != os.Stderr {
		t.Error("Writer should be os.Stderr — thats where logs go bestie")
	}
}
