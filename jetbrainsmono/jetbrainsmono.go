// Package jetbrainsmono embeds the JetBrains Mono typeface (variable weight axis).
package jetbrainsmono

import "embed"

// FS provides filesystem access to all embedded JetBrains Mono TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Variable is the JetBrains Mono variable-weight TTF data. Upstream ships this
// family as a single variable font rather than per-weight statics, so consumers
// select a weight on the wght axis at runtime.
//
//go:embed JetBrainsMono-VariableFont_wght.ttf
var Variable []byte
