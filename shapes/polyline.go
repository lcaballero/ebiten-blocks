package shapes

// A quadrolateral shape, so once with four sides, but no gaurantee
// similar angles or side lengths, unlike a Kite, Trapezoide,
// Parallelogram, Rhombus, Squares, Rectangles, etc.
type Quadro struct {
	Rot  Vec
	Pos  Vec
	Vecs Vecs
	Id   int
}

// NewParallelQuadro creates a parallelogram with object origin set to
// the bottom left corner
func NewParallelQuadro(rot float64, pos, dim Vec) Quadro {
	r := Rect{Dim: dim}
	vecs := r.Corners().Alter(func(i int, v Vec) Vec {
		return v.Rotate(rot).Translate(pos.Components())
	})
	rt := UnitX.Rotate(rot)
	return Quadro{
		Rot:  rt,
		Pos:  pos,
		Vecs: vecs,
	}
}

// Rotate alters the coodinates of the Quad to rotate around the
// object origin (lower left corner by default)
func (q Quadro) Rotate(t float64) Quadro {
	rot := q.Rot.Rotate(t)
	theta := rot.Angle()
	vecs := q.Vecs.Alter(func(i int, v Vec) Vec {
		return v.Sub(q.Pos).Rotate(theta).Add(q.Pos)
	})
	return Quadro{
		Id:   q.Id,
		Rot:  rot,
		Pos:  q.Pos,
		Vecs: vecs,
	}
}

// Translate alters the coordinates of the Quadro translating them by
// the given vector
func (q Quadro) Translate(v1 Vec) Quadro {
	pos := q.Pos.Add(v1)
	vecs := q.Vecs.Alter(func(i int, v0 Vec) Vec {
		return v0.Add(v1)
	})
	return Quadro{
		Id:   q.Id,
		Rot:  q.Rot,
		Pos:  pos,
		Vecs: vecs,
	}
}

func (q Quadro) Alter(fn func(i int, v Vec) Vec) Quadro {
	p := Zero
	vecs := q.Vecs.Alter(func(j int, v0 Vec) Vec {
		v1 := fn(j, v0)
		if j == 0 {
			p = v1
		}
		return v1
	})
	return Quadro{Id: q.Id, Rot: q.Rot, Pos: p, Vecs: vecs}
}

func (q Quadro) Divide() Quadro {
	dupes := q.Vecs
	rv := Vecs{}
	kn := len(q.Vecs)
	for i := 0; i < kn; i++ {
		n := (i + 1) % kn
		a := dupes[i]
		b := dupes[n]
		mid := b.Sub(a).ScaleXY(0.5).Add(a)
		rv = append(rv, a)
		rv = append(rv, mid)
	}
	return Quadro{
		Id:   q.Id,
		Rot:  q.Rot,
		Pos:  q.Pos,
		Vecs: rv,
	}
}
