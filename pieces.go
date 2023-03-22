package main

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/rand"
)

type Cells []*ebiten.Image

func (c Cells) Pick(rnd rand.Rnd) *ebiten.Image {
	n := rnd.Int(len(c))
	return c[n]
}

func (c Cells) At(n int) *ebiten.Image {
	ix := n % len(c)
	return c[ix]
}

type Pieces struct {
	image  *ebiten.Image
	blocks Cells
}

func NewPieces() *Pieces {
	f, err := os.Open("./blocks.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	all := ebiten.NewImageFromImage(img)
	blocks := make([]*ebiten.Image, 7)
	for i := 0; i < 7; i++ {
		min := image.Point{X: i * 10}
		max := image.Point{X: min.X + 10, Y: min.Y + 10}
		s := image.Rectangle{Min: min, Max: max}
		sub := all.SubImage(s)
		block := sub.(*ebiten.Image)
		blocks[i] = block
	}
	return &Pieces{
		image:  all,
		blocks: blocks,
	}
}
