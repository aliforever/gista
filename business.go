package gista

type business struct {
	ig *Instagram
}

func newBusiness(i *Instagram) *business {
	return &business{ig: i}
}
