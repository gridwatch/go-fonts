// Package robotoslab embeds the Roboto Slab typeface (Regular, Bold weights).
package robotoslab

import "embed"

// FS provides filesystem access to all embedded Roboto Slab TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Roboto Slab Regular (400) TTF data.
//
//go:embed RobotoSlab-Regular.ttf
var Regular []byte

// Bold is the Roboto Slab Bold (700) TTF data.
//
//go:embed RobotoSlab-Bold.ttf
var Bold []byte
