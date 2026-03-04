package main

import (
	"regexp"
)

// Rewriter transforms a line of command output into gen alpha energy 🔥
// each Go command gets its own implementation that knows how to parse
// and rewrite that commands specific output format no cap
type Rewriter interface {
	// Rewrite transforms a single line of output.
	// returns the rewritten line, or the original line if no rule matched.
	// NEVER returns empty string for non-empty input — we dont yeet lines, we transform em 💅
	Rewrite(line string) string
}

// Rule is a single regex-to-replacement mapping — the atomic unit of rewriting bestie 🧬
type Rule struct {
	Pattern *regexp.Regexp
	Replace func(matches []string) string
}

// RuleRewriter applies rules in order, first match wins.
// if nothing matches, falls back to an optional parent rewriter.
// this is the composable chain that makes the whole architecture bussin 💅
type RuleRewriter struct {
	Rules    []Rule
	Fallback Rewriter
}

// Rewrite applies the first matching rule, or falls back, or passes through unchanged
func (r *RuleRewriter) Rewrite(line string) string {
	for _, rule := range r.Rules {
		if matches := rule.Pattern.FindStringSubmatch(line); matches != nil {
			return rule.Replace(matches)
		}
	}
	if r.Fallback != nil {
		return r.Fallback.Rewrite(line)
	}
	return line
}
