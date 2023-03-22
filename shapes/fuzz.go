package shapes

type Fuzz float64

func (f Fuzz) Eq(a, b float64) bool {
	d := float64(f)
	return a-d <= b && b <= a+d
}

// Between returns true if v is greater than or equal to 'a' and less
// than or equal to 'b', inclusive.
func (f Fuzz) Between(v, a, b float64) bool {
	k := float64(f)
	return (a-k) <= v && v <= (b+k)
}

func (f Fuzz) Less(v, max float64) bool {
	d := float64(f)
	return v < (max - d)
}

func (f Fuzz) Greater(v, min float64) bool {
	d := float64(f)
	return v > (min + d)
}
