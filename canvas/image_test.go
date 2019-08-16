package canvas

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage_AlphaDefault(t *testing.T) {
	img := &Image{}

	assert.Equal(t, 1.0, img.Alpha())
}

func TestImage_TranslucencyDefault(t *testing.T) {
	img := &Image{}

	assert.Equal(t, 0.0, img.Translucency)
}

func TestNewImageFromFile(t *testing.T) {
	pwd, _ := os.Getwd()
	path := filepath.Join(filepath.Dir(pwd), "theme", "icons", "fyne.png")

	img := NewImageFromFile(path)
	assert.NotNil(t, img)
	assert.Equal(t, img.File, path)
}

func TestNewImageFromURL(t *testing.T) {
	urlStr := "https://fyne.io/img/favicon.png"

	img := NewImageFromURL(urlStr)
	assert.NotNil(t, img)
	assert.NotNil(t, img.Resource)
	assert.Equal(t, img.Resource.Name(), filepath.Base(urlStr))
}
