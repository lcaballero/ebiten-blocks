package shapes

import (
	"image"
	"math"
)

// Vecs represent a simple list of Tuples that can be used to
// represent control points for curves.  Namely these curves: Quad,
// Cubic, or general Control Points and Control Lines, etc.
type Vecs []Vec

func NewRangeVecs(n int, step Vec) Vecs {
	vecs := Vecs{}
	pos := Zero
	for i := 0; i < n; i++ {
		vecs = append(vecs, pos)
		pos = pos.Add(step)
	}
	return vecs
}

// Call executes the function for every Vec in the collection
func (vs Vecs) Call(fn func(int, Vec)) {
	for i := 0; i < vs.Len(); i++ {
		fn(i, vs[i])
	}
}

func (vs Vecs) Points() []image.Point {
	rs := []image.Point{}
	for _, v := range vs {
		rs = append(rs, v.Point())
	}
	return rs
}

// Close repeats the first vec as the last vec closing the slice
func (vs Vecs) Close() Vecs {
	*(&vs) = append(vs, vs[0])
	return vs
}

func (vs Vecs) Add(vecs ...Vec) Vecs {
	*(&vs) = append(vs, vecs...)
	return vs
}

func (vs Vecs) Filter(fn func(v Vec) bool) Vecs {
	vecs := Vecs{}
	for i := 0; i < vs.Len(); i++ {
		v := vs[i]
		if fn(v) {
			vecs = append(vecs, v)
		}
	}
	return vecs
}

// Alter executes the function for every Vec in the collection
func (vs Vecs) Alter(fn func(i int, v Vec) Vec) Vecs {
	vecs := make(Vecs, vs.Len())
	for i := 0; i < vs.Len(); i++ {
		vecs[i] = fn(i, vs[i])
	}
	return vecs
}

// Empty returns true if the number of Vecs is 0 and false
// otherwise.
func (v Vecs) Empty() bool {
	return v.Len() == 0
}

// Vecs2 returns the first two tuples as Vec
func (p Vecs) Vecs2() (v1, v2 Vec) {
	return p[0], p[1]
}

// Vecs3 returns the first three tuples as Vec
func (p Vecs) Vecs3() (v1, v2, v3 Vec) {
	return p[0], p[1], p[2]
}

// Vecs4 returns the first four tuples as Vec
func (p Vecs) Vecs4() (v1, v2, v3, v4 Vec) {
	return p[0], p[1], p[2], p[3]
}

// ToVecs returns the Vecs tuples as Vecs.
func (p Vecs) ToVecs() Vecs {
	n := len(p)
	vecs := make([]Vec, n)
	for i := 0; i < n; i++ {
		vecs[i] = p[0]
	}
	return vecs
}

// Swap exchages the values at index 'i' and 'j'
func (s Vecs) Swap(i, j int) {
	a, b := s[j], s[i]
	s[j] = b
	s[i] = a
}

// Len return the number of vectors in the collection.
func (v Vecs) Len() int {
	return len(v)
}

// Each iterates over the vectors in the collection providing
// the call back function the index and the vector.
func (v Vecs) Each(fn func(i int, v Vec)) {
	for i := 0; i < v.Len(); i++ {
		fn(i, v[i])
	}
}

// Add returns a new Vecs slice with the given vector appended
// to the slice, while the original Vecs is unaltered.
// func (v Vecs) Add(a Vec) Vecs {
// 	return append(v, a)
// }

// C1 provides the x,y coordinate pair from the first tuple in the set.
func (p Vecs) C1() (x0, y0 float64) {
	p0 := p[0]
	return p0.Components()
}

// C2 provides the x,y pairs from first two tuples in the set.
func (p Vecs) C2() (x0, y0, x1, y1 float64) {
	p0, p1 := p[0], p[1]
	x0, y0 = p0.Components()
	x1, y1 = p1.Components()
	return
}

// C3 provides the x,y pairs from first three tuples in the set.
func (p Vecs) C3() (x0, y0, x1, y1, x2, y2 float64) {
	p0, p1, p2 := p[0], p[1], p[2]
	x0, y0 = p0.Components()
	x1, y1 = p1.Components()
	x2, y2 = p2.Components()
	return
}

// C4 provides the x,y pairs from first four tuples in the set.
func (p Vecs) C4() (x0, y0, x1, y1, x2, y2, x3, y3 float64) {
	p0, p1, p2, p3 := p[0], p[1], p[2], p[3]
	x0, y0 = p0.Components()
	x1, y1 = p1.Components()
	x2, y2 = p2.Components()
	x3, y3 = p3.Components()
	return
}

// FindHull finds the rectangle extents of the set.
func (pts Vecs) FindHull() (x, y, w, h float64) {
	mn, mx := -(math.MaxFloat64 - 1.0), math.MaxFloat64
	minX, minY, maxX, maxY := mx, mx, mn, mn

	min, max := math.Min, math.Max

	for i := 0; i < pts.Len(); i++ {
		p := pts[i]
		maxX = max(p.X(), maxX)
		maxY = max(p.Y(), maxY)
		minX = min(p.X(), minX)
		minY = min(p.Y(), minY)
	}

	x, y = minX, minY
	w = maxX - minX
	h = maxY - minY

	return
}

// Line is a parametric function providing a Vec to a point
// on the given Line relative to it's end point.  `t` should
// be a value in the range [0.0,1.0].
func (pts Vecs) Line(t float64) (x, y float64) {
	v1, v2 := pts.Vecs2()
	return Line{A: v1, B: v2}.InLine(t)
}

