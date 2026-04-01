# AGENTS.md — go-fonts

Shared Go module embedding typeface TTF files. Each font family is a separate package so consumers import only what they need.

## Usage

```go
import "github.com/gridwatch/go-fonts/inter"

data := inter.Regular      // []byte — raw TTF
fs := inter.FS             // embed.FS — for walking/extracting
```

## Packages

| Package | Typeface | Weights |
|---------|----------|---------|
| `courierprime` | Courier Prime | Regular, Bold, Italic |
| `crimsontext` | Crimson Text | Regular, Bold, Italic, BoldItalic |
| `ebgaramond` | EB Garamond | Regular, Bold, Italic, BoldItalic |
| `ibmplexmono` | IBM Plex Mono | Regular, Bold, Italic |
| `inter` | Inter | Regular, Bold, Italic, SemiBold |
| `lato` | Lato | Regular, Bold, Italic, BoldItalic |
| `oswald` | Oswald | Variable (weight 200–700) |
| `roboto` | Roboto | Variable (width + weight) |
| `robotoslab` | Roboto Slab | Regular, Bold |
| `specialelite` | Special Elite | Regular |

## Build & Test

Always use Taskfile commands.

| What | Command |
|------|---------|
| Test | `task test` |
| Lint | `task lint` |

## Adding a Font

1. Create a new package directory (lowercase, no hyphens).
2. Add TTF files to the directory.
3. Create a `.go` file with `//go:embed` directives exposing `FS embed.FS` and named `[]byte` vars per weight.
4. Add a `_test.go` verifying non-empty bytes and valid TTF headers.
