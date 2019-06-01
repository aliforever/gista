package gista

type hashtag struct {
	ig *Instagram
}

func newHashtag(i *Instagram) *hashtag {
	return &hashtag{ig: i}
}
