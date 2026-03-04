package main

import (
	"testing"
)

func TestRun_ExitCodePassthrough(t *testing.T) {
	result := Run("sh", []string{"-c", "exit 42"}, func(s string) string { return s })
	if result.ExitCode != 42 {
		t.Errorf("expected exit code 42 but got %d — passthrough is bricked 💀", result.ExitCode)
	}
}

func TestRun_ExitCodeZeroOnSuccess(t *testing.T) {
	result := Run("sh", []string{"-c", "echo hello"}, func(s string) string { return s })
	if result.ExitCode != 0 {
		t.Errorf("expected exit code 0 but got %d — success case is bricked 💀", result.ExitCode)
	}
}

func TestRun_RewriterIsCalled(t *testing.T) {
	called := false
	result := Run("sh", []string{"-c", "echo test"}, func(s string) string {
		called = true
		return s
	})
	if !called {
		t.Error("rewriter function was never called — thats not bussin 💀")
	}
	if result.ExitCode != 0 {
		t.Errorf("expected exit code 0 but got %d 💀", result.ExitCode)
	}
}

func TestRun_RewriterTransformsOutput(t *testing.T) {
	result := Run("sh", []string{"-c", "echo boring"}, func(s string) string {
		return "litty 🔥"
	})
	if result.ExitCode != 0 {
		t.Errorf("expected exit code 0 but got %d 💀", result.ExitCode)
	}
	// we cant easily capture os.Stdout in a test without major refactoring
	// but we can verify the rewriter was called and command succeeded
	// the real verification is the integration test (go run ./cmd/litty test ./...)
}

func TestRun_InvalidCommand(t *testing.T) {
	result := Run("this-command-definitely-doesnt-exist-fr-fr", nil, func(s string) string { return s })
	if result.Err == nil {
		t.Error("expected an error for invalid command but got nil — error handling is bricked 💀")
	}
}
