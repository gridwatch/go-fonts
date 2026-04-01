package roboto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func TestVariable(t *testing.T) {
	require.Greater(t, len(Variable), 1000)
	assertParseable(t, Variable, "Variable")
}

func TestFS(t *testing.T) {
	entries, err := FS.ReadDir(".")
	require.NoError(t, err)
	assert.Len(t, entries, 1)

	data, err := FS.ReadFile(entries[0].Name())
	require.NoError(t, err)
	assertParseable(t, data, entries[0].Name())
}

func assertParseable(t *testing.T, data []byte, name string) {
	t.Helper()
	f, err := opentype.Parse(data)
	require.NoError(t, err, "font %s should parse as valid TTF", name)

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size: 16, DPI: 72, Hinting: font.HintingFull,
	})
	require.NoError(t, err, "font %s should produce a usable face", name)
	assert.NotNil(t, face)
}
