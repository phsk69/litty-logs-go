package main

import (
	"fmt"
	"regexp"

	litty "github.com/phsk69/litty-logs-go"
)

// vet rewriter patterns — go vet findings but make them gen alpha 🔍
var vetRules = []Rule{
	{
		// package header gets a vet-specific prefix 🔍
		Pattern: regexp.MustCompile(`^# (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("🔍 vetting the vibes in %s...", m[1])
		},
	},
	{
		// vet finding — file:line:col: message 😤
		Pattern: regexp.MustCompile(`^(.+):(\d+):(\d+): (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s😤 %s:%s:%s — vet says: %s%s",
				litty.Yellow, m[1], m[2], m[3], m[4], litty.Reset)
		},
	},
}

// NewVetRewriter creates a rewriter for go vet output 🔍
// falls back to build rewriter just in case 💅
func NewVetRewriter() Rewriter {
	return &RuleRewriter{
		Rules:    vetRules,
		Fallback: NewBuildRewriter(),
	}
}
