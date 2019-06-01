package gista

type location struct {
	ig *Instagram
}

func newLocation(i *Instagram) *location {
	return &location{ig: i}
}
