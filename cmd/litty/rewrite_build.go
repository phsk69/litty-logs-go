package main

import (
	"fmt"
	"regexp"

	litty "github.com/phsk69/litty-logs-go"
)

// build rewriter patterns — compiled once, used forever 🔥
// order matters bestie — specific patterns before general ones
var buildRules = []Rule{
	{
		// package header — "# github.com/foo/bar"
		// this is the first thing go build spits out when something is cooked 📦
		Pattern: regexp.MustCompile(`^# (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("📦 checking vibes in %s...", m[1])
		},
	},
	{
		// imported and not used — bestie why you importing stuff you dont need 🗑️
		Pattern: regexp.MustCompile(`^(.+):(\d+):(\d+): imported and not used: "(.+)"$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s💀 %s:%s:%s — bestie you imported %q and never used it, thats wasteful fr 🗑️%s",
				litty.Red, m[1], m[2], m[3], m[4], litty.Reset)
		},
	},
	{
		// declared and not used — yeet it bestie 🗑️
		Pattern: regexp.MustCompile(`^(.+):(\d+):(\d+): (.+ declared (and not used|but not used))$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s💀 %s:%s:%s — you declared this and never used it, yeet it bestie 🗑️%s",
				litty.Red, m[1], m[2], m[3], litty.Reset)
		},
	},
	{
		// general compile error — the catch-all for file:line:col: message patterns 💀
		Pattern: regexp.MustCompile(`^(.+):(\d+):(\d+): (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s💀 big L at %s:%s:%s — %s%s",
				litty.Red, m[1], m[2], m[3], m[4], litty.Reset)
		},
	},
}

// NewBuildRewriter creates a rewriter for go build output 🏗️
// this is the base rewriter that other rewriters fall back to
func NewBuildRewriter() Rewriter {
	return &RuleRewriter{Rules: buildRules}
}
