package main

import (
	"fmt"
	"game_top/game/entities"
	"game_top/game/entities/player"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Object struct {
	Sprite *ebiten.Image
}

type Game struct {
	Map      [36][240]*Object
	Player   *player.Player
	Entities []*entities.Entity
	Inputs   []ebiten.Key
}

func (g *Game) Update() error {
	g.Inputs = inpututil.AppendPressedKeys(g.Inputs[:0])

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for y := range g.Map {
		for x := range g.Map[y] {
			x1 := float32(x * 16)
			y1 := float32(y * 16)
			vector.StrokeLine(screen, x1, y1, x1+16, y1, 1, color.RGBA{0xFF, 0xFF, 0xFF, 255}, false)
			vector.StrokeLine(screen, float32(x1+16), y1, float32(x1+16), float32(y1+16), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0x00}, false)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(960/2, 540/2)
	if g.Player.Sprite != nil {
		screen.DrawImage(g.Player.Sprite, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %f\nPlayer Y: %f", g.Player.X, g.Player.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	game.Player = &player.Player{
		Entity: entities.Entity{
			Sprite: el_quadrado_vermelho,
		},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
