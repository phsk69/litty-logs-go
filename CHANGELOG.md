# changelog 📜

all the glow ups and level ups for litty-logs-go no cap

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
