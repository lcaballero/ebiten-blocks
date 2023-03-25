package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

type Tetromino struct {
	image    *ebiten.Image
	pos      shapes.Vec
	size     float64
	tetro    Tetro
	rot      Rotation
	velocity shapes.Vec
	isFrozen bool
}

func (t *Tetromino) Update(elapsed time.Duration, dt float64) {
	if t.isFrozen {
		return
	}
	t.pos = t.pos.Add(t.velocity.Scale(dt, dt))
}

func (t *Tetromino) Draw(screen *ebiten.Image) {
	blk := t.blocks()
	size := t.size
	for _, p := range blk {
		pos := t.pos.Add(p.Scale(size, size))
		m := &ebiten.GeoM{}
		m.Translate(pos.Components())
		opts := &ebiten.DrawImageOptions{GeoM: *m}
		screen.DrawImage(t.image, opts)
	}
}

func (t *Tetromino) MoveRight() {
	if t.isFrozen {
		return
	}
	t.pos = t.pos.Add(shapes.Vec{t.size, 0})
}

func (t *Tetromino) MoveLeft() {
	if t.isFrozen {
		return
	}
	t.pos = t.pos.Add(shapes.Vec{-t.size, 0})
}

func (t *Tetromino) RotateRight() {
	if t.isFrozen {
		return
	}
	t.rot = t.rot.Inc(t.tetro)
}

func (t *Tetromino) roundPosToSize() shapes.Vec {
	return shapes.Vec{
		math.Floor(t.pos.X()/t.size) * t.size,
		math.Floor(t.pos.Y()/t.size) * t.size,
	}
}

func (t *Tetromino) blocks() shapes.Vecs {
	blks := positions[t.tetro]
	blk := blks[int(t.rot)-1]
	return blk
}

func (t *Tetromino) Max() shapes.Vec {
	var v shapes.Vec
	switch t.tetro {
	case I:
		v = shapes.Vec{1, 4}
	case O:
		v = shapes.Vec{2, 2}
	case T:
		v = shapes.Vec{3, 2}
	case S:
		v = shapes.Vec{3, 2}
	case Z:
		v = shapes.Vec{3, 2}
	case J:
		v = shapes.Vec{2, 3}
	case L:
		v = shapes.Vec{2, 3}
	default:
		panic("not a tetrimino")
	}
	return t.pos.Add(v.Scale(t.size, t.size))
}

func (t *Tetromino) Center() shapes.Vec {
	var v shapes.Vec
	switch t.tetro {
	case I:
		v = shapes.Vec{1, 4}
	case O:
		v = shapes.Vec{2, 2}
	case T:
		v = shapes.Vec{3, 2}
	case S:
		v = shapes.Vec{3, 2}
	case Z:
		v = shapes.Vec{3, 2}
	case J:
		v = shapes.Vec{2, 3}
	case L:
		v = shapes.Vec{2, 3}
	default:
		panic("not a tetrimino")
	}
	return v.Scale(t.size, t.size).Half()
}

func (t *Tetromino) Accelerate() {
	t.velocity = shapes.Vec{0, 600}
}

func (t *Tetromino) MoveCenterTo(c shapes.Vec) *Tetromino {
	center := t.Center()
	t.pos = c.Sub(center.Scale(1, -1)).Sub(shapes.Vec{0, 10})
	return t
}
