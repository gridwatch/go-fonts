// Package oswald embeds the Oswald typeface (variable weight).
package oswald

import "embed"

// FS provides filesystem access to all embedded Oswald TTF files.
//
//go:embed *.ttf
var FS embed.FS

// Variable is the Oswald variable-weight TTF data (weight axis 200–700).
//
//go:embed Oswald-VariableFont_wght.ttf
var Variable []byte
