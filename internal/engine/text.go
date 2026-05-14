package engine

import (
	"bytes"
	"image/color"
	"project_m/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func LoadFont(path string) *text.GoTextFaceSource {
	dadosFonte, err := assets.Assets.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fonte, err := text.NewGoTextFaceSource(bytes.NewReader(dadosFonte))
	if err != nil {
		panic(err)
	}

	return fonte
}

func DrawText(screen *ebiten.Image, fonte *text.GoTextFaceSource, texto string, x, y float64, size float64, color color.Color) {
	face := &text.GoTextFace{
		Source: fonte,
		Size:   size,
	}

	op := &text.DrawOptions{}

	op.GeoM.Translate(x,y)
	op.ColorScale.ScaleWithColor(color)

	text.Draw(screen, texto, face, op)
}