package shapes

import (
	"math"
	"sort"
)

type VecsByAngle Vecs

func (v VecsByAngle) Len() int {
	return len(v)
}

func (v VecsByAngle) Less(i, j int) bool {
	a := v[i].Angle() + math.Pi
	b := v[j].Angle() + math.Pi
	return a < b
}

func (v VecsByAngle) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v VecsByAngle) Sort() Vecs {
	sort.Sort(v)
	return Vecs(v)
}
