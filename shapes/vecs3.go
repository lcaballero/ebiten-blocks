package shapes

import (
	"math"
)

type Vec3 [3]float64

func (a Vec3) Scale(s float64) Vec3 {
	return Vec3{a[0] * s, a[1] * s, a[2] * s}
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func (a Vec3) Negate() Vec3 {
	return Vec3{-a[0], -a[1], -a[2]}
}

func (a Vec3) Cos() Vec3 {
	return Vec3{
		math.Cos(a[0]),
		math.Cos(a[1]),
		math.Cos(a[2]),
	}
}

func (a Vec3) Components() (float64, float64, float64) {
	return a[0], a[1], a[2]
}

func (a Vec3) Mult(b Vec3) Vec3 {
	x, y, z := a[0]*b[0], a[1]*b[1], a[2]*b[2]
	return Vec3{x, y, z}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}
