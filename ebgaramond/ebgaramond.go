// Package ebgaramond embeds the EB Garamond typeface (Regular, Bold, Italic, BoldItalic weights).
package ebgaramond

import "embed"

// FS provides filesystem access to all embedded EB Garamond TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the EB Garamond Regular (400) TTF data.
//
//go:embed EBGaramond-Regular.ttf
var Regular []byte

// Bold is the EB Garamond Bold (700) TTF data.
//
//go:embed EBGaramond-Bold.ttf
var Bold []byte

// Italic is the EB Garamond Italic (400) TTF data.
//
//go:embed EBGaramond-Italic.ttf
var Italic []byte

// BoldItalic is the EB Garamond Bold Italic (700) TTF data.
//
//go:embed EBGaramond-BoldItalic.ttf
var BoldItalic []byte
