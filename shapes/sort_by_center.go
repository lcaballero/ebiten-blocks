package shapes

type SortByCenter Rects

func (sc SortByCenter) Len() int {
	return len(sc)
}

func (sc SortByCenter) Swap(i, j int) {
	sc[i], sc[j] = sc[j], sc[i]
}

func (sc SortByCenter) Less(i, j int) bool {
	a := sc[i].Center()
	b := sc[j].Center()
	if a.Y() == b.Y() {
		return a.X() < b.X()
	}
	if a.X() < b.X() {
		return a.Y() < b.Y()
	}
	if a.Y() < b.Y() {
		return true
	}
	return a.X() < b.X()
}
