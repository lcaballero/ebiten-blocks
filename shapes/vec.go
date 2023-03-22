package shapes

import (
	"errors"
	"fmt"
	"image"
	"math"
)

// ErrFindingRadianAngleBetweenVectors where vectors might include a
// zero vector
var ErrFindingRadianAngleBetweenVectors = errors.New(
	`error finding radians between vectors
	where one or both vectors is a zero vector.`,
)

// Vec is a vector of float64 with dimensionality of 2.
type Vec [2]float64

// Zero is a Vec where x, y are 0.0.
var Zero = Vec{0.0, 0.0}
var UnitX = Vec{1.0, 0.0}
var UnitY = Vec{0.0, 1.0}
var Ident = Vec{1.0, 1.0}

func (v Vec) IsNaN() bool {
	return math.IsNaN(v.X()) || math.IsNaN(v.Y())
}

func (v Vec) Components() (x, y float64) {
	return v.X(), v.Y()
}

func (v Vec) Size() (x, y int) {
	return int(v.X()), int(v.Y())
}

// X provides the x component of the vector.
func (v Vec) X() float64 {
	return v[0]
}

// Y provides the y component of the vector.
func (v Vec) Y() float64 {
	return v[1]
}

// Dim provides the dimensionality of the Vector which should always be 2.
func (v Vec) Dim() int {
	return len(v)
}

// Equals compares this vector against the parameter Vec component by
// component, and returns true if each component is equal (without any
// fuzz).  Otherwise Equals returns false
func (a Vec) Equals(b Vec) bool {
	return a[0] == b[0] && a[1] == b[1]
}

// EqualWithFuzz compares this Vec to the one provided using some
// amount of fuzz
func (a Vec) EqualWithFuzz(fuzz float64, b Vec) bool {
	v := a.Sub(b).Mag()
	return v <= fuzz
}

func (a Vec) EqualAny(fuzz float64, vs ...Vec) bool {
	for i := 0; i < len(vs); i++ {
		if a.EqualWithFuzz(fuzz, vs[i]) {
			return true
		}
	}
	return false
}

// Mag returns the square root of the sum of squares which is
// the magnitude of the vector.
func (v Vec) Mag() float64 {
	return math.Sqrt(v.Dot(v))
}

// Scale multiplies each component by the value s.
func (v Vec) Scale(x, y float64) Vec {
	v[0] *= x
	v[1] *= y
	return v
}

// ScaleX returns a Vec where the X is scaled by the given factor
func (v Vec) ScaleX(s float64) Vec {
	return Vec{v[0] * s, v[1]}
}

// ScaleY returns a Vec where the Y is scaled by the given factor
func (v Vec) ScaleY(s float64) Vec {
	return Vec{v[0], v[1] * s}
}

func (v Vec) ScaleXY(scale float64) Vec {
	return Vec{v[0] * scale, v[1] * scale}
}

// String produces a string in the form: "[x,y]"
func (v Vec) String() string {
	x, y := v.Components()
	fuzz := Fuzz(0.00001)
	if fuzz.Eq(x, 0.0) {
		x = 0.0
	}
	if fuzz.Eq(y, 0.0) {
		y = 0.0
	}
	return fmt.Sprintf("[%.2f, %.2f]", x, y)
}

// Normalize produces a unit vector from v.  However, if the vector is the
// zero vector then a zero vector is returned.  Normalizing is
// mathematically: (1 / magnitude) * v.  Which explains why Normalizing
// the Zero vector would cause a division by 0 and why the zero vector is
// returned. (May change to return an error in the future.)
//
// TODO: consider returning an error here with the vector, to handle
//
//	normalizing a zero vector.
func (v Vec) Normalize() Vec {
	d := v.Mag()
	if d == 0.0 {
		return Vec{}
	}
	return v.ScaleXY(1.0 / d)
}

// Add returns a new vector whereby the two vectors produce a new vector
// from adding each corresponding component of the two vectors.
func (a Vec) Add(b Vec) Vec {
	return Vec{
		a[0] + b[0],
		a[1] + b[1],
	}
}

// Sub is a convenience for a.Add(b.Negate()).
func (a Vec) Sub(b Vec) Vec {
	return a.Add(b.Negate())
}

// Translate adds the delta x,y values to the given Vec.
func (a Vec) Translate(x, y float64) Vec {
	return Vec{a[0] + x, a[1] + y}
}

// Negate returns a new vector with the same magnitude just pointing
// in the opposite direction.
func (a Vec) Negate() Vec {
	return a.ScaleXY(-1.0)
}

// Dot computes the dot-product between this vector and the provided
// vector.
func (a Vec) Dot(b Vec) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// Angle finds the rotation of the given vector.
func (a Vec) Angle() float64 {
	if a == Zero {
		return 0.0
	}
	n := a.Normalize()
	b := math.Acos(n.X())
	if n.Y() >= 0.0 {
		return b
	} else {
		return -b
	}
}

