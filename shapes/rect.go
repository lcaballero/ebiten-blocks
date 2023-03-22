package shapes

import "math"

// Rect is a geometric model which captures an id, the center
// and the dimensions.
type Rect struct {
	Pos Vec
	Dim Vec
}

func NewSqr(sz float64) Rect {
	return NewRect(sz, sz)
}

func NewRectAt(x, y, w, h float64) Rect {
	return Rect{
		Pos: Vec{x, y},
		Dim: Vec{w, h},
	}
}

func NewRect(w, h float64) Rect {
	return NewRectAt(0, 0, w, h)
}

func (r Rect) Components() (x, y, w, h float64) {
	x, y = r.Pos.Components()
	w, h = r.Dim.Components()
	return
}

func (r Rect) DownSlash(top, bot float64) Line {
	return Line{A: r.OnTop(top), B: r.OnBottom(bot)}
}

func (r Rect) CrossSlash(top, bot float64) Line {
	return Line{A: r.OnLeft(top), B: r.OnRight(bot)}
}

func (r Rect) InRect(x, y float64) Vec {
	return r.Dim.Scale(x, y).Add(r.Pos)
}

func (r Rect) OnLeft(f float64) Vec {
	//n0 := Line{A: r.BottomLeft(), B: r.TopLeft()} // left
	return r.LeftSide().OnLine(f)
}

func (r Rect) OnTop(f float64) Vec {
	n1 := Line{A: r.TopLeft(), B: r.TopRight()} // top
	return n1.OnLine(f)
}

func (r Rect) OnRight(f float64) Vec {
	n2 := Line{A: r.TopRight(), B: r.BottomRight()} // right
	return n2.OnLine(f)
}

func (r Rect) OnBottom(f float64) Vec {
	n3 := Line{A: r.BottomRight(), B: r.BottomLeft()} // bottom
	return n3.OnLine(f)
}

func (r Rect) Equals(b Rect) bool {
	return r.Pos.Equals(b.Pos) && r.Dim.Equals(b.Dim)
}

func (r Rect) Normalize() Rect {
	px, py := r.Pos.Components()
	dx, dy := r.Dim.Components()
	// case: quad 2
	if dx < 0.0 && dy < 0.0 {
		return Rect{
			Pos: Vec{px + dx, py + dy},
			Dim: Vec{-dx, -dy},
		}
	}
	// case: quad 1
	if dx < 0.0 {
		return Rect{
			Pos: Vec{px + dx, py},
			Dim: Vec{-dx, dy},
		}
	}
	// case: quad 3
	if dy < 0.0 {
		return Rect{
			Pos: Vec{px, py + dy},
			Dim: Vec{dx, -dy},
		}
	}
	return r
}

func (r Rect) MirrorLeft() Rect {
	return Rect{
		Pos: r.Pos.Sub(r.Dim.VecX()),
		Dim: r.Dim,
	}
}

func (r Rect) MirrorBottom() Rect {
	return Rect{
		Pos: r.Pos.Sub(r.Dim.VecY()),
		Dim: r.Dim,
	}
}

func (r Rect) Expand(d float64) Rect {
	e0 := Vec{-d, -d}
	e1 := Vec{d * 2, d * 2}
	b := Rect{Pos: r.Pos.Add(e0), Dim: r.Dim.Add(e1)}
	return b
}

func (r Rect) Shrink(d float64) Rect {
	e0 := Vec{d, d}
	e1 := e0.ScaleXY(2.0)
	b := Rect{Pos: r.Pos.Add(e0), Dim: r.Dim.Sub(e1)}
	return b
}

func (r Rect) Translate(v Vec) Rect {
	return Rect{Pos: r.Pos.Add(v), Dim: r.Dim}
}

func (r Rect) Scale(x, y float64) Rect {
	return Rect{Pos: r.Pos, Dim: r.Dim.Scale(x, y)}
}

