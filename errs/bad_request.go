package errs

type BadRequest GistaError

func (br BadRequest) Error() string {
	m := "unknown"
	if br.Message != nil {
		m = *br.Message
	}
	return m
}
