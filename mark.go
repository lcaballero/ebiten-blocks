package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

type mark struct {
	image *ebiten.Image
	pos   shapes.Vec
	size  float64
	rc    [2]int
	down  int
}

func (m *mark) in(grid grid) bool {
	_, in := grid[m.rc]
	return in
}

func (m *mark) up() *mark {
	m.rc = [2]int{m.rc[0], m.rc[1] - 1}
	return m
}

func (m *mark) Draw(screen *ebiten.Image) {
	mat := &ebiten.GeoM{}
	mat.Translate(m.pos.Components())
	opts := &ebiten.DrawImageOptions{GeoM: *mat}
	screen.DrawImage(m.image, opts)
}
