package shapes

import "math"

// A Line described by two Vectors.
type Line struct {
	A, B Vec
}

func VertLine(ht float64) Line {
	return Line{A: Zero, B: Vec{0, ht}}
}

// NewLine composed of the two endpoints x0,y0 and x1,y1.
func NewLine(x0, y0, x1, y1 float64) Line {
	return Line{
		A: Vec{x0, y0},
		B: Vec{x1, y1},
	}
}

// NewLineFromeVecs creates a line based on the points provided by the
// vectors 'a' and 'b'.  No checking is done to determine if the
// vectors provided makes sense as line endpoints.
func NewLineFromVecs(a, b Vec) Line {
	return NewLine(a.X(), a.Y(), b.X(), b.Y())
}

// Angle provides the angle between the endpoints of the line
func (e Line) Angle() float64 {
	return e.ToVec().Angle()
}

// Reverse exchanges endpoints so that the line start and end are
// swapped.
func (e Line) Reverse() Line {
	return NewLineFromVecs(e.B, e.A)
}

func (e Line) Slope() float64 {
	v := e.A.Sub(e.B)
	return v.X() / v.Y()
}

// V1 returns the vector where the Line begins.
func (e Line) V1() Vec {
	return e.A
}

// V2 returns the vector where the Line ends.
func (e Line) V2() Vec {
	return e.B
}

// ToVec returns the vector which represents the direction and
// magnitude between the start and end vectors describing the Line.
func (e Line) ToVec() Vec {
	return e.B.Sub(e.A)
}

// Len returns the length of the Line.
func (e Line) Len() float64 {
	return e.ToVec().Mag()
}

func (e Line) MidPt() Vec {
	return e.Mid().Add(e.A)
}

func (e Line) Rotate(t float64) Line {
	v0 := e.ToVec()
	v1 := v0.Rotate(t)
	line := Line{A: e.A, B: v1.Add(e.A)}
	return line
}

// Mid returns the mid point of the line as a Vec
func (e Line) Mid() Vec {
	v := e.ToVec()
	a := v.ScaleXY(0.5)
	return a
}

// Half produces two smaller lines which are co-linear with the
// basis Line but half as long and share an endpoint at the midpoint
// of the basis Line.
func (e Line) Half() (Line, Line) {
	return e.Split(0.5)
}

// Split divides the basis Line into two separate lines where the
// ratio between the two resulting lines is determined by the factor
// parameter.  Factor should be between 0.0 and the 1.0 exclusive.
func (e Line) Split(factor float64) (Line, Line) {
	v := e.ToVec()
	a := v.ScaleXY(factor)
	b := a.Translate(e.A.Components())

	return Line{A: e.A, B: b}, Line{A: b, B: e.B}
}

// IsVertical returns true if the Line is vertical by comparing
// the X value for the A and B points.
func (e Line) IsVertical() bool {
	return e.A.X() == e.B.X()
}

// IsHorizontal returns true if the Line horizontal by comparing
// the Y value for the A and B points.
func (e Line) IsHorizontal() bool {
	return e.A.Y() == e.B.Y()
}

// ToPts returns the structure as a Vecs collection.
func (e Line) ToPts() Vecs {
	return Vecs{e.A, e.B}
}

// Components provides the ends of the Line (x0,y0) (x1,y1)
func (e Line) Components() (x0, y0, x1, y1 float64) {
	x0, y0 = e.A.Components()
	x1, y1 = e.B.Components()
	return
}

// Translate moves the endpoints of the lines by the values provided
func (e Line) Translate(v Vec) Line {
	x, y := v.X(), v.Y()
	return Line{
		A: e.A.Translate(x, y),
		B: e.B.Translate(x, y),
	}
}

// IntersectsAt return the Pt where the two lines intersect and a
// true boolean value indicating the intersection is valid. Otherwise,
// a zero value Pt and false is returned indicating that the lines
// are parallel or overlap and don't intersect at all.
func (e Line) IntersectsAt(b Line) (Vec, bool) {
	p1, p2 := e.A, e.B
	p3, p4 := b.A, b.B
	return IntersectLinePts(p1, p2, p3, p4)
}

// InLine is a parametric function providing x,y pair to a point
// on the given Line relative to it's end points. `t` should be a
// value in the range [0.0, 1.0].
func (e Line) InLine(t float64) (x, y float64) {
	x0, y0 := e.A.Components()
	x1, y1 := e.B.Components()
	x = ((1 - t) * x0) + (x1 * t)
	y = ((1 - t) * y0) + (y1 * t)
	return x, y
}

func (e Line) OnLine(t float64) Vec {
	x0, y0 := e.InLine(t)
	return Vec{x0, y0}
}

func (e Line) ScaleFromMidPt(t float64) Line {
	v0 := e.ToVec()
	v1 := v0.ScaleXY(t)
	dt := v1.Sub(v0)
	line := Line{A: e.A.Sub(dt), B: e.B.Add(dt)}
	return line
}

func (e Line) ScaleFromA(t float64) Line {
	v := e.ToVec().ScaleXY(t)
	line := Line{B: v}
	return line.Translate(e.A)
}

func (e Line) ScaleFromB(t float64) Line {
	v := e.ToVec().ScaleXY(t).Rotate(math.Pi)
	line := Line{B: v}
	return line.Translate(e.A)
}

func IntersectLineCoords(x1, y1, x2, y2, x3, y3, x4, y4 float64) (Vec, bool) {
	nx := (x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)
	ny := (x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)
	d := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if d == 0 {
		return Vec{}, false
	}
	return Vec{nx / d, ny / d}, true
}

// IntersectLinePts determines if the lines made from p1,p2 and p3,p4
// intersect.  If they are parallel a zero pt is returned along with a
// false, else the Pt of intersection is returned with true.
func IntersectLinePts(p1, p2, p3, p4 Vec) (Vec, bool) {
	x1, y1 := p1.Components()
	x2, y2 := p2.Components()
	x3, y3 := p3.Components()
	x4, y4 := p4.Components()
	return IntersectLineCoords(x1, y1, x2, y2, x3, y3, x4, y4)
}
