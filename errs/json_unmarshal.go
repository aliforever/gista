package errs

type JsonUnmarshal GistaError

func (ad JsonUnmarshal) Error() string {
	m := "unknown"
	if ad.Message != nil {
		m = *ad.Message
	}
	return m
}
