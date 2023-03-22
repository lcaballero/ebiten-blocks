package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/shapes"
)

type Background struct {
	ctx     Context
	w, h    int
	scoring ScoreBoard
	canvas  shapes.Rect
	board   shapes.Rect
	next    shapes.Rect
	score   shapes.Rect
	level   shapes.Rect
	lines   shapes.Rect
}

func NewBackground() *Background {
	return &Background{
		scoring: ScoreBoard{Score: 0, Lines: 0, Level: 1},
		canvas:  shapes.NewRectAt(0, 0, 640, 480),
		board:   shapes.NewRectAt(20, 20, 100, 200),
		next:    shapes.NewRectAt(170, 20, 60, 60),
		score:   shapes.NewRectAt(170, 110, 120, 20),
		level:   shapes.NewRectAt(170, 150, 120, 20),
		lines:   shapes.NewRectAt(170, 190, 120, 20),
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	if b.ctx == nil {
		bounds := screen.Bounds()
		b.w, b.h = bounds.Max.X, bounds.Max.Y
		b.ctx = NewContextFromEbiten(screen)
		b.bg(b.ctx)
	}
	b.copy(screen)
	b.labels(b.ctx)
}

func (b *Background) copy(screen *ebiten.Image) {
	img := b.ctx.Image()
	for x := 0; x < b.w; x++ {
		for y := 0; y < b.h; y++ {
			screen.Set(x, y, img.At(x, y))
		}
	}
}

func (b *Background) labels(ctx Context) {
	ls := shapes.Vec{0, -2}
	ctx.Text("Next", b.next.Pos.Add(ls))
	ctx.Text("Score", b.score.Pos.Add(ls))
	ctx.Text("Level", b.level.Pos.Add(ls))
	ctx.Text("Lines", b.lines.Pos.Add(ls))

	ls = shapes.Vec{5, 15}
	score := fmt.Sprintf("%d", b.scoring.Score)
	level := fmt.Sprintf("%d", b.scoring.Level)
	lines := fmt.Sprintf("%d", b.scoring.Lines)
	ctx.Text(score, b.score.Pos.Add(ls))
	ctx.Text(level, b.level.Pos.Add(ls))
	ctx.Text(lines, b.lines.Pos.Add(ls))
}

func (b *Background) bg(ctx Context) {
	red := color.RGBA{R: uint8(255), A: uint8(255)}

	ctx.SetColor(red)
	ctx.DrawRectangle(b.canvas)
	ctx.Fill()

	ctx.SetColor(color.Black)
	ctx.DrawRectangle(b.board)
	ctx.Fill()

	ctx.SetColor(color.Black)
	ctx.DrawRectangle(b.next)
	ctx.Fill()

	ctx.SetColor(color.Black)
	ctx.DrawRectangle(b.score)
	ctx.Fill()

	ctx.SetColor(color.Black)
	ctx.DrawRectangle(b.level)
	ctx.Fill()

	ctx.SetColor(color.Black)
	ctx.DrawRectangle(b.lines)
	ctx.Fill()
}
