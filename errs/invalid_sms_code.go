package errs

type InvalidSmsCode GistaError

func (isc InvalidSmsCode) Error() string {
	m := "unknown"
	if isc.Message != nil {
		m = *isc.Message
	}
	return m
}
