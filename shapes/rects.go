package shapes

import "sort"

type Rects []Rect

func (rects Rects) Translate(v Vec) Rects {
	rs := Rects{}
	for _, r := range rects {
		rs = append(rs, r.Translate(v))
	}
	return rs
}

func (rects Rects) SortByCenter() Rects {
	sort.Sort(SortByCenter(rects))
	return rects
}
