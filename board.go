package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

type grid map[[2]int]*mark

type Board struct {
	box  shapes.Rect
	grid grid
}

func NewBoard(box shapes.Rect) *Board {
	return &Board{
		box:  box,
		grid: grid{},
	}
}

func (b *Board) Has(x, y int) bool {
	return false
}

func (b *Board) In(x, y int) bool {
	return false
}

func (b *Board) Fill(t *Tetromino) {
}

func (b *Board) CanGoRight(t *Tetromino) bool {
	blks := positions[t.tetro]
	blk := blks[int(t.rot)-1]
	size := t.size
	curr := t.pos.Add(shapes.Vec{size})
	for _, p := range blk {
		pos := curr.Add(p.Scale(size, size))
		if pos.X() >= b.box.MaxX() {
			return false
		}
	}
	return true
}

func (b *Board) CanGoLeft(t *Tetromino) bool {
	blk := t.blocks()
	size := t.size
	curr := t.pos.Add(shapes.Vec{-size})
	for _, p := range blk {
		pos := curr.Add(p.Scale(size, size))
		if pos.X() < b.box.MinX() {
			return false
		}
	}
	return true
}

func (b *Board) CheckBounds(t *Tetromino) {
	size := t.size
	y := t.pos.Y() / size
	maxY := (b.box.MaxY() - size) / size

	// case: require new active peice
	if t.isFrozen {
		return
	}

	// case: below the bottom of the box
	if y > maxY {
		t.pos = shapes.Vec{t.pos.X(), float64(maxY * size)}
		t.isFrozen = true
		b.finalizePosition(t)
		return
	}

	if b.checkCollide(t) {
		t.pos = t.roundPosToSize()
		t.isFrozen = true
		b.finalizePosition(t)
		return
	}
}

func (b *Board) checkCollide(t *Tetromino) bool {
	blks := t.blocks()
	size := t.size
	for _, bk := range blks {
		p := t.pos.Add(bk.Scale(size, size)).Add(shapes.Vec{0, size})
		xm, ym := p.IntComponents()
		x, y := xm/10, ym/10
		rc := [2]int{x, y}
		_, isFinal := b.grid[rc]
		if isFinal {
			return true
		}
	}
	return false
}

func (b *Board) finalizePosition(t *Tetromino) {
	log.Printf("finalizing piece")
	blks := t.blocks()
	size := t.size
	for _, bk := range blks {
		p := t.pos.Add(bk.Scale(size, size))
		xm, ym := p.IntComponents()
		x, y := xm/10, ym/10
		rc := [2]int{x, y}
		b.grid[rc] = &mark{
			image: t.image,
			pos:   p,
			rc:    rc,
			size:  t.size,
		}
	}
}

func (b *Board) Draw(screen *ebiten.Image) {
	for _, m := range b.grid {
		m.Draw(screen)
	}
}

func (b *Board) ClearFullRows(t *Tetromino) []int {
	blks := t.blocks()
	log.Printf("clearing full rows, type: %s, blks: %v", t.tetro, blks)
	size := t.size
	marks := [][]*mark{}
	rows := []int{}
	set := map[int]bool{}
	for _, bk := range blks {
		p := t.pos.Add(bk.Scale(size, size))
		py := int(p.Y() / 10)
		x, y, w, h := b.box.Components()
		bx, by, bw, bh := int(x/size), int(y/size), int(w/size), int(h/size)
		bw, bh = bw+bx, bh+by
		for iy := by; iy < bh; iy++ {
			if iy != py || set[iy] {
				continue
			}
			set[iy] = true
			log.Printf("iy: %d, bx: %d, by: %d, bw: %d, bh: %d", iy, bx, by, bw, bh)
			hasRow := true
			for ix := bx; ix < bw; ix++ {
				rc := [2]int{ix, iy}
				_, isInRow := b.grid[rc]
				log.Printf("ix: %d, rc: %v, isInRow: %v", ix, rc, isInRow)
				hasRow = hasRow && isInRow
			}
			if !hasRow {
				continue
			}
			rows = append(rows, iy)
			row := []*mark{}
			for ix := bx; ix < bw; ix++ {
				rc := [2]int{ix, iy}
				mark, _ := b.grid[rc]
				row = append(row, mark)
			}
			marks = append(marks, row)
		}
	}
	if len(marks) > 0 {
		log.Printf("removing row: %v", marks[0][0])
	}
	for _, row := range marks {
		log.Printf("len row: %d", len(row))
		for _, mark := range row {
			delete(b.grid, mark.rc)
		}
	}
	up := []*mark{}
	for _, row := range marks {
		for _, mk := range row {
			x, iy := mk.rc[0], mk.rc[1]
			for y := iy - 1; y >= 0; y-- {
				rc := [2]int{x, y}
				spot, isInGrid := b.grid[rc]
				if isInGrid {
					spot.down++
					up = append(up, spot)
				}
			}
		}
	}
	for _, m := range up {
		delete(b.grid, m.rc)
		h := float64(m.down) * m.size
		m.pos = m.pos.Add(shapes.Vec{0, h})
		m.rc = [2]int{m.rc[0], m.rc[1] + m.down}
		m.down = 0
		b.grid[m.rc] = m
	}
	return rows
}