func (r Rect) CenterInRect(a Rect) Rect {
	pos := a.Pos.Sub(r.Pos).Half()
	return Rect{Pos: pos, Dim: r.Dim}
}

// Dims the width and height of the frame
func (r Rect) Dims() Vec {
	return r.Dim
}

// Position of the bottom left corner in Cartesian coordinates
func (r Rect) Position() Vec {
	return r.Pos
}

func (r Rect) XRange() (float64, float64) {
	return r.BottomLeft().X(), r.BottomRight().X()
}

// MinX bottom left
func (r Rect) MinX() float64 {
	return r.BottomLeft().X()
}

// MaxX bottom left
func (r Rect) MaxX() float64 {
	return r.TopRight().X()
}

// MinY bottom left
func (r Rect) MinY() float64 {
	return r.BottomLeft().Y()
}

// MaxY bottom left
func (r Rect) MaxY() float64 {
	return r.TopRight().Y()
}

// X position returned
func (r Rect) X() float64 {
	return r.Pos.X()
}

// Y position returned
func (r Rect) Y() float64 {
	return r.Pos.Y()
}

// W width of the rectangle
func (r Rect) W() float64 {
	return r.Dim.X()
}

// H height of the rectangle
func (r Rect) H() float64 {
	return r.Dim.Y()
}

// TopLeft point relative to the Pos and given Dim of the rect
func (r Rect) TopLeft() Vec {
	return r.Pos.Add(r.Dim.VecY())
}

// TopCenter point relative to the Pos and given Dim of the rect
func (r Rect) TopCenter() Vec {
	return r.Pos.Add(r.Dim.Scale(0.5, 1.0))
}

// TopRight point relative to the Pos and given Dim of the rect
func (r Rect) TopRight() Vec {
	return r.Pos.Add(r.Dim)
}

// MiddleLeft point relative to the Pos and given Dim of the rect
func (r Rect) MiddleLeft() Vec {
	return r.Pos.Add(r.Dim.Scale(0.0, 0.5))
}

// Center is the same as MiddleCenter point relative to the Pos and
// given Dim of the rect
func (r Rect) Center() Vec {
	return r.MiddleCenter()
}

// MiddleCenter point relative to the Pos and given Dim of the rect
func (r Rect) MiddleCenter() Vec {
	return r.Pos.Add(r.Dim.Scale(0.5, 0.5))
}

// MiddleRight point relative to the Pos and given Dim of the rect
func (r Rect) MiddleRight() Vec {
	return r.Pos.Add(r.Dim.Scale(1.0, 0.5))
}

// BottomLeft point relative to the Pos and given Dim of the rect
func (r Rect) BottomLeft() Vec {
	return r.Pos
}

// BottomCenter point relative to the Pos and given Dim of the rect
func (r Rect) BottomCenter() Vec {
	return r.Pos.Add(r.Dim.Scale(0.5, 0.0))
}

// BottomRight point relative to the Pos and given Dim of the rect
func (r Rect) BottomRight() Vec {
	return r.Pos.Add(r.Dim.VecX())
}

func (r Rect) ToBottomRight() Rect {
	return Rect{Pos: r.BottomRight(), Dim: r.Dim}
}

func (r Rect) ToBottomCenter() Rect {
	return Rect{Pos: r.BottomCenter(), Dim: r.Dim}
}

func (r Rect) ToTopLeft() Rect {
	return Rect{Pos: r.TopLeft(), Dim: r.Dim}
}

func (r Rect) ToTopRight() Rect {
	return Rect{Pos: r.TopRight(), Dim: r.Dim}
}

func (r Rect) ToBottomLeft() Rect {
	return Rect{Pos: r.Pos, Dim: r.Dim} // an identity really
}

func (r Rect) Points() []Vec {
	return []Vec{
		r.TopLeft(),
		r.TopCenter(),
		r.TopRight(),
		r.MiddleLeft(),
		r.MiddleCenter(),
		r.MiddleRight(),
		r.BottomLeft(),
		r.BottomCenter(),
		r.BottomRight(),
	}
}

