// Package ibmplexmono embeds the IBM Plex Mono typeface (Regular, Bold weights).
package ibmplexmono

import "embed"

// FS provides filesystem access to all embedded IBM Plex Mono TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Regular is the IBM Plex Mono Regular (400) TTF data.
//
//go:embed IBMPlexMono-Regular.ttf
var Regular []byte

// Bold is the IBM Plex Mono Bold (700) TTF data.
//
//go:embed IBMPlexMono-Bold.ttf
var Bold []byte

// Italic is the IBM Plex Mono Italic (400) TTF data.
//
//go:embed IBMPlexMono-Italic.ttf
var Italic []byte
