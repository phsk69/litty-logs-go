package litty

import (
	"log/slog"
	"testing"
)

// TestGetLevelInfo_AllLevels — every level gets the right emoji, label, and color no cap 🔥
func TestGetLevelInfo_AllLevels(t *testing.T) {
	tests := []struct {
		name  string
		level slog.Level
		emoji string
		label string
		color string
	}{
		{"trace gets the eyes", slog.Level(-8), "👀", "trace", Cyan},
		{"debug gets the magnifier", slog.LevelDebug, "🔍", "debug", Blue},
		{"info gets the fire", slog.LevelInfo, "🔥", "info", Green},
		{"warn gets the angry face", slog.LevelWarn, "😤", "warn", Yellow},
		{"error gets the skull", slog.LevelError, "💀", "err", Red},
		{"error+4 still gets the skull", slog.LevelError + 4, "💀", "err", Red},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := GetLevelInfo(tt.level)
			if info.Emoji != tt.emoji {
				t.Errorf("emoji: got %q, want %q — wrong vibes bestie 💀", info.Emoji, tt.emoji)
			}
			if info.Label != tt.label {
				t.Errorf("label: got %q, want %q — thats not it 😤", info.Label, tt.label)
			}
			if info.Color != tt.color {
				t.Errorf("color: got %q, want %q — drip check failed 🎨", info.Color, tt.color)
			}
		})
	}
}
