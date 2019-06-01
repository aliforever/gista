package gista

type creative struct {
	ig *Instagram
}

func newCreative(i *Instagram) *creative {
	return &creative{ig: i}
}
