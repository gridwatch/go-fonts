// Package barlowcondensed embeds the Barlow Condensed typeface (SemiBold, Bold weights).
package barlowcondensed

import "embed"

// FS provides filesystem access to all embedded Barlow Condensed TTF files.
//
//go:embed *.ttf
var FS embed.FS

// SemiBold is the Barlow Condensed SemiBold (600) TTF data.
//
//go:embed BarlowCondensed-SemiBold.ttf
var SemiBold []byte

// Bold is the Barlow Condensed Bold (700) TTF data.
//
//go:embed BarlowCondensed-Bold.ttf
var Bold []byte
