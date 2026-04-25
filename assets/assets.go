package assets

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)
// NÃO TIRE ISTO, NÃO É UM COMENTÁRIO, É SERIO! É IMPORTANTE ISTO AQUI EM BAIXO 
//go:embed *
var Assets embed.FS

func LoadImage(path string) *ebiten.Image {
	f, err := Assets.ReadFile(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
