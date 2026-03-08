package main

import (
	"fmt"
	"regexp"
)

// clean rewriter patterns — see whats getting yeeted 🗑️
var cleanRules = []Rule{
	{
		// rm -f lines — file getting yeeted 🗑️
		Pattern: regexp.MustCompile(`^rm -f (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("🗑️ yeeted: %s", m[1])
		},
	},
	{
		// cd lines — entering directory to clean 📂
		Pattern: regexp.MustCompile(`^cd (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("📂 cleaning vibes in %s...", m[1])
		},
	},
}

// NewCleanRewriter creates a rewriter for go clean -x output 🗑️
func NewCleanRewriter() Rewriter {
	return &RuleRewriter{Rules: cleanRules}
}
