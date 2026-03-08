package main

import (
	"strings"
	"testing"
)

func TestVetRewriter(t *testing.T) {
	rw := NewVetRewriter()

	tests := []struct {
		name    string
		input   string
		wantSub string
	}{
		{"package header", "# github.com/foo/bar", "🔍"},
		{"package header has name", "# github.com/foo/bar", "github.com/foo/bar"},
		{"vet finding", "./main.go:15:2: printf: fmt.Sprintf format %d has arg str of wrong type string", "😤"},
		{"vet finding has location", "./main.go:15:2: printf: fmt.Sprintf format %d has arg str of wrong type string", "./main.go:15:2"},
		{"vet finding has message", "./main.go:15:2: printf: fmt.Sprintf format %d has arg str of wrong type string", "vet says:"},
		{"unmatched passthrough", "some random line", "some random line"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rw.Rewrite(tt.input)
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("Rewrite(%q) = %q, missing %q — vet rewriter is bricked 💀",
					tt.input, got, tt.wantSub)
			}
		})
	}
}
