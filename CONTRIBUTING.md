# Contributing

Thanks for taking the time to contribute.

## Conventions

[AGENTS.md](AGENTS.md) is the authoritative reference for how code here is
written, tested and structured. Read it before opening a pull request — it
carries the rules that are not obvious from the source, particularly around the
dual `//go:embed` pattern and the per-package test assertions.

## Local setup

Most development tasks run through [Task](https://taskfile.dev). The Taskfile pulls shared
 task definitions from the `tooling` repository, checked out beside this one:
```shell
git clone https://github.com/gridwatch/tooling.git ../tooling
task test    # go test with the race detector
task lint    # gofmt, golangci-lint, govulncheck
```

If you do not have access to `tooling`, `go test -race ./...` is a reasonable substitute
for `task test` and is enough to validate most changes.

## Adding or updating a typeface

This is the part most likely to go wrong, so it has hard rules.

- **Read the licence out of the font, not off a specimen page.** Every entry in
  `LICENSE-FONTS` comes from the TTF's own `name` table — name ID 0 for the
  copyright notice, 13 and 14 for the licence. Specimen pages go stale: Roboto
  was recorded here as Apache 2.0 long after Google relicensed it to OFL.
- **Ship the licence text with the font.** Each family directory carries its own
  `OFL.txt` (SIL OFL families) or `LICENSE.txt` (Apache families), headed by that
  family's copyright notice. OFL 1.1 clause 2 requires the notice and licence to
  travel with the fonts; a link is not enough.
- **Take every weight from one upstream release.** Mixing cuts of the same family
  produces a package whose weights do not match each other — `inter` shipped
  Bold and Italic from Google Fonts' optical-size split alongside Regular from
  rsms's release, and the metrics disagreed.
- **Do not modify, subset or re-hint the files.** Several families carry a
  Reserved Font Name, which under OFL 1.1 clause 3 a modified version may not
  use. Unmodified upstream releases sidestep this entirely.

The step-by-step checklist is in [AGENTS.md](AGENTS.md) under *Adding a Font*.

## Pull requests

- Commit subjects follow [Conventional Commits](https://www.conventionalcommits.org):
  `feat:`, `fix:`, `docs:`, `chore:`.
- Keep CI green. Pull requests are squash-merged, so the PR title becomes the
  commit subject on `main`.
- Some files in this repository are rendered centrally by Terraform and are
  overwritten on the next apply — `.gitattributes`, `.github/workflows/`,
  `.github/CODEOWNERS`, `.github/copilot-instructions.md` and `SECURITY.md`.
  Editing them here has no lasting effect; raise the change against the
  template instead.

## Security

Please do not open a public issue for security problems. See
[SECURITY.md](SECURITY.md) for the disclosure process.
