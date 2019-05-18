package errs

type LoginRequired GistaError

func (lr LoginRequired) Error() string {
	m := "unknown"
	if lr.Message != nil {
		m = *lr.Message
	}
	return m
}
