// Package roboto embeds the Roboto typeface (variable weight+width).
package roboto

import "embed"

// FS provides filesystem access to all embedded Roboto TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Variable is the Roboto variable-weight TTF data (width + weight axes).
//
//go:embed Roboto-VariableFont_wdth_wght.ttf
var Variable []byte
