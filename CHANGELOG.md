# changelog 📜

all the glow ups and level ups for litty-logs-go no cap

## [0.2.0] — 2026-03-09

### the big drop — CLI tool + JSON handler, absolute unit of a release 🔥

the litty CLI tool AND structured JSON logging both dropped in one release bestie. your Go commands now hit different AND your log aggregators are eating good no cap

#### added — litty CLI tool 🛠️
- `litty test` — `go test` but every line hits different with emojis and colors 🧪
- `litty build` — `go build` with litty error messages that actually slap 🏗️
- `litty run` — `go run` with litty compile error drip (program output passes through untouched) 🏃
- `litty vet` — `go vet` findings but the vibes are gen alpha 🔍
- `litty clean` — see whats getting yeeted in real time with `-x` auto-injection 🗑️
- auto-injects `-v` for test so output is always verbose enough to rewrite 💅
- composable rewriter architecture — test falls back to build for compile errors
- dual-stream goroutine capture — both stdout and stderr get the litty treatment
- exit code passthrough for CI/CD compatibility no cap
- installable via `go install github.com/phsk69/litty-logs-go/cmd/litty@latest` 📦
- `litty help` and `litty version` commands for the basics bestie
- 1MB scanner buffer for handling even the most unhinged test output dumps

#### added — JSONHandler 📦
- `JSONHandler` implementing `slog.Handler` — structured JSON with literal emojis for log aggregators
- `NewJSONHandler()`, `NewJSONHandlerWithOptions()`, `NewJSONLogger()` — same API shape as text handler 💅
- one compact JSON object per line with fields: timestamp, level, emoji, category, message + flat attrs
- proper JSON types — ints stay ints, bools stay bools, floats stay floats bestie
- literal emoji serialization — 🔥 not `\uD83D\uDD25`, Go is built different
- `examples/json/` — example showing all levels, groups, attrs

#### added — CI vibes 🔥
- CI pipeline now uses `litty build/vet/test` for fire emoji output in pipeline logs

## [0.1.1] - 2026-03-02

### hotfix — proxy poke timing 🔧

- fixed release pipeline poking proxy.golang.org before GitHub mirror had the tag synced (404 city 💀)
- moved proxy poke step to after GitHub mirror sync confirmation

## [0.1.0] - 2026-03-02

### the genesis drop — Go edition 🔥

the very first release of litty-logs for Go bestie. console logging with emojis, ANSI colors, and gen alpha energy. implements slog.Handler because we respect the Go way no cap

#### added
- `Handler` implementing `slog.Handler` — the core of the whole operation 🧠
- emoji-coded log levels: 👀 trace, 🔍 debug, 🔥 info, 😤 warn, 💀 err
- ANSI color codes for terminal vibrancy that hits different 🎨
- category shortening — yeets namespace bloat from logger names
- ISO 8601 UTC timestamps for that international rizz 🌍
- configurable output format: level-first (default) or timestamp-first
- `Options` struct with bussin defaults out the box
- `NewHandler()` and `NewLogger()` for easy setup
- log injection prevention — newlines in messages get sanitized 🔒
- structured attribute support (slog key=value pairs)
- `WithGroup()` for categorized loggers
- `WithAttrs()` for pre-resolved attributes
- zero external dependencies — stdlib only bestie 💅
- basic example in examples/basic/
- CI/CD with Forgejo Actions
- justfile with build, test, vet, bump, and release recipes
