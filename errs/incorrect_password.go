package errs

type IncorrectPassword GistaError

func (ip IncorrectPassword) Error() string {
	m := "unknown"
	if ip.Message != nil {
		m = *ip.Message
	}
	return m
}
