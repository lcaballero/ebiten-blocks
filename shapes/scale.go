package shapes

// Scale represents the ratio which to apply to X and Y of the context
type Scale Vec

// Invert scrates a scale that will undo the Scale if both were
// applied in succession
func (s Scale) Invert() Scale {
	x, y := Vec(s).Components()
	v := Vec{1.0 / x, 1 / y}
	return Scale(v)
}

// Vec returns this scale as a Vec
func (s Scale) Vec() Vec {
	return Vec(s)
}