// Rotate returns a vector as rotated about the origin by the given
// angle.
func (v Vec) Rotate(theta float64) Vec {
	x, y := v.X(), v.Y()
	xp := (x * math.Cos(theta)) - (y * math.Sin(theta))
	yp := (y * math.Cos(theta)) + (x * math.Sin(theta))
	return Vec{xp, yp}
}

// Transpose returns a Vec where x and y are swapped.
func (v Vec) Transpose() Vec {
	return Vec{v.Y(), v.X()}
}

// Perp returns a Vec perpendicular to this Vec
func (v Vec) Perp() Vec {
	return v.Rotate(math.Pi / 2.0)
}

// Copy returns a copy of the Vec
func (v Vec) Copy() Vec {
	return Vec{v.X(), v.Y()}
}

// VecX return a new Vec made from the <X,0>
func (v Vec) VecX() Vec {
	return Vec{v.X(), 0}
}

// VecY return a new Vec made from the <0,Y>
func (v Vec) VecY() Vec {
	return Vec{0, v.Y()}
}

func (v Vec) Half() Vec {
	return v.ScaleXY(0.5)
}

func (a Vec) AngleBetween(b Vec) float64 {
	magA := a.Mag()
	magB := b.Mag()
	dotAB := a.Dot(b)
	v := dotAB / (magA * magB)
	t := math.Acos(v)
	return t
}

func (a Vec) Cross(b Vec) float64 {
	magA := a.Mag()
	magB := b.Mag()
	dotAB := a.Dot(b)
	ab := magA * magB
	v := dotAB / ab
	t := math.Acos(v)
	return ab * math.Sin(t)
}

// CrossOn is the cross-product standing at 'a' looking at 'b'
// considering the side of 'c'
func (a Vec) CrossOn(b, c Vec) float64 {
	y1 := a.Y() - b.Y()
	y2 := a.Y() - c.Y()
	x1 := a.X() - b.X()
	x2 := a.X() - c.X()
	return y2*x1 - y1*x2
}

// Pt provides the x, y into the grid where this PointAgent resides.
func (a Vec) IntComponents() (x, y int) {
	return int(a.X()), int(a.Y())
}

// Point converts the vec to image.Point
func (a Vec) Point() image.Point {
	x, y := a.X(), a.Y()
	return image.Point{X: int(x), Y: int(y)}
}

// Octant returns the 1-based anti-clockwise number of the octant (or
// edge) which the Vec points toward.  Given that the unit vectors
// bridge the each quadrant, and therefor each octant, a vector
// parallel to an axis returns a number greater than the max octant.
// Each vec where x == y follows the edges of each quadrant.
func (a Vec) Octant() Octant {
	x, y := a.Components()

	xZero := x == 0.0
	yZero := y == 0.0

	if xZero && yZero { // no quad, octant, or edge
		return NoDirection
	}

	e := x > 0.0 // east-ward
	n := y > 0.0 // north-ward
	w := x < 0.0 // west-ward
	s := y < 0.0 // sourth-ward

	if xZero {
		if n {
			return North
		}
		return South
	}
	if yZero {
		if w {
			return West
		}
		return East
	}

	m := y / x
	pos := m > 0.0
	neg := m < 0.0
	edge := math.Abs(y) / math.Abs(x)

	if pos {
		if edge == 1.0 {
			if n {
				return NE
			}
			if s {
				return SW
			}
		}
		if edge < 1.0 {
			if e {
				return ENE
			}
			if w {
				return WSW
			}
		}
		if edge > 1.0 {
			if e {
				return NNE
			}
			if w {
				return SSW
			}
		}
	}
	if neg {
		if edge == 1.0 {
			if n {
				return NW
			}
			if s {
				return SE
			}
		}
		if edge < 1.0 {
			if e {
				return ESE
			}
			if w {
				return WNW
			}
		}
		if edge > 1.0 {
			if e {
				return SSE
			}
			if w {
				return NNW
			}
		}
	}
	return NoDirection
}

type Octant int

const (
	NoDirection Octant = 0

	ENE Octant = 1 // east-north-east
	NNE Octant = 2 // north-north-east
	NNW Octant = 3 // north-north-weast
	WNW Octant = 4 // west-north-west
	WSW Octant = 5 // west-south-west
	SSW Octant = 6 // south-south-west
	SSE Octant = 7 // south-south-east
	ESE Octant = 8 // east-sourth-east

	East  Octant = 9
	North Octant = 10
	West  Octant = 11
	South Octant = 12

	NE Octant = 13 // north-east
	NW Octant = 14 // north-east
	SW Octant = 15 // south-west
	SE Octant = 16 // south-east
)

func (oct Octant) Upward() bool {
	switch oct {
	case North, ENE, NNE, NNW, WNW, NE, NW:
		return true
	}
	return false
}

func (oct Octant) Downward() bool {
	switch oct {
	case South, ESE, SSE, SSW, WSW, SW, SE:
		return true
	}
	return false
}

func (oct Octant) Rightward() bool {
	switch oct {
	case East, NE, SE, ENE, NNE, SSE, ESE:
		return true
	}
	return false
}

func (oct Octant) Leftward() bool {
	switch oct {
	case West, NW, SW, NNW, WNW, WSW, SSW:
		return true
	}
	return false
}
