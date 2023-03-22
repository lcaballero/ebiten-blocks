package main

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/lcaballero/ebiten-01/shapes"
	"golang.org/x/image/font/basicfont"
)

type Context interface {
	Text(string, shapes.Vec) Context
	Set(x, y int, clr color.Color) Context
	Image() image.Image
	SetColor(color.Color) Context
	Fill() Context
	DrawRectangle(shapes.Rect) Context
}

type context struct {
	image *ebiten.Image
	ctx   *gg.Context
}

func NewContextFromEbiten(img *ebiten.Image) Context {
	return &context{
		image: img,
		ctx:   gg.NewContextForImage(img),
	}
}

func (c *context) Fill() Context {
	c.ctx.Fill()
	return c
}

func (c *context) SetColor(col color.Color) Context {
	c.ctx.SetColor(col)
	return c
}

func (c *context) Image() image.Image {
	return c.ctx.Image()
}

func (c *context) Text(s string, pos shapes.Vec) Context {
	//log.Printf("s: %s, pos: %v", s, pos)
	x, y := pos.IntComponents()
	text.Draw(c.image, s, basicfont.Face7x13, x, y, color.White)
	return c
}

func (c *context) Set(x, y int, clr color.Color) Context {
	c.image.Set(x, y, clr)
	return c
}

func (c *context) DrawRectangle(r shapes.Rect) Context {
	c.ctx.DrawRectangle(r.Components())
	return c
}
