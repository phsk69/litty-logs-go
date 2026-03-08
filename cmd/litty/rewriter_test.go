package main

import (
	"regexp"
	"testing"
)

func TestRuleRewriter_FirstMatchWins(t *testing.T) {
	rw := &RuleRewriter{
		Rules: []Rule{
			{
				Pattern: regexp.MustCompile(`^hello`),
				Replace: func(_ []string) string { return "first rule hit 🔥" },
			},
			{
				Pattern: regexp.MustCompile(`^hello`),
				Replace: func(_ []string) string { return "second rule should never hit 💀" },
			},
		},
	}

	got := rw.Rewrite("hello world")
	if got != "first rule hit 🔥" {
		t.Errorf("first match should win but got %q — composable pattern is bricked 💀", got)
	}
}

func TestRuleRewriter_FallbackWhenNoMatch(t *testing.T) {
	fallback := &RuleRewriter{
		Rules: []Rule{
			{
				Pattern: regexp.MustCompile(`.*`),
				Replace: func(_ []string) string { return "fallback caught it 💅" },
			},
		},
	}

	rw := &RuleRewriter{
		Rules: []Rule{
			{
				Pattern: regexp.MustCompile(`^wont match this`),
				Replace: func(_ []string) string { return "nope" },
			},
		},
		Fallback: fallback,
	}

	got := rw.Rewrite("some random line")
	if got != "fallback caught it 💅" {
		t.Errorf("should fall back when no rules match but got %q 💀", got)
	}
}

func TestRuleRewriter_PassthroughWhenNoMatchNoFallback(t *testing.T) {
	rw := &RuleRewriter{
		Rules: []Rule{
			{
				Pattern: regexp.MustCompile(`^wont match`),
				Replace: func(_ []string) string { return "nope" },
			},
		},
	}

	line := "this line stays untouched bestie"
	got := rw.Rewrite(line)
	if got != line {
		t.Errorf("unmatched line should pass through unchanged but got %q 💀", got)
	}
}

func TestRuleRewriter_CaptureGroups(t *testing.T) {
	rw := &RuleRewriter{
		Rules: []Rule{
			{
				Pattern: regexp.MustCompile(`^hello (\w+)$`),
				Replace: func(m []string) string { return "yo " + m[1] + " 🔥" },
			},
		},
	}

	got := rw.Rewrite("hello bestie")
	if got != "yo bestie 🔥" {
		t.Errorf("capture groups should work but got %q 💀", got)
	}
}

func TestRuleRewriter_EmptyRules(t *testing.T) {
	rw := &RuleRewriter{}

	line := "nothing to see here"
	got := rw.Rewrite(line)
	if got != line {
		t.Errorf("empty rules should passthrough but got %q 💀", got)
	}
}
