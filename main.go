package main // import "github.com/lcaballero/ebiten-01"

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
