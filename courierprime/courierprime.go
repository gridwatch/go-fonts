// Package courierprime embeds the Courier Prime typeface (Regular, Bold weights).
package courierprime

import "embed"

// FS provides filesystem access to all embedded Courier Prime TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the Courier Prime Regular (400) TTF data.
//
//go:embed CourierPrime-Regular.ttf
var Regular []byte

// Bold is the Courier Prime Bold (700) TTF data.
//
//go:embed CourierPrime-Bold.ttf
var Bold []byte

// Italic is the Courier Prime Italic (400) TTF data.
//
//go:embed CourierPrime-Italic.ttf
var Italic []byte
