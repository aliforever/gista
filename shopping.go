package gista

type shopping struct {
	ig *Instagram
}

func newShopping(i *Instagram) *shopping {
	return &shopping{ig: i}
}
