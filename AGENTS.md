# AGENTS.md — go-fonts

Shared Go module embedding typeface TTF files. Each font family is a separate package so consumers import only what they need.

---

## Table of Contents

1. [Usage](#usage)
2. [Build & Test](#build--test)
3. [Conventions](#conventions)
4. [Adding a Font](#adding-a-font)
5. [Gotchas](#gotchas)

---

## Usage

```go
import "github.com/gridwatch/go-fonts/inter"

data := inter.Regular      // []byte — raw TTF
fs := inter.FS             // embed.FS — for walking/extracting
```

## Build & Test

Always use Taskfile commands. Tests use `github.com/stretchr/testify` (`require.*` for fatal, `assert.*` for soft). Linting uses shared config at `../tooling/golangci.yml`.

## Conventions

- **One package per font family** — directory name is lowercase, no hyphens.
- **Dual embed pattern** — each package embeds via both `//go:embed *.ttf` (on `FS embed.FS`) and per-file embeds (on named `[]byte` vars per weight). Both must stay in sync.
- **Variable fonts export `Variable` only** — packages like `oswald` and `roboto` export a single `Variable []byte`, not per-weight vars. Consumers handle axis ranges at runtime.
- **Per-package test isolation** — each package defines its own `assertParseable()` helper locally rather than sharing a test utility. This is intentional — packages are independent.
- **Per-weight test functions** — multi-weight packages (e.g. lato) define individual `TestRegular()`, `TestBold()`, etc. Variable font packages use only `TestVariable()`.
- **File count assertion via `FS.ReadDir(".")`** — tests verify the exact number of embedded TTF files matches expectations. Catches accidental additions.

## Adding a Font

1. Create a new package directory (lowercase, no hyphens).
2. Add TTF files to the directory. Take every weight from the **same upstream release** — see the Gotchas.
3. Add the family's licence text beside them: `OFL.txt` for SIL OFL families, `LICENSE.txt` for Apache ones. Head it with the copyright notice read out of the font's own `name` table (name ID 0) — not from a specimen page.
4. Add the family to `LICENSE-FONTS`, with the licence taken from name IDs 13/14 of the font itself.
5. Create a `.go` file with `//go:embed` directives exposing `FS embed.FS` and named `[]byte` vars per weight.
6. Add a `_test.go` with an `assertParseable(t, data)` helper verifying non-empty bytes and valid TTF headers.
7. Add per-weight test functions (`TestRegular`, `TestBold`, etc.) or `TestVariable` for variable fonts.
8. **Update the `assert.Len` count** in the test to match the number of TTF files in the directory.

The licence files sit in the package directory but are not embedded — `FS` is `//go:embed *.ttf`, so adding them does not move the `assert.Len` count.

## Web consumers use woff2, not this module

This module is TTF for **rasterisation** — its tests parse every file through
`golang.org/x/image/font/opentype`, because the consumers draw text into images.

A browser wants woff2, which is roughly a seventeenth of the bytes (Inter Regular:
402 KB TTF against 23 KB woff2). So a web consumer does **not** import these packages and serve the bytes.
`backbrief` used to vendor woff2 of the same families; it now links them from
Google Fonts instead, so there is no longer a copy to keep in step.

What this module owns is which typefaces are house typefaces and what their
licence is — not the delivery format for every consumer.

**`barlowcondensed` and `jetbrainsmono` currently have no importer.** Both were
added for `backbrief` before it moved to Google Fonts. They are kept as the
canonical statement that these are house typefaces, which is this module's stated
job; delete them only if that job is being narrowed to "fonts something
rasterises", in which case say so here.

## Gotchas

- **File count must match** — adding a TTF without updating the `assert.Len` count fails tests. Adding a TTF without a corresponding `[]byte` var still includes it in `FS`, which also triggers the count assertion.
- **No shared test utility** — each package's `assertParseable()` is local. If you change the assertion pattern, update all packages.
- **The font's own metadata is authoritative on licensing, not the specimen page.** `LICENSE-FONTS` recorded Roboto as Apache 2.0 while the embedded file's `name` table said OFL 1.1 — Google relicensed the family and the summary was never revisited. Roboto Slab genuinely is still Apache, which makes the two easy to conflate. Read name IDs 0, 13 and 14 out of the TTF before writing a licence into `LICENSE-FONTS`.
- **Keep a family's statics on one upstream cut.** `inter` shipped Bold and Italic from Google Fonts' optical-size split (family `Inter 24pt`, `git-66647c0bb`) alongside Regular/Medium/SemiBold from rsms's plain `Inter` (`git-9221beed3`) — two different designs with different metrics in one package, so `inter.Bold` did not match `inter.Regular`. Check the `name` table's family and version strings agree across every file in a directory.
- **CI delegates to shared workflows** — the test, lint and Scorecard callers reference `gridwatch/.github/.github/workflows/`. No inline CI logic in this repo.
