package gista

type usertag struct {
	ig *Instagram
}

func newUsertag(i *Instagram) *usertag {
	return &usertag{ig: i}
}
