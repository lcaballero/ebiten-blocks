package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

func showAllTetronimos(screen *ebiten.Image, blocks Cells) {
	s := shapes.Vec{20, 40}
	ix := 0

	blk := blocks.At(ix + 0)
	t := &Tetromino{image: blk, pos: s, tetro: S, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 1)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{30}), tetro: S, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 2)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 30}), tetro: T, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 3)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{30, 30}), tetro: T, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 4)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{50, 30}), tetro: T, rot: R3}
	t.Draw(screen)

	blk = blocks.At(ix + 5)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{80, 40}), tetro: T, rot: R4}
	t.Draw(screen)

	blk = blocks.At(ix + 6)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 60}), tetro: S, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 7)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{30, 60}), tetro: S, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 8)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 90}), tetro: Z, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 9)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{30, 90}), tetro: Z, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 10)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 120}), tetro: J, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 11)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{20, 120}), tetro: J, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 12)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{50, 120}), tetro: J, rot: R3}
	t.Draw(screen)

	blk = blocks.At(ix + 13)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{70, 120}), tetro: J, rot: R4}
	t.Draw(screen)

	blk = blocks.At(ix + 14)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 150}), tetro: L, rot: R1}
	t.Draw(screen)

	blk = blocks.At(ix + 15)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{20, 150}), tetro: L, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 16)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{50, 150}), tetro: L, rot: R3}
	t.Draw(screen)

	blk = blocks.At(ix + 17)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{70, 150}), tetro: L, rot: R4}
	t.Draw(screen)

	blk = blocks.At(ix + 18)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{0, 170}), tetro: I, rot: R2}
	t.Draw(screen)

	blk = blocks.At(ix + 19)
	t = &Tetromino{image: blk, pos: s.Add(shapes.Vec{50, 170}), tetro: I, rot: R1}
	t.Draw(screen)

}
