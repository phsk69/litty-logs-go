package main

import (
	"fmt"
	"os"
)

// version gets set by ldflags at build time, or defaults to dev 🏷️
var version = "dev"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	var rewriter Rewriter
	var goCmd string
	var successMsg string
	var injectArgs []string

	switch subcommand {
	case "test":
		rewriter = NewTestRewriter()
		goCmd = "test"
		// go test prints PASS/ok lines, rewriter handles those — no extra success msg needed
		successMsg = ""
		// auto-inject -v if not present — gotta have verbose output to rewrite bestie 🔥
		if !containsFlag(args, "-v") {
			injectArgs = []string{"-v"}
		}
	case "build":
		rewriter = NewBuildRewriter()
		goCmd = "build"
		successMsg = "✅ build absolutely slayed bestie, zero errors lets gooo 🏆🔥"
	case "run":
		rewriter = NewRunRewriter()
		goCmd = "run"
		// go run output IS the program — no success msg needed
		successMsg = ""
	case "vet":
		rewriter = NewVetRewriter()
		goCmd = "vet"
		successMsg = "✅ vet found zero issues, code is immaculate bestie 💅🔥"
	case "clean":
		rewriter = NewCleanRewriter()
		goCmd = "clean"
		successMsg = "🗑️🔥 all artifacts yeeted into the void — clean af bestie"
		// auto-inject -x if not present — gotta see whats getting yeeted 🗑️
		if !containsFlag(args, "-x") {
			injectArgs = []string{"-x"}
		}
	case "help", "--help", "-h":
		printUsage()
		os.Exit(0)
	case "version", "--version":
		printVersion()
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "💀 bruh %q is not a valid litty command — try `litty help` bestie\n", subcommand)
		os.Exit(1)
	}

	// build final args: [go_subcommand] + injected flags + user args
	finalArgs := []string{goCmd}
	finalArgs = append(finalArgs, injectArgs...)
	finalArgs = append(finalArgs, args...)

	// lil header so you know litty is cooking 🔥
	fmt.Fprintf(os.Stderr, "🔥 litty %s — lets cook bestie\n", subcommand)

	result := Run("go", finalArgs, rewriter.Rewrite)

	if result.Err != nil {
		fmt.Fprintf(os.Stderr, "💀 %v\n", result.Err)
		os.Exit(1)
	}

	// print success message if the command succeeded and we have one
	// go build/vet/clean produce NO stdout on success (Go moment fr fr 💀)
	if result.ExitCode == 0 && successMsg != "" {
		fmt.Println(successMsg)
	}

	os.Exit(result.ExitCode)
}

// containsFlag checks if a flag is already in the args list
// we dont wanna double-inject flags, thats not bussin 💀
func containsFlag(args []string, flag string) bool {
	for _, a := range args {
		if a == flag {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Print(`🔥 litty — go commands but make them gen alpha fr fr

usage: litty <command> [args...]

commands:
  test     wraps 'go test' with bussin output (auto-injects -v) 🧪
  build    wraps 'go build' with litty error messages 🏗️
  run      wraps 'go run' with litty compile errors 🏃
  vet      wraps 'go vet' with gen alpha energy 🔍
  clean    wraps 'go clean' and shows whats getting yeeted 🗑️
  help     shows this bussin help text 💅
  version  shows the version bestie 🏷️

examples:
  litty test ./...              run all tests with litty output 🧪
  litty test -run TestFoo       run specific test 🎯
  litty build ./cmd/myapp       build with litty errors 🏗️
  litty run ./cmd/myapp         run with litty compile errors 🏃
  litty vet ./...               vet with gen alpha vibes 🔍
  litty clean                   see whats getting yeeted 🗑️

all args after the command get passed straight to the go tool 🔥
install: go install github.com/phsk69/litty-logs-go/cmd/litty@latest
`)
}

func printVersion() {
	fmt.Printf("litty %s — the most bussin Go CLI no cap 🔥\n", version)
}