func (r Rect) Mids() []Vec {
	return []Vec{
		r.TopCenter(),
		r.MiddleRight(),
		r.BottomCenter(),
		r.MiddleLeft(),
	}
}

func (r Rect) Divide(divs int) []Rects {
	rects := []Rect{r}
	stages := []Rects{}
	for i := 0; i < divs; i++ {
		var next []Rect
		for _, rr := range rects {
			next = append(next, rr.Split()...)
		}
		stages = append(stages, next)
		rects = next
	}
	return stages
}

func (r Rect) Split() []Rect {
	a := r.BottomLeft()
	b := r.BottomCenter()
	c := r.MiddleLeft()
	d := r.MiddleCenter()
	dim := r.Dim.ScaleXY(0.5)
	return []Rect{
		Rect{Pos: a, Dim: dim},
		Rect{Pos: b, Dim: dim},
		Rect{Pos: c, Dim: dim},
		Rect{Pos: d, Dim: dim},
	}
}

func (b Rect) ContainsWithFuzz(p0 Vec, fuzz Fuzz) bool {
	x0, y0 := b.Pos.Components()
	x1, y1 := b.Pos.Add(b.Dim).Components()
	x2, y2 := p0.Components()
	inX := fuzz.Between(x2, x0, x1)
	inY := fuzz.Between(y2, y0, y1)
	return inX && inY
}

func (b Rect) Contains(p0 Vec) bool {
	x0, y0 := b.Pos.Components()
	x1, y1 := b.Pos.Add(b.Dim).Components()
	x2, y2 := p0.Components()
	inX := x0 <= x2 && x2 <= x1
	inY := y0 <= y2 && y2 <= y1
	return inX && inY
}

func (b Rect) Clamp(v Vec) Vec {
	if b.Contains(v) {
		return v
	}
	x0, y0 := v.Components()
	x := math.Max(x0, b.MinX())
	x = math.Min(x, b.MaxX()-1.0)
	y := math.Max(y0, b.MinY())
	y = math.Min(y, b.MaxY()-1.0)
	return Vec{x, y}
}

// in returns true if the components of p0 are between the
// components of p1 and p2 respectively.
func (b *Rect) in(p0, p1, p2 Vec) bool {
	x0, y0 := p0.Components()
	x1, y1 := p1.Components()
	x2, y2 := p2.Components()
	inX := x1 <= x0 && x0 <= x2 // x0 in [x1,x2]?
	inY := y1 <= y0 && y0 <= y2 // y0 in [y1,y2]?
	return inX && inY
}

// within only checks that 'a' overlaps 'b', but not if 'b' overlaps
// 'a'.  The distinction being if 'a' is inside 'b' then the corners
// of 'b' are not inside of 'a' which is how this algorithm detects
// if the boxes overlap.  See 'In' which tests both if 'a' overlaps
// 'b' and if 'b' overlaps 'a' to cover this case.
func (a Rect) within(b Rect, margin float64) bool {
	m := Vec{margin, margin}

	b1 := b.Pos.Sub(b.Dim.ScaleXY(.5)).Sub(m)
	b2 := b1.Add(b.Dim.Add(m.ScaleXY(2.0)))

	a1 := a.Pos.Sub(a.Dim.ScaleXY(.5))
	a2 := a1.Add(a.Dim)
	a3 := a2.Sub(Vec{a.Dim.X(), 0.0})
	a4 := a2.Sub(Vec{0.0, a.Dim.Y()})

	in := a.in

	return in(a1, b1, b2) ||
		in(a2, b1, b2) ||
		in(a3, b1, b2) ||
		in(a4, b1, b2)
}

// In determines if this box overlaps the other box with additional
// margin added to normal geometric model.
func (a Rect) In(b Rect, margin float64) bool {
	return a.within(b, margin) || b.within(a, margin)
}

