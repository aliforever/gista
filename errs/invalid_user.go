package errs

type InvalidUser GistaError

func (iu InvalidUser) Error() string {
	m := "unknown"
	if iu.Message != nil {
		m = *iu.Message
	}
	return m
}
