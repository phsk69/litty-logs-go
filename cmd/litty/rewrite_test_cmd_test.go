package main

import (
	"strings"
	"testing"
)

func TestTestRewriter_AllPatterns(t *testing.T) {
	rw := NewTestRewriter()

	tests := []struct {
		name    string
		input   string
		wantSub string
	}{
		{"run", "=== RUN   TestFoo", "🏃"},
		{"run has name", "=== RUN   TestFoo", "TestFoo"},
		{"run subtest", "=== RUN   TestFoo/sub_test", "TestFoo/sub_test"},
		{"pass", "--- PASS: TestFoo (0.00s)", "✅"},
		{"pass has name", "--- PASS: TestFoo (0.00s)", "TestFoo"},
		{"pass has time", "--- PASS: TestFoo (0.00s)", "0.00s"},
		{"pass has slayed", "--- PASS: TestFoo (0.00s)", "slayed"},
		{"fail", "--- FAIL: TestBar (0.01s)", "💀"},
		{"fail has name", "--- FAIL: TestBar (0.01s)", "TestBar"},
		{"fail has fat L", "--- FAIL: TestBar (0.01s)", "fat L"},
		{"skip", "--- SKIP: TestBaz (0.00s)", "⏭️"},
		{"skip has name", "--- SKIP: TestBaz (0.00s)", "TestBaz"},
		{"skip has not today", "--- SKIP: TestBaz (0.00s)", "not today bestie"},
		{"pause", "=== PAUSE TestConcurrent", "⏸️"},
		{"pause has name", "=== PAUSE TestConcurrent", "TestConcurrent"},
		{"cont", "=== CONT  TestConcurrent", "▶️"},
		{"cont has name", "=== CONT  TestConcurrent", "TestConcurrent"},
		{"standalone pass", "PASS", "🏆"},
		{"standalone pass has W", "PASS", "W in the chat"},
		{"standalone fail", "FAIL", "💀"},
		{"standalone fail has L", "FAIL", "massive L"},
		{"ok pkg with tab", "ok  \tgithub.com/foo/bar\t0.123s", "✅"},
		{"ok pkg has name", "ok  \tgithub.com/foo/bar\t0.123s", "github.com/foo/bar"},
		{"fail pkg with tab", "FAIL\tgithub.com/foo/bar\t0.456s", "💀"},
		{"fail pkg has name", "FAIL\tgithub.com/foo/bar\t0.456s", "github.com/foo/bar"},
		{"no test files", "?\tgithub.com/foo/bar\t[no test files]", "🤷"},
		{"no test files has name", "?\tgithub.com/foo/bar\t[no test files]", "github.com/foo/bar"},
		{"coverage", "coverage: 85.0% of statements", "📊"},
		{"coverage has percent", "coverage: 85.0% of statements", "85.0%"},
		{"testing warning", "testing: warning: no tests to run", "😤"},
		{"testing warning has msg", "testing: warning: no tests to run", "no tests to run"},
		{"indented output", "    bar_test.go:15: expected 5, got 3", "↳"},
		{"indented output has msg", "    bar_test.go:15: expected 5, got 3", "bar_test.go:15: expected 5, got 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rw.Rewrite(tt.input)
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("Rewrite(%q) = %q, missing %q — test rewriter pattern is bricked 💀",
					tt.input, got, tt.wantSub)
			}
		})
	}
}

func TestTestRewriter_FallsBackToBuild(t *testing.T) {
	rw := NewTestRewriter()

	// compile errors should hit the build rewriter fallback
	got := rw.Rewrite("# github.com/foo/bar")
	if !strings.Contains(got, "📦") {
		t.Errorf("test rewriter should fall back to build for package headers but got %q 💀", got)
	}

	got = rw.Rewrite("./main.go:10:5: undefined: thing")
	if !strings.Contains(got, "💀") || !strings.Contains(got, "big L") {
		t.Errorf("test rewriter should fall back to build for compile errors but got %q 💀", got)
	}
}

func TestTestRewriter_UnmatchedPassthrough(t *testing.T) {
	rw := NewTestRewriter()
	line := "some custom test log output that dont match nothing"
	got := rw.Rewrite(line)
	if got != line {
		t.Errorf("unmatched line should pass through but got %q 💀", got)
	}
}
