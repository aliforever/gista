package gista

type collection struct {
	ig *Instagram
}

func newCollection(i *Instagram) *collection {
	return &collection{ig: i}
}
