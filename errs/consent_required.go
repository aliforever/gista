package errs

type ConsentRequired GistaError

func (cr ConsentRequired) Error() string {
	m := "unknown"
	if cr.Message != nil {
		m = *cr.Message
	}
	return m
}
