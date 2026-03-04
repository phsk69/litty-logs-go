package main

import (
	"strings"
	"testing"
)

func TestRunRewriter_DelegatesToBuild(t *testing.T) {
	rw := NewRunRewriter()

	// compile errors should get the build rewriter treatment
	got := rw.Rewrite("./main.go:10:5: undefined: thing")
	if !strings.Contains(got, "💀") {
		t.Errorf("run rewriter should delegate to build for compile errors but got %q 💀", got)
	}

	// program output should pass through unchanged
	line := "hello from my bussin program 🔥"
	got = rw.Rewrite(line)
	if got != line {
		t.Errorf("program output should pass through unchanged but got %q 💀", got)
	}
}
