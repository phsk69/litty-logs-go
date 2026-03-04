package main

import (
	"strings"
	"testing"
)

func TestBuildRewriter(t *testing.T) {
	rw := NewBuildRewriter()

	tests := []struct {
		name    string
		input   string
		wantSub string
	}{
		{
			"package header",
			"# github.com/phsk69/litty-logs-go",
			"📦",
		},
		{
			"package header has pkg name",
			"# github.com/phsk69/litty-logs-go",
			"github.com/phsk69/litty-logs-go",
		},
		{
			"imported not used",
			`./main.go:10:5: imported and not used: "fmt"`,
			"bestie you imported",
		},
		{
			"imported not used has emoji",
			`./main.go:10:5: imported and not used: "fmt"`,
			"💀",
		},
		{
			"declared not used",
			"./main.go:10:5: x declared and not used",
			"yeet it bestie",
		},
		{
			"general compile error",
			"./main.go:10:5: undefined: thing",
			"💀 big L at",
		},
		{
			"general compile error has location",
			"./main.go:10:5: undefined: thing",
			"./main.go:10:5",
		},
		{
			"general compile error has message",
			"./main.go:10:5: undefined: thing",
			"undefined: thing",
		},
		{
			"unmatched line passes through",
			"some random output that aint a compile error",
			"some random output that aint a compile error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rw.Rewrite(tt.input)
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("Rewrite(%q) = %q, missing %q — build rewriter pattern is bricked 💀",
					tt.input, got, tt.wantSub)
			}
		})
	}
}

func TestBuildRewriter_SpecificBeforeGeneral(t *testing.T) {
	rw := NewBuildRewriter()

	// "imported and not used" should hit the specific rule, not the general one
	got := rw.Rewrite(`./main.go:10:5: imported and not used: "fmt"`)
	if !strings.Contains(got, "bestie you imported") {
		t.Errorf("specific import rule should fire before general error rule but got %q 💀", got)
	}
	if strings.Contains(got, "big L at") {
		t.Errorf("general error rule should NOT fire for import-not-used, got %q 💀", got)
	}
}
