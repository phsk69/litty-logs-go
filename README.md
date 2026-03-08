# litty-logs-go 🔥

the most bussin Go logging library no cap. transforms your boring `slog` output into gen alpha energy with emojis, ANSI colors, and abbreviated categories fr fr

## before vs after 💀➡️🔥

**before (boring corporate slog):**
```
2026/03/02 21:45:00 INFO request received method=GET path=/api/vibes
```

**after (litty-logs):**
```
[🔥 info] [2026-03-02T21:45:00.420Z] [app] request received method=GET path=/api/vibes
```

## install 📦

```bash
go get github.com/phsk69/litty-logs-go
```

## usage 🔥

### one-liner setup — bussin out the box

```go
package main

import "github.com/phsk69/litty-logs-go"

func main() {
    logger := litty.NewLogger()
    logger.Info("we vibing fr fr", "key", "value")
}
```

### override only what you want — defaults stay bussin 💅

```go
// just turn off colors — UTC, shortening, everything else stays fire
logger := litty.NewLogger(litty.WithColors(false))

// stack em like a combo meal bestie
logger := litty.NewLogger(
    litty.WithColors(false),
    litty.WithTimestampFirst(true),
    litty.WithLevel(slog.LevelDebug),
)
```

### categorized loggers via WithGroup 📦

```go
logger := litty.NewLogger()
serviceLogger := logger.WithGroup("MyService")
serviceLogger.Info("service is cooking bestie")
// output: [🔥 info] [2026-03-02T21:45:00.420Z] [MyService] service is cooking bestie
```

### set as default slog logger 🌍

```go
slog.SetDefault(litty.NewLogger())
slog.Info("the whole app is litty now bestie")
```

### timestamp-first mode for observability besties 📊

```go
logger := litty.NewLogger(litty.WithTimestampFirst(true))
// output: [2026-03-02T21:45:00.420Z] [🔥 info] [app] timestamp comes first bestie
```

### full struct control for power users 👑

```go
opts := litty.DefaultOptions()
opts.UseColors = false
opts.TimestampFirst = true
logger := slog.New(litty.NewHandlerWithOptions(opts))
```

## CLI tool 🔥

litty also comes as a CLI tool that wraps Go commands with gen alpha output. install it and never look at boring `go test` output again bestie

### install

```bash
go install github.com/phsk69/litty-logs-go/cmd/litty@latest
```

### usage

```bash
litty test ./...              # go test but bussin 🧪
litty test -run TestFoo       # run specific test 🎯
litty build ./cmd/myapp       # go build but litty 🏗️
litty run ./cmd/myapp         # go run with litty compile errors 🏃
litty vet ./...               # go vet with gen alpha vibes 🔍
litty clean                   # see whats getting yeeted 🗑️
```

all args after the command get passed straight to the go tool. auto-injects `-v` for test and `-x` for clean so theres always output to rewrite 💅

### before vs after 💀➡️🔥

**before (boring go test):**
```
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
=== RUN   TestBar
--- FAIL: TestBar (0.01s)
    bar_test.go:15: expected 5, got 3
FAIL
```

**after (litty test):**
```
🔥 litty test — lets cook bestie
🏃 running TestFoo...
✅ TestFoo absolutely slayed (0.00s) 🔥
🏃 running TestBar...
💀 TestBar took a fat L (0.01s)
    ↳ bar_test.go:15: expected 5, got 3
💀 tests took a massive L, not bussin at all
```

## log levels 🎯

| level | emoji | label | color | vibe |
|-------|-------|-------|-------|------|
| trace | 👀 | trace | cyan | for when you lowkey wanna see everything |
| debug | 🔍 | debug | blue | investigating whats going on under the hood |
| info | 🔥 | info | green | everything is bussin and vibing |
| warn | 😤 | warn | yellow | something kinda sus but we not panicking |
| error | 💀 | err | red | something took a fat L |

## options 🎛️

| option | type | default | what it does |
|--------|------|---------|-------------|
| `Level` | `slog.Leveler` | `slog.LevelInfo` | minimum log level — anything below gets yeeted 🗑️ |
| `UseColors` | `bool` | `true` | ANSI color codes for terminal vibrancy 🎨 |
| `ShortenCategories` | `bool` | `true` | yeets namespace bloat from category names |
| `TimestampFirst` | `bool` | `false` | put timestamp before level (observability style) |
| `UseUtcTimestamp` | `bool` | `true` | UTC timestamps for international rizz 🌍 |
| `Writer` | `io.Writer` | `os.Stderr` | where the output goes bestie |

## development 🛠️

```bash
# build
just build

# test with race detector
just test

# lint
just vet

# run the example
just example

# bump version
just bump patch    # 0.1.0 -> 0.1.1
just bump minor    # 0.1.0 -> 0.2.0
just bump major    # 0.1.0 -> 1.0.0

# gitflow release
just release patch
```

## roadmap 🗺️

- [x] CLI tool (`go install` litty command for litty-fied builds/tests)
- [ ] JSON structured output (litty-json)
- [ ] file sink with rotation and compression
- [ ] message rewriting (Go framework messages → gen alpha slang)
- [ ] webhook sinks (Matrix, Teams)

## license 📄

MIT — do whatever you want bestie, just keep the vibes going 💅
