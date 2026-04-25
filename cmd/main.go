package main

import (
	"log"
	"project_m/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.GameSetup()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
