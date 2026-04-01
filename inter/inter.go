// Package inter embeds the Inter typeface (Regular, SemiBold weights).
package inter

import "embed"

// FS provides filesystem access to all embedded Inter TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Inter Regular (400) TTF data.
//
//go:embed Inter-Regular.ttf
var Regular []byte

// Bold is the Inter Bold (700) TTF data.
//
//go:embed Inter-Bold.ttf
var Bold []byte

// Italic is the Inter Italic (400) TTF data.
//
//go:embed Inter-Italic.ttf
var Italic []byte

// SemiBold is the Inter SemiBold (600) TTF data.
//
//go:embed Inter-SemiBold.ttf
var SemiBold []byte
