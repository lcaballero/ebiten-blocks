package shapes

import (
	"math"
)

// Circle (a complete 2Pi) is defined as the Center and the Radius.
type Circle struct {
	Pos    Vec
	Radius float64
	Id     int
}

// X position returned
func (c Circle) X() float64 {
	return c.Pos.X()
}

// Y position returned
func (c Circle) Y() float64 {
	return c.Pos.Y()
}

// R size of the circle
func (c Circle) R() float64 {
	return c.Radius
}

// Start returns a point on the circle at 0 radians
func (c Circle) Start() Vec {
	return c.Pos.Add(Vec{c.Radius, 0.0})
}

// At returns a point on the circle at theta radians
func (c Circle) At(theta float64) Vec {
	return c.Start().Rotate(theta)
}

// Top returns a Vec to the top of the circle
func (c Circle) Top() Vec {
	return c.Start().Rotate(math.Pi / 2.0)
}

// Left returns a Vec of the left of the circle
func (c Circle) Left() Vec {
	return c.Start().Rotate(math.Pi)
}

// Right returns a Vec of the right of the circle
func (c Circle) Right() Vec {
	return c.Start()
}

// Bottom returns the a Vec of the bottom of the circle
func (c Circle) Bottom() Vec {
	return c.Start().Rotate(math.Pi * 1.5)
}

// Points returns a slice of Right, Top, Left, Bottom of the circle
func (c Circle) Points() []Vec {
	return []Vec{
		c.Right(),
		c.Top(),
		c.Left(),
		c.Bottom(),
	}
}
