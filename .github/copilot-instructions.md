# GitHub Copilot Instructions

This repository is a Go project in the [@gridwatch](https://github.com/gridwatch) fleet. `AGENTS.md` in the repository root carries the full conventions and is authoritative wherever it is more specific than this file.

## Centrally managed files

`.github/workflows/` and `.gitattributes` are rendered from templates in the `infrastructure-github` repository and pushed by Terraform. Editing them here is reverted on the next apply. Propose the change against the template instead, and say so when reviewing a pull request that touches one.

## Build, test and lint

Everything runs through Taskfile targets — `task build`, `task test`, `task lint`, `task run -- <subcommand>`. **Never run bare `go build`, `go test`, `go run` or `golangci-lint`.** If a needed command has no task, add it to the Taskfile first, then call it.

Operational checks go through the Taskfile too (`task status`, `task health`, `task restart`) rather than bare `curl`, `lsof`, `kill` or `pkill`.

Go 1.26 or later. Format with `gofmt`. `golangci-lint` uses the shared allowlist config in the sibling `tooling` repo (`tooling/golangci.yml`) and CI requires zero warnings.

## Testing

- Assertions use `github.com/stretchr/testify`: `require.*` for fatal checks, `assert.*` for soft checks.
- **Never write manual `if got != want { t.Errorf(...) }` comparisons.** `testifylint` is enabled in `golangci-lint` to enforce this.
- Prefer table-driven tests for anything with more than a couple of input cases.
- Mock external HTTP APIs with fixtures so tests make no network calls in CI.

## Coding style

- Standard Go naming: `MixedCaps`, not `snake_case`. Exported types and functions carry doc comments; unexported helpers only when the logic isn't self-evident.
- Use descriptive variable names. A short receiver-like name is fine in a tight local scope; a single letter for a long-lived value is not.
- **Always handle errors.** Never discard one with `_` unless a comment explains why.
- Wrap errors with context: `fmt.Errorf("failed to generate manifest: %w", err)`. Return them to the caller; log-and-continue only at the top-level handler.
- Use sentinel errors (`var ErrUnitNotFound = errors.New(...)`) for conditions callers need to check.
- Prefer the standard library. A new dependency needs justification, not just convenience.

## Concurrency

- Protect shared state with `sync.Mutex` or `sync.RWMutex`, and hold the lock for the minimum necessary duration — never across I/O.
- `context.Context` is the first parameter of any function that does I/O or may block.
- Pass context into goroutines and respect cancellation. Use `errgroup` for parallel fetches.

## Configuration

Configuration lives in `config.yml`, loaded via Viper, with `GRIDWATCH_`-prefixed environment overrides. Precedence is CLI flags > environment variables > config file > defaults. Read `config.yml` for the current keys rather than duplicating them into documentation.

## Container images

Image tags are **date-based**, `vYYYYMMDD-N` — never semantic versions. **Never push `latest` to the container registry**; the `latest` tag in a local Taskfile is for local development builds only.

## Module versions

**A repo consumed as a Go module tags semver instead** — `v0.1.0`, not `vYYYYMMDD-N`. This is the one exception to the date-based rule, and it is not a preference: the module proxy resolves versions by semver and rejects a date-based tag outright, so `go get` cannot see one.

At `v2` and above the module path takes a `/v2` suffix to match.

Private modules need `GOPRIVATE=github.com/gridwatch/*` locally as well as in CI. Without it `go mod download` asks the public proxy and the checksum database for a repo neither can see, and fails rather than falling back to a direct fetch.

## GitHub Actions

- Use `actions/checkout@v6` or later. Never `actions/checkout@v4`.
- Pin every action to an exact version, with a `# see <releases URL>` comment above the pin.

## Documentation

- Diagrams are **Mermaid** only. No ASCII box drawing.
- Markdown is linted with `markdownlint` at a 320-character line limit — don't hard-wrap prose at 80.
- `CLAUDE.md` is a thin entry point: a title line and `@AGENTS.md`, nothing else. Conventions belong in `AGENTS.md`.
