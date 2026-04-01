// Package lato embeds the Lato typeface (Regular, Bold, Italic, BoldItalic weights).
package lato

import "embed"

// FS provides filesystem access to all embedded Lato TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Lato Regular (400) TTF data.
//
//go:embed Lato-Regular.ttf
var Regular []byte

// Bold is the Lato Bold (700) TTF data.
//
//go:embed Lato-Bold.ttf
var Bold []byte

// Italic is the Lato Italic (400) TTF data.
//
//go:embed Lato-Italic.ttf
var Italic []byte

// BoldItalic is the Lato Bold Italic (700) TTF data.
//
//go:embed Lato-BoldItalic.ttf
var BoldItalic []byte
