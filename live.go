package gista

type live struct {
	ig *Instagram
}

func newLive(i *Instagram) *live {
	return &live{ig: i}
}
