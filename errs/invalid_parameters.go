package errs

type InvalidParameters GistaError

func (ip InvalidParameters) Error() string {
	m := "unknown"
	if ip.Message != nil {
		m = *ip.Message
	}
	return m
}
