package errs

type AccountDisabled GistaError

func (ad AccountDisabled) Error() string {
	m := "unknown"
	if ad.Message != nil {
		m = *ad.Message
	}
	return m
}
