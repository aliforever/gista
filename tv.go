package gista

type tv struct {
	ig *Instagram
}

func newTv(i *Instagram) *tv {
	return &tv{ig: i}
}
