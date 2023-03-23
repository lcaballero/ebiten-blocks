package main // import "github.com/lcaballero/ebiten-01"

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:generate go run ./scripts/cli/main.go

func main() {
	procs := Procs{
		NewGame: StartGame,
	}
	err := NewApp(procs).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func StartGame(vals Vals) error {
	game := NewGame(NewGameOpts{vals})
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tetris")
	err := ebiten.RunGame(game)
	return err
}
