// Package crimsontext embeds the Crimson Text typeface (Regular, Bold, Italic, BoldItalic weights).
package crimsontext

import "embed"

// FS provides filesystem access to all embedded Crimson Text TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Crimson Text Regular (400) TTF data.
//
//go:embed CrimsonText-Regular.ttf
var Regular []byte

// Bold is the Crimson Text Bold (700) TTF data.
//
//go:embed CrimsonText-Bold.ttf
var Bold []byte

// Italic is the Crimson Text Italic (400) TTF data.
//
//go:embed CrimsonText-Italic.ttf
var Italic []byte

// BoldItalic is the Crimson Text Bold Italic (700) TTF data.
//
//go:embed CrimsonText-BoldItalic.ttf
var BoldItalic []byte