func (pts Vecs) OnQuad(t float64) Vec {
	return pts.QuadPt(t)
}

// QuadPt is a parametric function providing a Vec to a point
// on the Quad curve relative to it's control points.  `t` should
// be a value in the range [0.0,1.0].
func (pts Vecs) QuadPt(t float64) Vec {
	x0, y0, x1, y1, x2, y2 := pts.C3()

	m := 1 - t
	m2 := m * m
	t2 := t * t

	x := (x0 * m2) + (x1 * 2 * m * t) + (x2 * t2)
	y := (y0 * m2) + (y1 * 2 * m * t) + (y2 * t2)

	return Vec{x, y}
}

// CubicPt is a parametric function providing a Vec to a point
// on the Cubic curve relative to it's control points.  `t` should
// be a value in the range [0.0,1.0].
func (pts Vecs) CubicPt(t float64) Vec {
	x0, y0, x1, y1, x2, y2, x3, y3 := pts.C4()

	pow := math.Pow

	m := 1 - t
	x := (x0 * pow(m, 3)) +
		(x1 * 3 * pow(m, 2) * t) +
		(x2 * 3 * m * t * t) +
		(x3 * t * t * t)

	y := (y0 * pow(m, 3)) +
		(y1 * 3 * pow(m, 2) * t) +
		(y2 * 3 * m * t * t) +
		(y3 * t * t * t)

	return Vec{x, y}
}

// Translate generates a new collection where each point is
// translated by the given values.
func (pts Vecs) Translate(v0 Vec) Vecs {
	n := pts.Len()
	res := make(Vecs, n)
	for i := 0; i < n; i++ {
		v := pts[i]
		res[i] = v.Add(v0)
	}
	return res
}

func (pts Vecs) ScaleXY(sv Vec) Vecs {
	n := pts.Len()
	res := make(Vecs, n)
	for i := 0; i < n; i++ {
		v := pts[i]
		res[i] = v.Scale(sv.X(), sv.Y())
	}
	return res
}

func (pts Vecs) Rotate(rads float64) Vecs {
	n := pts.Len()
	res := make(Vecs, n)
	for i := 0; i < n; i++ {
		v := pts[i]
		res[i] = v.Rotate(rads)
	}
	return res
}

func (pts Vecs) First() Vec {
	if pts.Empty() {
		return Zero
	}
	return pts[0]
}

func (pts Vecs) Last() Vec {
	if pts.Empty() {
		return Zero
	}
	return pts[pts.Len()-1]
}

/*
	Given two parameters t0 and t1 (and with u0 = (1 − t0), u1 = (1 − t1)),
	the part of the curve in the interval [t0, t1] is described by the new control points

Q1 = u0u0u0 P1 + (t0u0u0 + u0t0u0 + u0u0t0) P2 + (t0t0u0 + u0t0t0 + t0u0t0) P3 + t0t0t0 P4

Q2 = u0u0u1 P1 + (t0u0u1 + u0t0u1 + u0u0t1) P2 + (t0t0u1 + u0t0t1 + t0u0t1) P3 + t0t0t1 P4

Q3 = u0u1u1 P1 + (t0u1u1 + u0t1u1 + u0u1t1) P2 + (t0t1u1 + u0t1t1 + t0u1t1) P3 + t0t1t1 P4

Q4 = u1u1u1 P1 + (t1u1u1 + u1t1u1 + u1u1t1) P2 + (t1t1u1 + u1t1t1 + t1u1t1) P3 + t1t1t1 P4
*/
func (pts Vecs) Cut(t0, t1 float64) Vecs {
	u0 := 1 - t0
	u1 := 1 - t1

	p1, p2, p3, p4 := pts[0], pts[1], pts[2], pts[3]

	q1 := p1.ScaleXY(u0 * u0 * u0).
		Add(p2.ScaleXY(t0*u0*u0 + u0*t0*u0 + u0*u0*t0)).
		Add(p3.ScaleXY(t0*t0*u0 + u0*t0*t0 + t0*u0*t0)).
		Add(p4.ScaleXY(t0 * t0 * t0))

	q2 := p1.ScaleXY(u0 * u0 * u1).
		Add(p2.ScaleXY(t0*u0*u1 + u0*t0*u1 + u0*u0*t1)).
		Add(p3.ScaleXY(t0*t0*u1 + u0*t0*t1 + t0*u0*t1)).
		Add(p4.ScaleXY(t0 * t0 * t1))

	q3 := p1.ScaleXY(u0 * u1 * u1).
		Add(p2.ScaleXY(t0*u1*u1 + u0*t1*u1 + u0*u1*t1)).
		Add(p3.ScaleXY(t0*t1*u1 + u0*t1*t1 + t0*u1*t1)).
		Add(p4.ScaleXY(t0 * t1 * t1))

	q4 := p1.ScaleXY(u1 * u1 * u1).
		Add(p2.ScaleXY(t1*u1*u1 + u1*t1*u1 + u1*u1*t1)).
		Add(p3.ScaleXY(t1*t1*u1 + u1*t1*t1 + t1*u1*t1)).
		Add(p4.ScaleXY(t1 * t1 * t1))

	return Vecs{q1, q2, q3, q4}
}

func (pts Vecs) Reverse() Vecs {
	n := pts.Len()
	rv := make(Vecs, n)
	for i, v := range pts {
		rv[n-i-1] = v
	}
	return rv
}
