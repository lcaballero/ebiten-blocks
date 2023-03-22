package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

type mark struct {
	image *ebiten.Image
	pos   shapes.Vec
	rc    [2]int
	size  float64
	down  int
}

func (m *mark) Draw(screen *ebiten.Image) {
	mat := &ebiten.GeoM{}
	mat.Translate(m.pos.Components())
	opts := &ebiten.DrawImageOptions{GeoM: *mat}
	screen.DrawImage(m.image, opts)
}
