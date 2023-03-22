package shapes

type Arc struct {
	Pos    Vec
	Radius float64
	A      float64
	B      float64
}

func (a Arc) Translate(v Vec) Arc {
	return Arc{
		Pos:    a.Pos.Add(v),
		Radius: a.Radius,
		A:      a.A,
		B:      a.B,
	}
}

func (a Arc) Expand(dt float64) Arc {
	return Arc{
		Pos:    a.Pos,
		Radius: a.Radius + dt,
		A:      a.A,
		B:      a.B,
	}
}
