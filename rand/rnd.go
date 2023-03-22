package rand

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/lcaballero/ebiten-01/shapes"
)

// Rnd for generating a random number for a component
type Rnd struct {
	RNG  *rand.Rand
	Seed int64 `yaml:"seed"`
}

// NewRnd tying the name of the component and the seed for a random
// number generateor
func NewRnd(seed int64) Rnd {
	return Rnd{
		RNG:  rand.New(rand.NewSource(seed)),
		Seed: seed,
	}
}

// UnmarshalYAML unpacks the Rnd from Yaml
func (rnd *Rnd) UnmarshalYAML(fn func(interface{}) error) error {
	r := struct {
		Seed int64 `yaml:"seed"`
	}{}
	fn(&r)
	fmt.Println(r)
	*rnd = NewRnd(r.Seed)
	return nil
}

// Copy creates a new Rnd based on the seed of thie Rnd
func (rnd Rnd) Copy() Rnd {
	return NewRnd(rnd.Seed)
}

// Float64 generates a number between [0, 1)
func (rnd Rnd) F64() float64 {
	return rnd.RNG.Float64()
}

// Less returns true if a random number [0,1] is less than the value
func (rnd Rnd) Less(prob float64) bool {
	return rnd.In(0, 1) < prob
}

func (rnd Rnd) Vals(fs ...float64) []float64 {
	n := len(fs) / 2
	rs := make([]float64, n)
	for i, j := 1, 0; i < len(fs); i, j = i+2, j+1 {
		a, b := fs[i-1], fs[i]
		if b-a > 0 {
			rs[j] = rnd.In(a, b)
		}
	}
	return rs
}

// Between generates a non-inclusive float64 between (a, b)
func (rnd Rnd) In(a, b float64) float64 {
	f := (b - a) * rnd.F64()
	return a + f
}

// Pos returns a Vec randomly shifted by the given Orthogonal vectors
func (rnd Rnd) Pos(p, w, h shapes.Vec) shapes.Vec {
	x := rnd.In(w.Components())
	y := rnd.In(h.Components())
	v := shapes.Vec{x, y}
	return v.Add(p)
}

func (rnd Rnd) Centered(c float64) float64 {
	f := rnd.F64()
	r := (f * 2.0) - 1.0
	return c + r
}

// ZeroCentered produces random number between (-1.0, 1.0)
func (rnd Rnd) ZeroCentered() float64 {
	return rnd.Centered(0.0)
}

// ZC is an alias to ZeroCentered
func (rnd Rnd) ZC() float64 {
	return rnd.ZeroCentered()
}

func (rnd Rnd) OverInt(n int) func() int {
	return func() int {
		return rnd.Int(n)
	}
}

// UnitVec produces a unit vector with an angle between (pi, -pi)
func (rnd Rnd) UnitVec() shapes.Vec {
	a := rnd.ZeroCentered() * math.Pi
	return shapes.Vec{math.Cos(a), math.Sin(a)}
}

// Int picks a number between [a,b]
func (rnd Rnd) Intn(a, b int) int {
	return a + rnd.RNG.Intn(b-a)
}

// Mod over a random number for every n it returns true
func (rnd Rnd) Mod(in int) bool {
	return rnd.Int(in+1)%in == 0
}

// Int returns integer from [0,a)
func (rnd Rnd) Int(a int) int {
	return rnd.RNG.Intn(a)
}

// Shuffle calls the given func passing indexes that to swap
func (rnd Rnd) Shuffle(n int, fn func(j, k int)) {
	rnd.RNG.Shuffle(n, fn)
}

func (rnd Rnd) OnCircle(rads, radius float64) shapes.Vec {
	return shapes.UnitX.
		Rotate(rnd.F64() * rads).
		ScaleXY(radius * rnd.F64())
}

// Bool returns true or false randomly
func (rnd Rnd) Bool() bool {
	return rnd.F64() < 0.5
}

func (rnd Rnd) ChooseVec(vecs ...shapes.Vec) (shapes.Vec, int) {
	n := rnd.Int(len(vecs))
	return vecs[n], n
}

// F returns a random valued scaled by the first parameter if there is
// only 1, else between the two numbers
func (rnd Rnd) F(fs ...float64) float64 {
	n := len(fs)
	switch n {
	case 1:
		return rnd.F64() * fs[0]
	case 2:
		a := fs[0]
		b := fs[1]
		df := b - a
		return a + (rnd.F64() * df)
	default:
		return rnd.F64()
	}
}

func (rnd Rnd) InVec(v shapes.Vec) shapes.Vec {
	x := rnd.F64() * v.X()
	y := rnd.F64() * v.Y()
	return shapes.Vec{x, y}
}

func (rnd Rnd) InRect(r shapes.Rect) shapes.Vec {
	v := rnd.InVec(r.Dim)
	return v.Add(r.Pos)
}

func (rnd Rnd) Vecs(n int, dim shapes.Vec) shapes.Vecs {
	vecs := make(shapes.Vecs, n)
	for i := 0; i < n; i++ {
		vecs[i] = rnd.InVec(dim)
	}
	return vecs
}
