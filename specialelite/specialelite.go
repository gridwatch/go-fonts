// Package specialelite embeds the Special Elite typeface (Regular weight).
package specialelite

import "embed"

// FS provides filesystem access to all embedded Special Elite TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Special Elite Regular (400) TTF data.
//
//go:embed SpecialElite-Regular.ttf
var Regular []byte
