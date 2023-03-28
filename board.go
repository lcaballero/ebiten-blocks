package main

import (
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

func (b *Board) reset() {
	marks := []*mark{}
	for _, m := range b.grid {
		marks = append(marks, m)
	}
	for _, m := range marks {
		delete(b.grid, m.rc)
	}
}

func (b *Board) CanGoRight(t *Tetromino) bool {
	blks := positions[t.tetro][int(t.rot)-1]
	size := t.size
	curr := t.roundPosToSize().Add(shapes.Vec{size})
	for _, blk := range blks {
		pos := curr.Add(blk.Scale(size, size))
		if pos.X() >= b.box.MaxX() {
			return false
		}
		x, y := pos.IntComponents()
		rc := [2]int{x / 10, y / 10}
		_, inGrid := b.grid[rc]
		if inGrid {
			return false
		}
	}
	return true
}

func (b *Board) CanGoLeft(t *Tetromino) bool {
	blks := positions[t.tetro][int(t.rot)-1]
	size := t.size
	curr := t.roundPosToSize().Add(shapes.Vec{-size})
	for _, blk := range blks {
		pos := curr.Add(blk.Scale(size, size))
		if pos.X() < b.box.MinX() {
			return false
		}
		x, y := pos.IntComponents()
		rc := [2]int{x / 10, y / 10}
		_, inGrid := b.grid[rc]
		if inGrid {
			return false
		}
	}
	return true
}

func (b *Board) CanRotate(t *Tetromino) bool {
	shape := positions[t.tetro]
	blks := shape[t.rot.Inc(t.tetro).AsIndex()]
	size := t.size
	curr := t.pos
	for _, p := range blks {
		pos := curr.Add(p.Scale(size, size))
		x := pos.X() + size
		if x > b.box.MaxX() {
			return false
		}
		if x < b.box.MinX() {
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
		marks := b.positions(t)
		b.topOfStack(t, marks)
		b.finalizePosition(t)
		return
	}

	if b.checkCollide(t) {
		t.pos = t.roundPosToSize()
		t.isFrozen = true
		marks := b.positions(t)
		b.topOfStack(t, marks)
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

// topOfStack moves the Tetromino to the top of the stack based on
// it's current position, which due to velocity may have over-shot the
// top of the stack but was halted because it either exited the board
// or collided with pieces below the top of the stack
func (b *Board) topOfStack(t *Tetromino, marks []*mark) {
	hasOverlap := func(ms []*mark) bool {
		in0 := ms[0].in(b.grid)
		in1 := ms[1].in(b.grid)
		in2 := ms[2].in(b.grid)
		in3 := ms[3].in(b.grid)
		return in0 || in1 || in2 || in3
	}
	moveUp := func(ms []*mark) {
		ms[0].up()
		ms[1].up()
		ms[2].up()
		ms[3].up()
	}
	up := 0
	for hasOverlap(marks) {
		moveUp(marks)
		up++
	}
	if up == 0 {
		return
	}
	t.pos = t.pos.Sub(shapes.Vec{0, t.size * float64(up)})
}

func (b *Board) positions(t *Tetromino) []*mark {
	blks := t.blocks()
	size := t.size
	marks := []*mark{}
	for _, bk := range blks {
		p := t.pos.Add(bk.Scale(size, size))
		xm, ym := p.IntComponents()
		x, y := xm/10, ym/10
		rc := [2]int{x, y}
		m := &mark{
			image: t.image,
			pos:   p,
			rc:    rc,
			size:  t.size,
		}
		marks = append(marks, m)
	}
	return marks
}

func (b *Board) finalizePosition(t *Tetromino) {
	marks := b.positions(t)
	for _, m := range marks {
		b.grid[m.rc] = m
	}
}

func (b *Board) Draw(screen *ebiten.Image) {
	for _, m := range b.grid {
		m.Draw(screen)
	}
}

func (b *Board) ClearFullRows(t *Tetromino) []int {
	blks := t.blocks()
	//log.Printf("clearing full rows, type: %s, blks: %v", t.tetro, blks)
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
			//log.Printf("iy: %d, bx: %d, by: %d, bw: %d, bh: %d", iy, bx, by, bw, bh)
			hasRow := true
			for ix := bx; ix < bw; ix++ {
				rc := [2]int{ix, iy}
				_, isInRow := b.grid[rc]
				//log.Printf("ix: %d, rc: %v, isInRow: %v", ix, rc, isInRow)
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
	for _, row := range marks {
		//log.Printf("len row: %d", len(row))
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

func (b *Board) IsGameOver() bool {
	return false
}
