package main

import (
	"fmt"
	"regexp"
	"strings"

	litty "github.com/phsk69/litty-logs-go"
)

// test rewriter patterns — the BIGGEST rewriter, handles all go test -v output 🧪
// falls back to build rewriter for compile errors that show up in test output
var testRules = []Rule{
	{
		// === RUN — a test is about to cook 🏃
		Pattern: regexp.MustCompile(`^=== RUN\s+(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("🏃 running %s...", strings.TrimSpace(m[1]))
		},
	},
	{
		// --- PASS — test absolutely slayed 🔥
		Pattern: regexp.MustCompile(`^--- PASS: (.+) \((.+)\)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s✅ %s absolutely slayed (%s) 🔥%s",
				litty.Green, m[1], m[2], litty.Reset)
		},
	},
	{
		// --- FAIL — test took a fat L 💀
		Pattern: regexp.MustCompile(`^--- FAIL: (.+) \((.+)\)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s💀 %s took a fat L (%s)%s",
				litty.Red, m[1], m[2], litty.Reset)
		},
	},
	{
		// --- SKIP — test said not today 🔥
		Pattern: regexp.MustCompile(`^--- SKIP: (.+) \((.+)\)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s⏭️ %s said \"not today bestie\" (%s)%s",
				litty.Yellow, m[1], m[2], litty.Reset)
		},
	},
	{
		// === PAUSE — test on pause, chilling ⏸️
		Pattern: regexp.MustCompile(`^=== PAUSE\s+(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("⏸️ %s on pause, chilling for a sec...", strings.TrimSpace(m[1]))
		},
	},
	{
		// === CONT — test back in the game ▶️
		Pattern: regexp.MustCompile(`^=== CONT\s+(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("▶️ %s back in the game lesgooo", strings.TrimSpace(m[1]))
		},
	},
	{
		// standalone PASS — all tests bussin 🏆
		Pattern: regexp.MustCompile(`^PASS$`),
		Replace: func(_ []string) string {
			return fmt.Sprintf("%s🏆 W in the chat, all tests bussin fr fr 🔥%s",
				litty.Green, litty.Reset)
		},
	},
	{
		// standalone FAIL — tests took a massive L 💀
		Pattern: regexp.MustCompile(`^FAIL$`),
		Replace: func(_ []string) string {
			return fmt.Sprintf("%s💀 tests took a massive L, not bussin at all%s",
				litty.Red, litty.Reset)
		},
	},
	{
		// ok package summary — package vibed clean 💅
		Pattern: regexp.MustCompile(`^ok\s+(\S+)\s+(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s✅ %s vibed in %s — clean 💅%s",
				litty.Green, m[1], m[2], litty.Reset)
		},
	},
	{
		// FAIL package summary — package is cooked 💀
		Pattern: regexp.MustCompile(`^FAIL\s+(\S+)\s+(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s💀 %s is absolutely cooked (%s)%s",
				litty.Red, m[1], m[2], litty.Reset)
		},
	},
	{
		// no test files — living dangerously 🤷
		Pattern: regexp.MustCompile(`^\?\s+(\S+)\s+\[no test files\]$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("🤷 %s has no tests — living dangerously bestie", m[1])
		},
	},
	{
		// coverage line — are we eating? 📊
		Pattern: regexp.MustCompile(`^coverage: (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("📊 coverage check: %s — are we eating?", m[1])
		},
	},
	{
		// testing warning — heads up bestie 😤
		Pattern: regexp.MustCompile(`^testing: warning: (.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("😤 heads up bestie: %s", m[1])
		},
	},
	{
		// indented test output (4+ spaces) — add arrow for visual hierarchy ↳
		Pattern: regexp.MustCompile(`^(\s{4,})(.+)$`),
		Replace: func(m []string) string {
			return fmt.Sprintf("%s↳ %s", m[1], m[2])
		},
	},
}

// NewTestRewriter creates a rewriter for go test output 🧪
// falls back to build rewriter for compile errors — composable chain bestie 💅
func NewTestRewriter() Rewriter {
	return &RuleRewriter{
		Rules:    testRules,
		Fallback: NewBuildRewriter(),
	}
}
