package main

import (
	"strings"
	"testing"
)

func TestCleanRewriter(t *testing.T) {
	rw := NewCleanRewriter()

	tests := []struct {
		name    string
		input   string
		wantSub string
	}{
		{"rm file", "rm -f /path/to/binary", "🗑️"},
		{"rm file has path", "rm -f /path/to/binary", "/path/to/binary"},
		{"cd dir", "cd /path/to/dir", "📂"},
		{"cd dir has path", "cd /path/to/dir", "/path/to/dir"},
		{"unmatched passthrough", "some other clean output", "some other clean output"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rw.Rewrite(tt.input)
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("Rewrite(%q) = %q, missing %q — clean rewriter is bricked 💀",
					tt.input, got, tt.wantSub)
			}
		})
	}
}
