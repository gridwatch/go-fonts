# go-fonts

> Go module embedding the GRIDWATCH typefaces, one package per family

Each family is its own package, so a consumer compiles in only the bytes it asks
for rather than the whole set. The fonts are TTF because the consumers
rasterise text into images; a browser wants woff2 and should not use this module.

## Install

```shell
go get github.com/gridwatch/go-fonts
```

## Usage

```go
import "github.com/gridwatch/go-fonts/inter"

data := inter.Regular   // []byte — raw TTF
fsys := inter.FS        // embed.FS — for walking or extracting
```

Static families export one `[]byte` per weight. Variable families export a
single `Variable` and leave axis selection to the caller.

| Package | Family | Exports |
|---|---|---|
| `barlowcondensed` | Barlow Condensed | `SemiBold`, `Bold` |
| `courierprime` | Courier Prime | `Regular`, `Italic`, `Bold` |
| `crimsontext` | Crimson Text | `Regular`, `Italic`, `Bold`, `BoldItalic` |
| `ebgaramond` | EB Garamond | `Regular`, `Italic`, `Bold`, `BoldItalic` |
| `ibmplexmono` | IBM Plex Mono | `Regular`, `Italic`, `Bold` |
| `inter` | Inter | `Regular`, `Medium`, `SemiBold`, `Bold`, `Italic` |
| `jetbrainsmono` | JetBrains Mono | `Variable` (wght) |
| `lato` | Lato | `Regular`, `Italic`, `Bold`, `BoldItalic` |
| `oswald` | Oswald | `Variable` (wght) |
| `roboto` | Roboto | `Variable` (wdth, wght) |
| `robotoslab` | Roboto Slab | `Regular`, `Bold` |
| `specialelite` | Special Elite | `Regular` |

Every package also exports `FS embed.FS` covering that family's TTF files.

## Licensing

The Go code in this repository is Apache 2.0 — see [LICENSE](LICENSE).

The embedded typefaces are third-party and keep their own licences, either the
SIL Open Font License 1.1 or the Apache License 2.0 depending on the family.
Each package directory carries the full licence text next to its TTF files
(`OFL.txt` or `LICENSE.txt`), and [LICENSE-FONTS](LICENSE-FONTS) indexes them
with the copyright notice read from each font's own metadata.

Both licences permit redistribution and embedding in applications, including
commercial ones, **provided the copyright notice and licence text stay with the
font files**. Keep those files in place when vendoring a package.

The fonts here are unmodified upstream releases. Lato carries the Reserved Font
Name "Lato", so leave it that way — under OFL 1.1 clause 3 a modified version
may not use the reserved name.

## Development

```shell
task test
task lint
```

See [AGENTS.md](AGENTS.md) for conventions and
[CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.
