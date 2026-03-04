package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

// maxScanBuf is the max line size for the scanner — 1MB should handle
// even the most unhinged test output dumps bestie 📦
const maxScanBuf = 1024 * 1024

// RunResult holds the exit code and any runner-level error
// this is what you get back after the subprocess finishes cooking 🍳
type RunResult struct {
	ExitCode int
	Err      error
}

// Run spawns a command, rewrites both stdout and stderr line-by-line through
// the rewriter function, and returns the exit code.
// stdout lines go to os.Stdout, stderr lines go to os.Stderr.
// non-matching lines pass through unchanged — safe for any output bestie 🔥
func Run(name string, args []string, rewrite func(string) string) RunResult {
	cmd := exec.Command(name, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return RunResult{ExitCode: 1, Err: fmt.Errorf("failed to pipe stdout: %w — thats not bussin 💀", err)}
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return RunResult{ExitCode: 1, Err: fmt.Errorf("failed to pipe stderr: %w — thats not bussin 💀", err)}
	}

	if err := cmd.Start(); err != nil {
		return RunResult{ExitCode: 1, Err: fmt.Errorf("failed to start %s: %w — command is bricked 💀", name, err)}
	}

	var wg sync.WaitGroup

	// goroutine for stdout — rewrite and send to os.Stdout 🔥
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdoutPipe)
		scanner.Buffer(make([]byte, 0, maxScanBuf), maxScanBuf)
		for scanner.Scan() {
			fmt.Fprintln(os.Stdout, rewrite(scanner.Text()))
		}
	}()

	// goroutine for stderr — rewrite and send to os.Stderr 🔥
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderrPipe)
		scanner.Buffer(make([]byte, 0, maxScanBuf), maxScanBuf)
		for scanner.Scan() {
			fmt.Fprintln(os.Stderr, rewrite(scanner.Text()))
		}
	}()

	// wait for both streams to finish reading before calling Wait
	// otherwise we might miss output bestie 💀
	wg.Wait()

	err = cmd.Wait()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return RunResult{ExitCode: 1, Err: fmt.Errorf("command wait failed: %w 💀", err)}
		}
	}

	return RunResult{ExitCode: exitCode}
}