func (a Rect) Overlaps(b Rect) bool {
	inA := a.Contains(b.TopLeft()) ||
		a.Contains(b.TopRight()) ||
		a.Contains(b.BottomLeft()) ||
		a.Contains(b.BottomRight())
	inB := b.Contains(a.TopLeft()) ||
		b.Contains(a.TopRight()) ||
		b.Contains(a.BottomLeft()) ||
		b.Contains(a.BottomRight())
	return inA || inB
}

// Rect provides the normal geometric values for rendering the box
// as a rectangle with the corner as x,y and w,h.
func (b *Rect) Rect() (x, y, w, h float64) {
	x, y = b.Pos.Sub(b.Dim.ScaleXY(.5)).Components()
	w, h = b.Dim.Components()
	return
}

// Lines returns each of the line segments that make up the edges of
// the boxes.
func (b *Rect) Lines() (s1, s2, s3, s4 Line) {
	x, y, w, h := b.Rect()
	v := Vec{x, y}

	s1 = Line{A: v, B: v.Translate(w, 0.0)}
	s2 = Line{A: s1.B, B: s1.B.Translate(0.0, h)}
	s3 = Line{A: s2.B, B: s2.B.Translate(-w, 0.0)}
	s4 = Line{A: s3.B, B: s3.B.Translate(0.0, -h)}
	return s1, s2, s3, s4
}

func (r Rect) Tiles(divs int) []Rect {
	size := r.Dim.ScaleXY(1 / float64(divs))
	maxW, maxH := r.Dim.Components()
	w0, h0 := size.Components()
	result := []Rect{}
	for w := 0.0; w < maxW; w += w0 {
		for h := 0.0; h < maxH; h += h0 {
			c := r.Pos.Translate(w, h)
			result = append(result, Rect{Pos: c, Dim: size})
		}
	}
	return result
}

func (r Rect) MidPts() Vecs {
	return Vecs{
		r.MiddleLeft(),
		r.TopCenter(),
		r.MiddleRight(),
		r.BottomCenter(),
	}
}

func (r Rect) Corners() Vecs {
	return Vecs{
		r.BottomLeft(),
		r.TopLeft(),
		r.TopRight(),
		r.BottomRight(),
	}
}

func (r Rect) Sides() []Line {
	n0 := Line{A: r.BottomLeft(), B: r.TopLeft()} // left
	n1 := Line{A: n0.B, B: r.TopRight()}          // top
	n2 := Line{A: n1.B, B: r.BottomRight()}       // right
	n3 := Line{A: n2.B, B: n0.A}                  // bottom
	return []Line{n0, n1, n2, n3}
}

func (r Rect) LeftSide() Line {
	return Line{A: r.BottomLeft(), B: r.TopLeft()}
}

func (r Rect) TopSide() Line {
	return Line{A: r.TopLeft(), B: r.TopRight()}
}

func (r Rect) RightSide() Line {
	return Line{A: r.TopRight(), B: r.BottomRight()}
}

func (r Rect) BottomSide() Line {
	return Line{A: r.BottomRight(), B: r.BottomLeft()}
}

func (r Rect) AddRight(amt float64) Rect {
	return Rect{
		Pos: r.Pos,
		Dim: Vec{
			r.Dim.X() + amt,
			r.Dim.Y(),
		},
	}
}

func (r Rect) AddTop(amt float64) Rect {
	return Rect{
		Pos: r.Pos,
		Dim: Vec{
			r.Dim.X(),
			r.Dim.Y() + amt,
		},
	}
}

func (r Rect) AddBottom(amt float64) Rect {
	return Rect{
		Pos: Vec{r.Pos.X(), r.Pos.Y() - amt},
		Dim: Vec{r.Dim.X(), r.Dim.Y() + amt},
	}
}

func (r Rect) AddLeft(amt float64) Rect {
	return Rect{
		Pos: Vec{r.Pos.X() - amt, r.Pos.Y()},
		Dim: Vec{r.Dim.X() + amt, r.Dim.Y()},
	}
}
