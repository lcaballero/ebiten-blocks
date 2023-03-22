package main

import (
	"github.com/lcaballero/ebiten-01/rand"
	"github.com/lcaballero/ebiten-01/shapes"
)

type Tetro int

const (
	I Tetro = 1
	O Tetro = 2
	T Tetro = 3
	S Tetro = 4
	Z Tetro = 5
	J Tetro = 6
	L Tetro = 7
)

func (t Tetro) String() string {
	switch t {
	case I:
		return "I"
	case O:
		return "O"
	case T:
		return "T"
	case S:
		return "S"
	case Z:
		return "Z"
	case J:
		return "J"
	case L:
		return "L"
	default:
		return "unkown"
	}
}

func RandTetro(rnd rand.Rnd) Tetro {
	maxTetro := int(L)
	n := rnd.Int(maxTetro)
	return Tetro(n + 1)
}

type Rotation int

const (
	R1 Rotation = 1
	R2 Rotation = 2
	R3 Rotation = 3
	R4 Rotation = 4
)

func (r Rotation) Inc(t Tetro) Rotation {
	length := len(positions[t])
	curr := int(r) - 1
	next := (curr + 1) % length
	return Rotation(next + 1)
}

/*
I:
   * ----
   * ----
   * ----
   * ****

O:
   --
   **
   **

T:
   --- -* --- *-
   *** ** -*- **
   x*- x* *** *-

S:
   --- *-
   -** **
   **- x*

Z:
   --- -*
   **- **
   x** *-

J:
   -* --- ** ---
   -* *-- *- ***
   ** *** *- x-*

L:
   *- --- ** ---
   *- *** -* --*
   ** *-- x* ***

*/

var positions = map[Tetro][]shapes.Vecs{
	I: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{0, -2},
			shapes.Vec{0, -3},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{2, 0},
			shapes.Vec{3, 0},
		},
	},
	O: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{1, 0},
			shapes.Vec{1, -1},
		},
	},
	T: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, -1},
			shapes.Vec{1, 0},
			shapes.Vec{1, -1},
			shapes.Vec{2, -1},
		},
		shapes.Vecs{
			shapes.Vec{1, 0},
			shapes.Vec{0, -1},
			shapes.Vec{1, -1},
			shapes.Vec{1, -2},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{2, 0},
			shapes.Vec{1, -1},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{0, -2},
			shapes.Vec{1, -1},
		},
	},
	S: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{1, -1},
			shapes.Vec{2, -1},
		},
		shapes.Vecs{
			shapes.Vec{1, 0},
			shapes.Vec{0, -1},
			shapes.Vec{1, -1},
			shapes.Vec{0, -2},
		},
	},
	Z: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, -1},
			shapes.Vec{1, 0},
			shapes.Vec{2, 0},
			shapes.Vec{1, -1},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{1, -1},
			shapes.Vec{1, -2},
		},
	},
	J: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{1, -1},
			shapes.Vec{1, -2},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{2, 0},
			shapes.Vec{0, -1},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{0, -2},
			shapes.Vec{1, -2},
		},
		shapes.Vecs{
			shapes.Vec{0, -1},
			shapes.Vec{1, -1},
			shapes.Vec{2, -1},
			shapes.Vec{2, 0},
		},
	},
	L: []shapes.Vecs{
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{0, -2},
			shapes.Vec{1, 0},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{0, -1},
			shapes.Vec{1, -1},
			shapes.Vec{2, -1},
		},
		shapes.Vecs{
			shapes.Vec{1, 0},
			shapes.Vec{1, -1},
			shapes.Vec{1, -2},
			shapes.Vec{0, -2},
		},
		shapes.Vecs{
			shapes.Vec{0, 0},
			shapes.Vec{1, 0},
			shapes.Vec{2, 0},
			shapes.Vec{2, -1},
		},
	},
}
