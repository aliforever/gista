package errs

type NotAuthorized GistaError

func (na NotAuthorized) Error() string {
	m := "unknown"
	if na.Message != nil {
		m = *na.Message
	}
	return m
}
